package tui

import (
	"strings"

	"github.com/WadeGulbrandsen/scoundrel_tui/internal/deck"
	"github.com/charmbracelet/lipgloss"
)

var cardStyle = lipgloss.NewStyle().
	Width(7).
	Height(5).
	Margin(0).
	BorderStyle(lipgloss.OuterHalfBlockBorder()).
	BorderForeground(lipgloss.Color("223")).
	BorderBackground(lipgloss.Color("188")).
	Background(lipgloss.Color("188"))

var redCardStyle = cardStyle.
	BorderForeground(lipgloss.Color("160")).
	Foreground(lipgloss.Color("160"))

var blkCardStyle = cardStyle.
	BorderForeground(lipgloss.Color("16")).
	Foreground(lipgloss.Color("16"))

var cardBackStyle = cardStyle.
	BorderForeground(lipgloss.Color("17")).
	Foreground(lipgloss.Color("62")).
	BorderBackground(lipgloss.Color("19")).
	Background(lipgloss.Color("19"))

const cardBackOddRow = "XYXYXYX"
const cardBackEvenRow = "YXYXYXY"

var cardFrontsTemplates = [...]string{
	"J     R\nO     E\nK  X  K\nE     O\nR     J",
	"A      \n       \n   X   \n       \n      A",
	"2      \n   X   \n       \n   X   \n      2",
	"3      \n   X   \n   X   \n   X   \n      3",
	"4      \n  X X  \n       \n  X X  \n      4",
	"5      \n  X X  \n   X   \n  X X  \n      5",
	"6      \n  X X  \n  X X  \n  X X  \n      6",
	"7 X X  \n   X   \n  X X  \n  X X  \n      7",
	"8 X X  \n   X   \n  X X  \n   X   \n  X X 8",
	"9 X X  \n  X X  \n   X   \n  X X  \n  X X 9",
	"10X X  \n  XXX  \n  X X  \n   X   \n  X X10",
	"J      \nX J A  \n       \n  C K X\n      J",
	"Q      \nX Q U  \n   E   \n  E E X\n      Q",
	"K      \nX K I  \n       \n  N G X\n      K",
}

func renderCardFront(c *deck.Card) string {
	suit := c.Suit.String()
	if c.Suit < 1 || c.Suit > 4 {
		suit = "\U0001F7CD"
	}
	template := "UNKNOWN\n       \n       \n       \nUNKNOWN"
	if int(c.Rank) >= 0 && int(c.Rank) <= len(cardFrontsTemplates) {
		template = cardFrontsTemplates[int(c.Rank)]
	}
	cardString := strings.ReplaceAll(template, "X", suit)
	style := blkCardStyle
	if c.Color == deck.Red {
		style = redCardStyle
	}
	return style.Render(cardString)
}

func renderCardBack() string {
	cardStrings := []string{}
	x := "░"
	y := "▒"
	evens := strings.ReplaceAll(strings.ReplaceAll(cardBackEvenRow, "X", x), "Y", y)
	odds := strings.ReplaceAll(strings.ReplaceAll(cardBackOddRow, "X", x), "Y", y)
	for i := range 5 {
		if i%2 == 0 {
			cardStrings = append(cardStrings, evens)
		} else {
			cardStrings = append(cardStrings, odds)
		}
	}
	return cardBackStyle.Render(strings.Join(cardStrings, "\n"))
}
