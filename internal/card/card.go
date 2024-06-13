package card

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID            uint
	Category      string
	QuestionText  string
	QuestionImage string
	AnswerText    string
	AnswerImage   string
	CreatedAt     time.Time
}

func CreateCard(db *gorm.DB, newCard *Card) error {
	return db.Create(newCard).Error
}

func ListCards(db *gorm.DB, category string) ([]Card, error) {
	var r []Card
	q := db.Order("id desc")

	if category != "" {
		q = q.Where("category = ?", category)
	}

	err := q.Find(&r).Error

	return r, err
}

func DeleteCard(db *gorm.DB, id uint) error {
	var card Card
	if err := db.First(&card, id).Error; err != nil {
		return err
	}

	return db.Delete(&card).Error
}
