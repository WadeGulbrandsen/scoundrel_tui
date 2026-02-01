package tui

import (
	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
	"github.com/charmbracelet/lipgloss"
)

var deckStyle = gameStyle.
	Width(9).
	Height(7).
	Margin(1)

func renderDeck(d *deck.Deck) string {
	content := ""
	if d.Length() > 0 {
		content = renderCardBack()
	}
	title := d.Name
	if len(title) > 9 {
		title = string([]rune(title)[:9])
	}
	top := lipgloss.PlaceHorizontal(9, lipgloss.Center, title, lipgloss.WithWhitespaceChars("═"))
	boarder := lipgloss.Border{
		Top:         top,
		Bottom:      "═",
		Left:        "║",
		Right:       "║",
		TopLeft:     "╔",
		TopRight:    "╗",
		BottomLeft:  "╚",
		BottomRight: "╝",
	}
	return deckStyle.Border(boarder).Render(content)
}
