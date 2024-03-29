package dealmaker

import (
	"bytes"
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/pack"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/data-preservation-programs/singularity/service/epochutil"
	commp "github.com/filecoin-project/go-fil-commp-hashhash"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockDealMaker struct {
	mock.Mock
}

func (m *MockDealMaker) MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig replication.DealConfig) (*model.Deal, error) {
	args := m.Called(ctx, walletObj, car, dealConfig)
	deal := *args.Get(0).(*model.Deal)
	deal.ID = 0
	deal.PieceCID = car.PieceCID
	deal.PieceSize = car.PieceSize
	deal.DatasetID = &car.DatasetID
	deal.ClientID = walletObj.ID
	deal.Provider = dealConfig.Provider
	deal.Verified = dealConfig.Verified
	deal.ProposalID = uuid.NewString()
	deal.State = model.DealProposed
	now := time.Now()
	startEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay))
	endEpoch := epochutil.TimeToEpoch(now.Add(dealConfig.StartDelay + dealConfig.Duration))
	if deal.StartEpoch == 0 {
		deal.StartEpoch = int32(startEpoch)
	}
	if deal.EndEpoch == 0 {
		deal.EndEpoch = int32(endEpoch)
	}
	err := args.Error(1)
	if err != nil {
		return &deal, err
	}
	return &deal, nil
}

func TestDealMakerService_StartRun(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealMakerService(db, "https://api.node.glif.io", "")
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan struct{})
	go func() {
		err = service.Run(ctx)
		require.Error(t, err, context.Canceled)
		close(done)
	}()
	cancel()
	<-done
}

func TestDealMakerService_Cron(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealMakerService(db, "https://api.node.glif.io", "")
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// All deal proposal will be accepted
	// Create test schedule
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Dataset: &model.Dataset{Name: "test", Wallets: []model.Wallet{
			{
				ID: client, Address: "f0xx",
			},
		}},
		State:            model.ScheduleActive,
		ScheduleCron:     "0 0 1 1 1",
		ScheduleDealSize: 1,
		Provider:         provider,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)

	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ScheduleID: &schedule.ID,
	}, nil)

	err = db.Create([]model.Car{
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize: 1024,
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize: 1024,
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
			PieceSize: 1024,
		},
	}).Error

	require.NoError(t, err)
	service.cron.Start()
	defer service.cron.Stop()
	service.runOnce(ctx)

	err = db.Model(&schedule).Update("schedule_cron", "* * * * * *").Error
	require.NoError(t, err)
	service.runOnce(ctx)

	time.Sleep(1 * time.Second)
	var deals []model.Deal
	err = db.Find(&deals).Error
	require.NoError(t, err)
	ndeals := len(deals)
	require.True(t, ndeals == 1 || ndeals == 2)

	err = db.Model(&schedule).Update("state", model.SchedulePaused).Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(2 * time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, ndeals)

	db.Model(&schedule).Update("state", model.ScheduleActive)
	service.runOnce(ctx)
	time.Sleep(3 * time.Second)
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Greater(t, len(deals), ndeals)
}

func TestDealMakerService_NewScheduleOneOff(t *testing.T) {
	db, closer, err := database.OpenInMemory()
	require.NoError(t, err)
	defer closer.Close()
	service, err := NewDealMakerService(db, "https://api.node.glif.io", "")
	require.NoError(t, err)
	mockDealmaker := new(MockDealMaker)
	service.dealMaker = mockDealmaker
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// All deal proposal will be accepted
	// Create test schedule
	provider := "f0miner"
	client := "f0client"
	schedule := model.Schedule{
		Dataset: &model.Dataset{Name: "test", Wallets: []model.Wallet{
			{
				ID: client, Address: "f0xx",
			},
		}},
		State:    model.ScheduleActive,
		Provider: provider,
	}
	err = db.Create(&schedule).Error
	require.NoError(t, err)

	mockDealmaker.On("MakeDeal", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ScheduleID: &schedule.ID,
	}, nil)

	pieceCIDs := []model.CID{
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
		model.CID(calculateCommp(t, generateRandomBytes(1000), 1024)),
	}
	err = db.Create([]model.Car{
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  pieceCIDs[0],
			PieceSize: 1024,
			FilePath:  "0",
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  pieceCIDs[1],
			PieceSize: 1024,
			FilePath:  "1",
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  pieceCIDs[2],
			PieceSize: 1024,
			FilePath:  "2",
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  pieceCIDs[3],
			PieceSize: 1024,
			FilePath:  "3",
		},
		{
			DatasetID: schedule.Dataset.ID,
			PieceCID:  pieceCIDs[4],
			PieceSize: 1024,
			FilePath:  "4",
		},
	}).Error
	require.NoError(t, err)

	// Test1 is already proposed
	// Test2 is expired proposal
	// Test3 is active
	// Test4 is proposed by other schedules
	// Test5 is not proposed
	err = db.Create([]model.Deal{
		{
			DatasetID:  &schedule.Dataset.ID,
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[0],
			PieceSize:  1024,
			State:      model.DealProposed,
		},
		{
			DatasetID:  &schedule.Dataset.ID,
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[1],
			PieceSize:  1024,
			State:      model.DealProposalExpired,
		},
		{
			DatasetID:  &schedule.Dataset.ID,
			ScheduleID: &schedule.ID,
			Provider:   provider,
			ClientID:   client,
			PieceCID:   pieceCIDs[2],
			PieceSize:  1024,
			State:      model.DealActive,
		},
		{
			DatasetID: &schedule.Dataset.ID,
			Provider:  provider,
			ClientID:  client,
			PieceCID:  pieceCIDs[3],
			PieceSize: 1024,
			State:     model.DealProposed,
		},
	}).Error
	require.NoError(t, err)
	service.runOnce(ctx)
	time.Sleep(time.Second)
	var deals []model.Deal
	err = db.Find(&deals).Error
	require.NoError(t, err)
	require.Len(t, deals, 6)
	require.Equal(t, pieceCIDs[1], deals[4].PieceCID)
	require.Equal(t, pieceCIDs[4], deals[5].PieceCID)
}

func calculateCommp(t *testing.T, content []byte, targetPieceSize uint64) cid.Cid {
	calc := &commp.Calc{}
	_, err := bytes.NewBuffer(content).WriteTo(calc)
	require.NoError(t, err)
	c, _, err := pack.GetCommp(calc, targetPieceSize)
	require.NoError(t, err)
	return c
}

func generateRandomBytes(n int) []byte {
	b := make([]byte, n)
	rand.Read(b)
	return b
}
