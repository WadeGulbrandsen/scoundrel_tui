package deck

import "fmt"

const CardBack = rune(0x1F0A0)

type Rank int

const (
	Joker Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (r Rank) String() string {
	if !r.IsRank() {
		return "Unknown"
	}
	return [...]string{
		"Joker",
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}[r]
}

func (r Rank) IsRank() bool {
	return r >= Joker && r <= King
}

func (r Rank) IsFace() bool {
	return r >= Jack && r <= King
}

type Suit int

const (
	None Suit = iota
	Spades
	Hearts
	Diamonds
	Clubs
)

func (s Suit) String() string {
	if s.IsSuit() {
		return [...]string{"None", "♠", "♥", "♦", "♣"}[s]
	}
	return "Unknown"
}

func (s Suit) IsSuit() bool {
	return s >= None && s <= Clubs
}

type Color bool

const (
	Black Color = false
	Red   Color = true
)

func (c Color) String() string {
	if c {
		return "Red"
	}
	return "Black"
}

type Card struct {
	Rank    Rank
	Suit    Suit
	Color   Color
	AceHigh bool
}

func (c *Card) Value() int {
	if c.Rank == Ace && c.AceHigh {
		return 14
	}
	return int(c.Rank)
}

func (c *Card) IsValid() bool {
	switch c.Suit {
	case None:
		return c.Rank == Joker
	case Hearts, Diamonds:
		return c.Rank >= Ace && c.Rank <= King && c.Color == Red
	case Spades, Clubs:
		return c.Rank >= Ace && c.Rank <= King && c.Color == Black
	}
	return false
}

func (c Card) Front() rune {
	if !c.IsValid() {
		return rune(0xFFFD)
	}
	switch c.Rank {
	case Joker:
		return rune(0x1F0BF)
	default:
		rune := rune(0x1F090 + (int(c.Suit) * 0x10) + int(c.Rank))
		if c.Rank >= Queen {
			rune++
		}
		return rune
	}
}

func (c Card) String() string {
	if !c.IsValid() {
		return fmt.Sprintf("{Rank: %d, Suit: %d, Color: %t}", c.Rank, c.Suit, c.Color)
	}
	if c.Rank == Joker {
		return fmt.Sprintf("%s Joker", c.Color)
	}
	if c.Rank > Ace && c.Rank < Jack {
		return fmt.Sprintf("%d%s", c.Rank, c.Suit)
	}
	return fmt.Sprintf("%s%s", c.Rank.String()[:2], c.Suit)
}
