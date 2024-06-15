package card

import (
	"time"

	"gorm.io/gorm"
)

type Result struct {
	ID        uint
	CardID    uint
	Success   bool
	CreatedAt time.Time
}

func CreateResult(db *gorm.DB, cardID uint, success bool) error {
	return db.Create(&Result{
		CardID:  cardID,
		Success: success,
	}).Error
}

func ListResultAfter(db *gorm.DB, after time.Time) ([]Result, error) {
	var results []Result
	err := db.Where("created_at >= ?", after).Find(&results).Error

	return results, err
}
