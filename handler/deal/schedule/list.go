package schedule

import (
	"github.com/data-preservation-programs/singularity/model"
	"gorm.io/gorm"
)

// @Summary List all deal making schedules
// @Tags Deal Schedule
// @Produce json
// @Success 200 {array} model.Schedule
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /schedules [get]
func listHandler(
	db *gorm.DB,
) ([]model.Schedule, error) {
	var schedules []model.Schedule
	err := db.Find(&schedules).Error
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func ListHandler(
	db *gorm.DB,
) ([]model.Schedule, error) {
	return listHandler(db)
}
