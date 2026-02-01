package tui

import "fmt"

func renderActions(m model) string {
	actions := m.game.GetActions()
	s := ""
	for i, action := range actions {
		cursor := "   "
		if m.cursor == i {
			cursor = " ->"
		}
		s += fmt.Sprintf("%s %s\n", cursor, action.Description)
	}
	return s
}

// type actions_for_card struct {
// 	card    *deck.Card
// 	actions []scoundrel.Action
// }

// func (m model) actionsByCard() []actions_for_card {
// 	actions_map := make(map[int][]scoundrel.Action)
// 	for _, action := range m.game.GetActions() {
// 		if _, ok := actions_map[action.CardIdx]; !ok {
// 			actions_map[action.CardIdx] = []scoundrel.Action{}
// 		}
// 		actions_map[action.CardIdx] = append(actions_map[action.CardIdx], action)
// 	}
// 	actions_by_card := []actions_for_card{}
// 	for i, card := range m.game.Room {
// 		if lst, ok := actions_map[i]; ok {
// 			actions_by_card = append(actions_by_card, actions_for_card{card: card, actions: lst})
// 		}
// 	}
// 	if lst, ok := actions_map[-1]; ok {
// 		actions_by_card = append(actions_by_card, actions_for_card{card: nil, actions: lst})
// 	}
// 	return actions_by_card
// }
