package pack

import (
	"bytes"
	"context"
	"github.com/data-preservation-programs/go-singularity/datasource"
	"github.com/data-preservation-programs/go-singularity/model"
	commcid "github.com/filecoin-project/go-fil-commcid"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	blocks "github.com/ipfs/go-block-format"
	"github.com/ipfs/go-cid"
	chunk "github.com/ipfs/go-ipfs-chunker"
	util "github.com/ipfs/go-ipfs-util"
	cbor "github.com/ipfs/go-ipld-cbor"
	format "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-log/v2"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	unixfs_pb "github.com/ipfs/go-unixfs/pb"
	"github.com/ipld/go-car"
	"github.com/multiformats/go-varint"
	"github.com/pkg/errors"
	"io"
	"os"
	"path"
)

const CidLength = 32

type BlockResult struct {
	CID    cid.Cid
	Offset uint64
	Length uint64
	Raw    []byte
	Error  error
}
type Result struct {
	CarFilePath string
	CarFileSize uint64
	PieceCID    cid.Cid
	PieceSize   uint64
	RootCID     cid.Cid
	Header      []byte
	RawBlocks   []model.RawBlock
	ItemBlocks  []model.ItemBlock
	CarBlocks   []model.CarBlock
}

const ChunkSize int64 = 1 << 20
const NumLinkPerNode = 1024

type Link struct {
	format.Link
	ChunkSize uint64
}

func createParentNode(links []Link) (format.Node, uint64, error) {
	node := unixfs.NewFSNode(unixfs_pb.Data_File)
	total := uint64(0)
	for _, link := range links {
		node.AddBlockSize(link.ChunkSize)
		total += link.ChunkSize
	}
	nodeBytes, err := node.GetBytes()
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to get bytes from node")
	}
	pbNode := merkledag.NodeWithData(nodeBytes)
	pbNode.SetCidBuilder(merkledag.V1CidPrefix())
	for _, link := range links {
		err = pbNode.AddRawLink("", &link.Link)
		if err != nil {
			return nil, 0, errors.Wrap(err, "failed to add link to node")
		}
	}
	return pbNode, total, nil
}

