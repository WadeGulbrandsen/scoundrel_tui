package scoundrel

import (
	"fmt"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
)

type Game struct {
	Health         int
	Weapon         *Weapon
	CanRun         bool
	Dungeon        deck.Deck
	Room           []deck.Card
	Discard        deck.Deck
	HasDrankPotion bool
	GameOver       bool
}

func (sg *Game) Score() int {
	switch {
	case !sg.GameOver:
		return 0
	case sg.Health <= 0:
		score := 0
		for _, card := range sg.Dungeon.Cards {
			if card.Color == deck.Black {
				score -= card.Value()
			}
		}
		return score
	case sg.Health == 20 && len(sg.Room) == 1 && sg.Room[0].Suit == deck.Hearts:
		return 20 + sg.Room[0].Value()
	default:
		return sg.Health
	}
}

func (sg *Game) BuildRoom() {
	for i := len(sg.Room); i < 4; i++ {
		card, err := sg.Dungeon.Draw()
		if err != nil {
			// Deck is empty
			break
		}
		sg.Room = append(sg.Room, card)
	}
}

func (sg *Game) DoAction(a Action) error {
	return a.callback(sg, a.CardIdx)
}

func (sg *Game) GetActions() []Action {
	if sg.Health <= 0 {
		// Player died!
		sg.GameOver = true
		return []Action{{
			Description: "Play again",
			callback:    play_again,
		}}
	}

	if len(sg.Room) <= 1 {
		// Current turn complet
		sg.HasDrankPotion = false
		sg.CanRun = true
		sg.BuildRoom()
	}

	if len(sg.Room) <= 1 && sg.Dungeon.Length() == 0 {
		// Player won!
		sg.GameOver = true
		return []Action{{
			Description: "Play again",
			callback:    play_again,
		}}
	}

	max_weapon := 0
	if sg.Weapon != nil {
		if last, err := sg.Weapon.Killed.Peek(); err == nil {
			max_weapon = last.Value()
		} else {
			max_weapon = 99
		}
	}

	actions := []Action{}
	for i, card := range sg.Room {
		switch card.Suit {
		case deck.Hearts:
			if sg.HasDrankPotion {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Discard %s potion", card),
					CardIdx:     i,
					callback:    discard_potion,
				})
			} else {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Drink %s potion", card),
					CardIdx:     i,
					callback:    drink_potion,
				})
			}
		case deck.Diamonds:
			actions = append(actions, Action{
				Description: fmt.Sprintf("Take %s weapon", card),
				CardIdx:     i,
				callback:    take_weapon,
			})
		default:
			if card.Value() < max_weapon {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Attack %s with your weapon", card),
					CardIdx:     i,
					callback:    fight_monster_with_weapon,
				})
			}
			actions = append(actions, Action{
				Description: fmt.Sprintf("Attack %s barehanded", card),
				CardIdx:     i,
				callback:    fight_monster_barehanded,
			})
		}
	}
	if sg.CanRun {
		actions = append(actions, Action{
			Description: "Run away",
			callback:    run_away,
		})
	}
	return actions
}

func New() Game {
	cards := make([]deck.Card, 0, 44)
	for suit := deck.Spades; suit <= deck.Clubs; suit++ {
		for rank := deck.Ace; rank <= deck.King; rank++ {
			card := deck.Card{
				Rank:    rank,
				Suit:    suit,
				Color:   (suit == deck.Hearts || suit == deck.Diamonds),
				AceHigh: true,
			}
			if card.Color == deck.Red && (rank == deck.Ace || card.Rank.IsFace()) {
				continue
			}
			cards = append(cards, card)
		}
	}
	game := Game{
		Health:   20,
		CanRun:   true,
		Dungeon:  deck.Deck{Cards: cards},
		Discard:  deck.Empty(),
		GameOver: false,
	}
	game.Dungeon.Shuffle()
	game.BuildRoom()
	return game
}
