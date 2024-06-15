package card

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetNextCard(db *gorm.DB, category string, practicePeriod time.Time) (*Card, error) {
	cards, err := ListCards(db, category)
	if err != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return nil, nil
	}

	practicedCardResults, _ := ListResultAfter(db, practicePeriod)
	cards = lo.Filter(cards, func(item Card, _ int) bool {
		for _, v := range practicedCardResults {
			if v.CardID == item.ID {
				return false
			}
		}

		return true
	})

	if len(cards) == 0 {
		return nil, nil
	}

	randomCard := lo.Sample(cards)

	return &randomCard, nil
}
