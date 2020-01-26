package card

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreateCards(t *testing.T) {
	start := time.Now()
	defer func() {
		end := time.Now()
		fmt.Println("exec time: ", end.Sub(start))
	}()

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	numberOfCards := 10
	cards := make([]*Base, numberOfCards)
	attr := []string{ // 属性
		"🔥",
		"🍃",
		"🌴",
		"💧",
		"👼",
		"👿",
	}

	min := 10
	for i := 0; i < numberOfCards; i++ {
		card := NewBase(attr)
		// sets status
		card.Cost.Set(float64(r.Intn(10) + 1))
		card.Attack.Set(float64(r.Intn(10) + min))
		card.Defence.Set(float64(r.Intn(10) + min))
		card.Magic.Set(float64(r.Intn(10) + min))
		card.MagicDefence.Set(float64(r.Intn(10) + min))
		// sets atteributes
		card.Attr.Set(r.Intn(len(attr)))
		card.Effect.Descript.Set(fmt.Sprintf("飛行を持つ%.f/%.fの%sトークンをX体生成", card.Attack.Value(), card.Defence.Value(), card.Attributes[card.Attr.Value()]))
		// stacks cards
		cards[i] = card
	}

	for i := range cards {
		fmt.Printf("%+v\n\n\n", cards[i].Status())
	}
}
