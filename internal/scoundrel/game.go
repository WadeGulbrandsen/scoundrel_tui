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
	Room           [4]*deck.Card
	Discard        deck.Deck
	HasDrankPotion bool
	GameOver       bool
}

func (game *Game) RemainingCardsInRoom() int {
	sum := 0
	for _, card := range game.Room {
		if card != nil {
			sum += 1
		}
	}
	return sum
}

func (game *Game) Score() int {
	remainingCards := []deck.Card{}
	for _, card := range game.Room {
		if card != nil {
			remainingCards = append(remainingCards, *card)
		}
	}
	switch {
	case !game.GameOver:
		return 0
	case game.Health <= 0:
		score := 0
		for _, card := range game.Dungeon.Cards {
			if card.Color == deck.Black {
				score -= card.Value()
			}
		}
		return score
	case game.Health == 20 &&
		len(remainingCards) == 1 &&
		remainingCards[0].Suit == deck.Hearts:
		return 20 + game.Room[0].Value()
	default:
		return game.Health
	}
}

func (game *Game) BuildRoom() {
	for i, card_ptr := range game.Room {
		if card_ptr != nil {
			// This slot in the room alread has a card
			continue
		}
		card, err := game.Dungeon.Draw()
		if err != nil {
			// Deck is empty
			break
		}
		game.Room[i] = &card
	}
}

func (game *Game) DoAction(a Action) error {
	return a.callback(game, a.CardIdx)
}

func (game *Game) GetActions() []Action {
	if game.Health <= 0 {
		// Player died!
		game.GameOver = true
		return []Action{{
			Description: "Play again",
			CardIdx:     -1,
			Effect:      PlayAgain,
			callback:    playAgain,
		}}
	}

	cards_in_room := game.RemainingCardsInRoom()

	if cards_in_room <= 1 && game.Dungeon.Length() == 0 {
		// Player won!
		game.GameOver = true
		return []Action{{
			Description: "Play again",
			CardIdx:     -1,
			callback:    playAgain,
		}}
	}

	if cards_in_room <= 1 {
		// Current turn complet
		game.HasDrankPotion = false
		game.CanRun = true
		game.BuildRoom()
	}

	max_weapon := 0
	if game.Weapon != nil {
		if last, err := game.Weapon.Killed.Peek(); err == nil {
			max_weapon = last.Value()
		} else {
			max_weapon = 99
		}
	}

	actions := []Action{}
	for i, card := range game.Room {
		if card == nil {
			// No card in this slot of the room
			continue
		}
		switch card.Suit {
		case deck.Hearts:
			if game.HasDrankPotion {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Discard %s potion", card),
					CardIdx:     i,
					Effect:      DiscardPotion,
					callback:    discardPotion,
				})
			} else {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Drink %s potion", card),
					CardIdx:     i,
					Effect:      DrinkPotion,
					callback:    drinkPotion,
				})
			}
		case deck.Diamonds:
			actions = append(actions, Action{
				Description: fmt.Sprintf("Take %s weapon", card),
				CardIdx:     i,
				Effect:      TakeWeapon,
				callback:    takeWeapon,
			})
		default:
			if card.Value() < max_weapon {
				actions = append(actions, Action{
					Description: fmt.Sprintf("Attack %s with your weapon", card),
					CardIdx:     i,
					Effect:      AttackWithWeapon,
					callback:    fightMonsterWithWeapon,
				})
			}
			actions = append(actions, Action{
				Description: fmt.Sprintf("Attack %s barehanded", card),
				CardIdx:     i,
				Effect:      AttackBarehanded,
				callback:    fightMonsterBarehanded,
			})
		}
	}
	if game.CanRun {
		actions = append(actions, Action{
			Description: "Run away",
			CardIdx:     -1,
			callback:    runAway,
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
		Dungeon:  deck.Deck{Cards: cards, Name: "Dungeon"},
		Discard:  deck.Empty("Discard"),
		GameOver: false,
	}
	game.Dungeon.Shuffle()
	game.BuildRoom()
	return game
}
