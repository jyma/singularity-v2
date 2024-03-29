package wallet

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

func RemoveWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) error {
	return removeWalletHandler(db, datasetName, wallet)
}

// @Summary Remove an associated wallet from a dataset
// @Tags Wallet
// @Param datasetName path string true "Dataset name"
// @Param wallet path string true "Wallet Address"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName}/wallet/{wallet} [delete]
func removeWalletHandler(
	db *gorm.DB,
	datasetName string,
	wallet string,
) error {
	if datasetName == "" {
		return handler.NewInvalidParameterErr("dataset name is required")
	}

	if wallet == "" {
		return handler.NewInvalidParameterErr("wallet address is required")
	}

	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}

	var w model.Wallet
	err = db.Where("address = ? OR id = ?", wallet, wallet).First(&w).Error
	if err != nil {
		return handler.NewInvalidParameterErr("failed to find wallet: " + err.Error())
	}

	err = database.DoRetry(func() error {
		return db.Where("dataset_id = ? AND wallet_id = ?", dataset.ID, w.ID).Delete(&model.WalletAssignment{}).Error
	})

	if err != nil {
		return err
	}

	return nil
}
