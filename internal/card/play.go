package card

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetNextCard(db *gorm.DB, category string, now time.Time) (*Card, error) {
	cards, err := ListCards(db, category)
	if err != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return nil, nil
	}

	randomCard := lo.Sample(cards)

	return &randomCard, nil
}
