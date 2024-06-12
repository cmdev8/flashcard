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
