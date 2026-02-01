package scoundrel

import (
	"testing"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
)

func TestScore(t *testing.T) {
	cases := []struct {
		input    Game
		expected int
	}{
		{
			input: Game{
				Health:   20,
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Hearts, Color: deck.Red}},
				GameOver: true,
			},
			expected: 25,
		},
		{
			input: Game{
				Health:   20,
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Hearts, Color: deck.Red}},
				GameOver: false,
			},
			expected: 0,
		},
		{
			input: Game{
				Health:   2,
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Hearts, Color: deck.Red}},
				GameOver: true,
			},
			expected: 2,
		},
		{
			input: Game{
				Health:   20,
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Clubs, Color: deck.Black}},
				GameOver: true,
			},
			expected: 20,
		},
		{
			input: Game{
				Health: 0,
				Dungeon: deck.Deck{Cards: []deck.Card{
					{Rank: deck.Ace, Suit: deck.Clubs, Color: deck.Black, AceHigh: true},
					{Rank: deck.Ace, Suit: deck.Hearts, Color: deck.Red, AceHigh: true},
					{Rank: deck.Ten, Suit: deck.Spades, Color: deck.Black, AceHigh: true},
					{Rank: deck.Six, Suit: deck.Clubs, Color: deck.Black, AceHigh: true},
				}},
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Clubs, Color: deck.Black}},
				GameOver: true,
			},
			expected: -30,
		},
		{
			input: Game{
				Health: -5,
				Dungeon: deck.Deck{Cards: []deck.Card{
					{Rank: deck.Ace, Suit: deck.Clubs, Color: deck.Black, AceHigh: true},
					{Rank: deck.Ace, Suit: deck.Hearts, Color: deck.Red, AceHigh: true},
					{Rank: deck.Ten, Suit: deck.Spades, Color: deck.Black, AceHigh: true},
					{Rank: deck.Six, Suit: deck.Clubs, Color: deck.Black, AceHigh: true},
				}},
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Clubs, Color: deck.Black}},
				GameOver: true,
			},
			expected: -30,
		},
		{
			input: Game{
				Health:   0,
				Dungeon:  deck.Empty(),
				Room:     []deck.Card{{Rank: deck.Five, Suit: deck.Clubs, Color: deck.Black}},
				GameOver: true,
			},
			expected: 0,
		},
	}

	for _, c := range cases {
		actual := c.input.Score()
		if actual != c.expected {
			t.Errorf("Expected %d but got %d", c.expected, actual)
		}
	}
}
