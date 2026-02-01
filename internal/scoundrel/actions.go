package scoundrel

import (
	"fmt"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
)

type Action struct {
	Description string
	CardIdx     int
	callback    func(*Game, int) error
}

func play_again(game *Game, _ int) error {
	*game = New()
	return nil
}

func run_away(game *Game, _ int) error {
	game.Discard.AddBottom(game.Room...)
	game.Room = []deck.Card{}
	game.BuildRoom()
	game.CanRun = false
	return nil
}

func discard_potion(game *Game, i int) error {
	if i >= len(game.Room) {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := game.Room[i]
	if card.Suit != deck.Hearts {
		return fmt.Errorf("Card at Room[%d] is not a potion", i)
	}
	game.Discard.AddTop(card)
	game.Room = append(game.Room[:i], game.Room[i+1:]...)
	game.CanRun = false
	return nil
}

func drink_potion(game *Game, i int) error {
	if i >= len(game.Room) {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := game.Room[i]
	if card.Suit != deck.Hearts {
		return fmt.Errorf("Card at Room[%d] is not a potion", i)
	}
	game.Health = min(game.Health+card.Value(), 20)
	game.HasDrankPotion = true
	game.Discard.AddTop(card)
	game.Room = append(game.Room[:i], game.Room[i+1:]...)
	game.CanRun = false
	return nil
}

func take_weapon(game *Game, i int) error {
	if i >= len(game.Room) {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := game.Room[i]
	if card.Suit != deck.Diamonds {
		return fmt.Errorf("Card at Room[%d] is not a weapon", i)
	}
	if game.Weapon != nil {
		game.Discard.AddTop(game.Weapon.Weapon)
		game.Discard.AddTop(game.Weapon.Killed.Cards...)
		game.Weapon = nil
	}
	game.Weapon = &Weapon{Weapon: card, Killed: deck.Empty()}
	game.Room = append(game.Room[:i], game.Room[i+1:]...)
	game.CanRun = false
	return nil
}

func fight_monster_barehanded(game *Game, i int) error {
	if i >= len(game.Room) {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := game.Room[i]
	if card.Color != deck.Black {
		return fmt.Errorf("Card at Room[%d] is not a monster", i)
	}
	game.Health -= card.Value()
	game.Discard.AddTop(card)
	game.Room = append(game.Room[:i], game.Room[i+1:]...)
	game.CanRun = false
	return nil
}

func fight_monster_with_weapon(game *Game, i int) error {
	if i >= len(game.Room) {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := game.Room[i]
	if card.Color != deck.Black {
		return fmt.Errorf("Card at Room[%d] is not a monster", i)
	}
	if game.Weapon == nil {
		return fmt.Errorf("You don't have a weapon")
	}
	if last, err := game.Weapon.Killed.Peek(); err == nil && last.Value() <= card.Value() {
		return fmt.Errorf("You can't fight %s with your current weapon.", card)
	}
	game.Health -= max(0, card.Value()-game.Weapon.Weapon.Value())
	game.Weapon.Killed.AddTop(card)
	game.Room = append(game.Room[:i], game.Room[i+1:]...)
	game.CanRun = false
	return nil
}
