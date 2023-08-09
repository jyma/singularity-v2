package datasource

import (
	"errors"
	"fmt"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/rjNemo/underscore"
	"gorm.io/gorm"
)

func ChunkItemHandler(
	db *gorm.DB,
	itemID uint64,
) (int64, error) {
	return chunkItemHandler(db, itemID)
}

// @Summary Chunk for the specified item
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path integer true "Item ID"
// @Success 201 {object}
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /item/{id}/chunk [post]
func chunkItemHandler(
	db *gorm.DB,
	itemID uint64,
) (int64, error) {
	var item model.Item
	err := db.Preload("Source.Dataset").Where("id = ?", itemID).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, handler.NewInvalidParameterErr("item not found")
	}
	if err != nil {
		return 0, err
	}

	var remainingParts []model.ItemPart
	err = db.Where("item_id = ? AND chunk_id is null", itemID).Order("item_parts.id asc").
		Find(&remainingParts).Error

	return ChunkItemParts(db, item.SourceID, item.Source.Dataset.MaxSize, remainingParts)
}

func ChunkItemParts(
	db *gorm.DB,
	sourceID uint32,
	maxChunkSize int64,
	remainingParts []model.ItemPart,
) (int64, error) {
	chunkSet := newChunkSet()

	for len(remainingParts) > 0 {
		nextChunk, err := nextAvailableChunk(db, sourceID)
		if err != nil {
			return 0, err
		}
		chunkSet.reset()
		chunkSet.add(nextChunk.ItemParts)
		for len(remainingParts) > 0 {
			if !chunkSet.addIfFits(remainingParts[0], maxChunkSize) {
				break
			}
			remainingParts = remainingParts[1:]
		}
		if len(chunkSet.itemParts) == 0 && len(remainingParts) > 0 {
			chunkSet.add(remainingParts[:1])
			remainingParts = remainingParts[1:]
		}
		// if we still have remaining parts, we've filled up this chunk
		chunkState := model.Created
		if len(remainingParts) > 0 {
			chunkState = model.Ready
		}
		err = updateChunk(db, nextChunk.ID, chunkState, chunkSet.itemIDs())
		if err != nil {
			return 0, err
		}
	}
	return chunkSet.carSize, nil
}

func SourceChunksReadyHandler(
	db *gorm.DB,
	sourceID uint32,
) error {
	return sourceChunksReadyHandler(db, sourceID)
}

// @Summary
// @Tags Data Source
// @Accept json
// @Produce json
// @Param id path integer true "Source ID"
// @Success 201
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /source/{id}/finalize [post]
func sourceChunksReadyHandler(
	db *gorm.DB,
	sourceID uint32,
) error {
	return database.DoRetry(func() error {
		return db.Model(&model.Chunk{}).Where("source_id = ? AND state = ?", sourceID, model.Created).Update("state", model.Ready).Error
	})
}

func updateChunk(
	db *gorm.DB,
	chunkID uint32,
	state model.WorkState,
	itemIDs []uint64,
) error {
	return database.DoRetry(func() error {
		return db.Transaction(
			func(db *gorm.DB) error {
				err := db.Model(&model.Chunk{}).Updates(model.Chunk{ID: chunkID, PackingState: state}).Error
				if err != nil {
					return fmt.Errorf("failed to create chunk: %w", err)
				}
				itemPartIDChunks := util.ChunkSlice(itemIDs, util.BatchSize)
				for _, itemPartIDChunks := range itemPartIDChunks {
					err = db.Model(&model.ItemPart{}).
						Where("id IN ?", itemPartIDChunks).Update("chunk_id", chunkID).Error
					if err != nil {
						return fmt.Errorf("failed to update items: %w", err)
					}
				}
				return nil
			},
		)
	})
}

type chunkSet struct {
	itemParts []model.ItemPart
	carSize   int64
}

const carHeaderSize = 59

func newChunkSet() *chunkSet {
	return &chunkSet{
		itemParts: make([]model.ItemPart, 0),
		// Some buffer for header
		carSize: carHeaderSize,
	}
}

func (cs *chunkSet) add(itemParts []model.ItemPart) {
	cs.itemParts = append(cs.itemParts, itemParts...)
	for _, itemPart := range itemParts {
		cs.carSize += toCarSize(itemPart.Length)
	}
}

func (cs *chunkSet) addIfFits(itemPart model.ItemPart, maxSize int64) bool {
	nextSize := toCarSize(itemPart.Length)
	if cs.carSize+nextSize > maxSize {
		return false
	}
	cs.itemParts = append(cs.itemParts, itemPart)
	cs.carSize += nextSize
	return true
}

func (cs *chunkSet) reset() {
	cs.itemParts = make([]model.ItemPart, 0)
	cs.carSize = carHeaderSize
}

func (cs *chunkSet) itemIDs() []uint64 {
	return underscore.Map(cs.itemParts, func(itemPart model.ItemPart) uint64 {
		return itemPart.ID
	})
}

func toCarSize(size int64) int64 {
	out := size
	nBlocks := size / 1024 / 1024
	if size%(1024*1024) != 0 {
		nBlocks++
	}

	// For each block, we need to add the bytes for the CID as well as varint
	out += nBlocks * (36 + 9)

	// For every 256 blocks, we need to add another block.
	// The block stores up to 256 CIDs and integers, estimate it to be 12kb
	if nBlocks > 1 {
		out += (((nBlocks - 1) / 256) + 1) * 12000
	}

	return out
}

func nextAvailableChunk(
	db *gorm.DB,
	sourceID uint32,
) (*model.Chunk, error) {
	var chunk model.Chunk
	err := db.Where(model.Chunk{SourceID: sourceID, PackingState: model.Created}).Preload("ItemParts").FirstOrCreate(&chunk).Error
	if err != nil {
		return nil, err
	}
	return &chunk, nil
}
