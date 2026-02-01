package deck

import "testing"

func TestRankNames(t *testing.T) {
	cases := []struct {
		input    int
		expected string
	}{
		{
			input:    0,
			expected: "Joker",
		},
		{
			input:    -1,
			expected: "Unknown",
		},
		{
			input:    15,
			expected: "Unknown",
		},
		{
			input:    1,
			expected: "Ace",
		},
		{
			input:    13,
			expected: "King",
		},
	}

	for _, c := range cases {
		actual := Rank(c.input).String()
		if actual != c.expected {
			t.Errorf("Expected %q but got %q", c.expected, actual)
		}
	}
}

func TestCardValues(t *testing.T) {
	cases := []struct {
		input    Card
		expected int
	}{
		{input: Card{Rank: Joker, Suit: None, Color: Black}, expected: 0},
		{input: Card{Rank: Ace, Suit: Hearts, Color: Red, AceHigh: false}, expected: 1},
		{input: Card{Rank: Ace, Suit: Hearts, Color: Red, AceHigh: true}, expected: 14},
		{input: Card{Rank: Five, Suit: Clubs, Color: Black}, expected: 5},
		{input: Card{Rank: Queen, Suit: Spades, Color: Black}, expected: 12},
	}

	for _, c := range cases {
		actual := c.input.Value()
		if actual != c.expected {
			t.Errorf("Expected %q but got %q", c.expected, actual)
		}
	}
}
