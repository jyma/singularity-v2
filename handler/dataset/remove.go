package dataset

import (
	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/handler"
	"gorm.io/gorm"
)

// @Summary Remove a specific dataset. This will not remove the CAR files.
// @Description Important! If the dataset is large, this command will take some time to remove all relevant data.
// @Tags Dataset
// @Param datasetName path string true "Dataset name"
// @Success 204
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dataset/{datasetName} [delete]
func removeHandler(
	db *gorm.DB,
	datasetName string,
) error {
	dataset, err := database.FindDatasetByName(db, datasetName)
	if err != nil {
		return handler.NewInvalidParameterErr("failed to find dataset: " + err.Error())
	}
	err = database.DoRetry(func() error { return db.Delete(&dataset).Error })
	if err != nil {
		return err
	}
	return nil
}

func RemoveHandler(
	db *gorm.DB,
	datasetName string,
) error {
	return removeHandler(db, datasetName)
}
