package deck

import "testing"

func TestDeckSuffle(t *testing.T) {
	deck := New()
	original := deck.String()
	deck.Shuffle()
	shuffled := deck.String()
	if original == shuffled {
		t.Errorf("Deck is still in the same order")
	}
}