func writeCarBlock(writer io.Writer, block blocks.BasicBlock) (uint64, error) {
	written := uint64(0)
	varintBytes := varint.ToUvarint(uint64(len(block.RawData()) + block.Cid().ByteLen()))
	n, err := io.Copy(writer, bytes.NewReader(varintBytes))
	if err != nil {
		return written, errors.Wrap(err, "failed to write varint")
	}
	written += uint64(n)

	n, err = io.Copy(writer, bytes.NewReader(block.Cid().Bytes()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write cid")
	}
	written += uint64(n)

	n, err = io.Copy(writer, bytes.NewReader(block.RawData()))
	if err != nil {
		return written, errors.Wrap(err, "failed to write raw")
	}
	written += uint64(n)
	return written, nil
}

func PackItems(
	ctx context.Context,
	handler datasource.Handler,
	items []model.Item,
	outDir string,
	pieceSize uint64) (*Result, error) {
	result := &Result{}
	offset := uint64(0)
	var headerBytes []byte

	calc := &commp.Calc{}
	var writer io.Writer = calc
	var filepath string
	if outDir != "" {
		filename := uuid.NewString() + ".car"
		filepath = path.Join(outDir, filename)
		file, err := os.Create(filepath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create file at "+filepath)
		}
		defer file.Close()
		writer = io.MultiWriter(calc, file)
	}

	for _, item := range items {
		links := make([]Link, 0)
		blockChan, err := streamItem(ctx, handler, item)
		if err != nil {
			return nil, errors.Wrap(err, "failed to stream item")
		}

		for block := range blockChan {
			links = append(links, Link{
				Link: format.Link{
					Name: "",
					Size: block.Length,
					Cid:  block.CID,
				}})
			if block.Error != nil {
				return nil, errors.Wrap(block.Error, "failed to stream block")
			}

			if offset == 0 {
				result.RootCID = block.CID
				header := car.CarHeader{
					Roots:   []cid.Cid{block.CID},
					Version: 0,
				}

				headerBytes, err = cbor.DumpObject(&header)
				if err != nil {
					return nil, errors.Wrap(err, "failed to dump header")
				}

				result.Header = headerBytes
				n, err := io.Copy(writer, bytes.NewReader(headerBytes))
				if err != nil {
					return nil, errors.Wrap(err, "failed to write header")
				}

				offset += uint64(n)
			}

			basicBlock, _ := blocks.NewBlockWithCid(block.Raw, block.CID)
			written, err := writeCarBlock(writer, *basicBlock)
			if err != nil {
				return nil, errors.Wrap(err, "failed to write block")
			}

			offset += written
			result.ItemBlocks = append(result.ItemBlocks, model.ItemBlock{
				CID:    block.CID.String(),
				ItemID: item.ID,
				Offset: block.Offset,
				Length: block.Length,
			})

			result.CarBlocks = append(result.CarBlocks, model.CarBlock{
				CID:    block.CID.String(),
				Offset: offset - block.Length,
				Length: written,
				Varint: block.Length + CidLength,
			})
		}

		for len(links) > 1 {
			newLinks := make([]Link, 0)
			for start := 0; start < len(links); start += NumLinkPerNode {
				newNode, total, err := createParentNode(links[start:Min(start+NumLinkPerNode, len(links))])
				if err != nil {
					return nil, errors.Wrap(err, "failed to create parent node")
				}

				basicBlock, _ := blocks.NewBlockWithCid(newNode.RawData(), newNode.Cid())
				written, err := writeCarBlock(writer, *basicBlock)
				if err != nil {
					return nil, errors.Wrap(err, "failed to write block")
				}
				offset += written

				result.RawBlocks = append(result.RawBlocks, model.RawBlock{
					CID:    basicBlock.Cid().String(),
					Block:  basicBlock.RawData(),
					Length: uint32(len(basicBlock.RawData())),
				})
				result.CarBlocks = append(result.CarBlocks, model.CarBlock{
					CID:    basicBlock.Cid().String(),
					Offset: offset - written,
					Length: written,
					Varint: uint64(len(basicBlock.RawData())) + CidLength,
				})

				newNodeSize, _ := newNode.Size()
				newLinks = append(newLinks, Link{
					ChunkSize: total,
					Link: format.Link{
						Name: "",
						Size: newNodeSize,
						Cid:  newNode.Cid(),
					}})
			}

			links = newLinks
		}
	}

	rawCommp, rawPieceSize, err := calc.Digest()
	if err != nil {
		return nil, errors.Wrap(err, "failed to calculate commp")
	}

	if rawPieceSize < pieceSize {
		rawCommp, err = commp.PadCommP(rawCommp, rawPieceSize, pieceSize)
		if err != nil {
			return nil, errors.Wrap(err, "failed to pad commp")
		}

		rawPieceSize = pieceSize
	} else if rawPieceSize > pieceSize {
		log.Logger("packing").Warn("piece size is larger than the target piece size")
	}

	commCid, err := commcid.DataCommitmentV1ToCID(rawCommp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert commp to cid")
	}

	result.PieceCID = commCid
	result.PieceSize = rawPieceSize
	result.CarFileSize = offset

	if outDir != "" {
		result.CarFilePath = path.Join(outDir, commCid.String()+".car")
	}
	if filepath != "" {
		err = os.Rename(filepath, result.CarFilePath)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create symlink")
		}
	}

	return result, nil
}

func Min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func streamItem(ctx context.Context, handler datasource.Handler, item model.Item) (<-chan BlockResult, error) {
	readStream, err := handler.Open(ctx, item.Path, item.Offset, item.Length)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}

	blockChan := make(chan BlockResult)
	chunker := chunk.NewSizeSplitter(readStream, ChunkSize)
	go func() {
		defer readStream.Close()
		defer close(blockChan)
		offset := item.Offset
		for {
			chunkerBytes, err := chunker.NextBytes()
			if err != nil {
				if err == io.EOF {
					return
				}
				blockChan <- BlockResult{Error: errors.Wrap(err, "failed to read chunk")}
				return
			}

			hash := util.Hash(chunkerBytes)
			c := cid.NewCidV1(cid.Raw, hash)
			blockChan <- BlockResult{
				CID:    c,
				Offset: offset,
				Length: uint64(len(chunkerBytes)),
				Raw:    chunkerBytes,
				Error:  nil,
			}

			offset += uint64(len(chunkerBytes))
		}
	}()

	return blockChan, err
}
