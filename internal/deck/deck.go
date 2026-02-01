package deck

import (
	"fmt"
	"math/rand"
)

type Deck struct {
	Cards []Card
}

func (d *Deck) Length() int {
	return len(d.Cards)
}

func (d *Deck) AddTop(cards ...Card) {
	d.Cards = append(d.Cards, cards...)
}

func (d *Deck) AddBottom(cards ...Card) {
	d.Cards = append(cards, d.Cards...)
}

func (d *Deck) Draw() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, fmt.Errorf("No cards left in the deck")
	}
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card, nil
}

func (d *Deck) Peek() (Card, error) {
	if len(d.Cards) == 0 {
		return Card{}, fmt.Errorf("No cards left in the deck")
	}
	card := d.Cards[len(d.Cards)-1]
	return card, nil
}

func (d Deck) String() string {
	return fmt.Sprintf("Deck{Cards: [%s]}", string(d.Fronts()))
}

func (d *Deck) Fronts() []rune {
	cardFronts := make([]rune, len(d.Cards))
	for i, card := range d.Cards {
		cardFronts[i] = card.Front()
	}
	return cardFronts
}

func New() Deck {
	cards := make([]Card, 0, 52)
	for suit := Spades; suit <= Clubs; suit++ {
		for rank := Ace; rank <= King; rank++ {
			cards = append(cards, Card{
				Rank:  rank,
				Suit:  suit,
				Color: (suit == Hearts || suit == Diamonds),
			})
		}
	}
	return Deck{Cards: cards}
}

func Empty() Deck {
	return Deck{Cards: []Card{}}
}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
