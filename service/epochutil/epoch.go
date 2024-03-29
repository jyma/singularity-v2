package epochutil

import (
	"context"
	"strings"
	"time"

	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/pkg/errors"
)

var GenesisTimestamp = int32(1598306400)

type result struct {
	Blocks []block
}

type block struct {
	Timestamp int32
}

func Initialize(ctx context.Context, lotusAPI string, lotusToken string) error {
	if strings.HasPrefix(lotusAPI, "https://api.node.glif.io/rpc") {
		GenesisTimestamp = int32(1598306400)
		return nil
	}

	client := util.NewLotusClient(lotusAPI, lotusToken)
	var r result
	err := client.CallFor(ctx, &r, "Filecoin.ChainGetGenesis")
	if err != nil {
		return errors.Wrap(err, "failed to decide genesis timestamp")
	}

	if len(r.Blocks) == 0 {
		return errors.New("length of blocks from genesis is 0")
	}

	GenesisTimestamp = r.Blocks[0].Timestamp
	return nil
}

func EpochToTime(epoch int32) time.Time {
	return time.Unix(int64(epoch)*30+int64(GenesisTimestamp), 0)
}

func UnixToEpoch(unix int64) abi.ChainEpoch {
	return abi.ChainEpoch(unix-int64(GenesisTimestamp)) / 30
}

func TimeToEpoch(t time.Time) abi.ChainEpoch {
	return abi.ChainEpoch((t.Unix() - int64(GenesisTimestamp)) / 30)
}
