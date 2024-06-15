package card

import (
	"time"

	"github.com/samber/lo"
	"gorm.io/gorm"
)

func GetNextCard(db *gorm.DB, category string, notpracticedBefore time.Time) (*Card, error) {
	cards, err := ListCards(db, category)
	if err != nil {
		return nil, err
	}

	if len(cards) == 0 {
		return nil, nil
	}

	// results, _ := ListResultAfter(db, notpracticedBefore)
	// cards = lo.Filter(cards, func(item Card, _ int) bool {
	// 	return !lo.ContainsBy(results, func(ri Result) bool {
	// 		return ri.CardID == item.ID
	// 	})
	// })

	randomCard := lo.Sample(cards)

	return &randomCard, nil
}
