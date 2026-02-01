package scoundrel

import (
	"fmt"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
)

type Action struct {
	Description string
	CardIdx     int
	Effect      Effect
	callback    func(*Game, int) error
}

type Effect int

const (
	PlayAgain = iota + 1
	DiscardPotion
	DrinkPotion
	TakeWeapon
	AttackWithWeapon
	AttackBarehanded
)

func playAgain(game *Game, _ int) error {
	*game = New()
	return nil
}

func runAway(game *Game, _ int) error {
	cards := []deck.Card{}
	for _, card := range game.Room {
		if card != nil {
			cards = append(cards, *card)
		}
	}
	game.Discard.AddBottom(cards...)
	game.Room = [4]*deck.Card{}
	game.BuildRoom()
	game.CanRun = false
	return nil
}

func discardPotion(game *Game, i int) error {
	if game.Room[i] == nil {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := *game.Room[i]
	if card.Suit != deck.Hearts {
		return fmt.Errorf("Card at Room[%d] is not a potion", i)
	}
	game.Discard.AddTop(card)
	game.Room[i] = nil
	game.CanRun = false
	return nil
}

func drinkPotion(game *Game, i int) error {
	if game.Room[i] == nil {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := *game.Room[i]
	if card.Suit != deck.Hearts {
		return fmt.Errorf("Card at Room[%d] is not a potion", i)
	}
	game.Health = min(game.Health+card.Value(), 20)
	game.HasDrankPotion = true
	game.Discard.AddTop(card)
	game.Room[i] = nil
	game.CanRun = false
	return nil
}

func takeWeapon(game *Game, i int) error {
	if game.Room[i] == nil {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := *game.Room[i]
	if card.Suit != deck.Diamonds {
		return fmt.Errorf("Card at Room[%d] is not a weapon", i)
	}
	if game.Weapon != nil {
		game.Discard.AddTop(game.Weapon.Weapon)
		game.Discard.AddTop(game.Weapon.Killed.Cards...)
		game.Weapon = nil
	}
	game.Weapon = &Weapon{Weapon: card, Killed: deck.Empty("Killed")}
	game.Room[i] = nil
	game.CanRun = false
	return nil
}

func fightMonsterBarehanded(game *Game, i int) error {
	if game.Room[i] == nil {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := *game.Room[i]
	if card.Color != deck.Black {
		return fmt.Errorf("Card at Room[%d] is not a monster", i)
	}
	game.Health -= card.Value()
	game.Discard.AddTop(card)
	game.Room[i] = nil
	game.CanRun = false
	return nil
}

func fightMonsterWithWeapon(game *Game, i int) error {
	if game.Room[i] == nil {
		return fmt.Errorf("No card at Room[%d]", i)
	}
	card := *game.Room[i]
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
	game.Room[i] = nil
	game.CanRun = false
	return nil
}
