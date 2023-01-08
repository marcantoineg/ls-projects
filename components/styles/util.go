package styles

import "github.com/charmbracelet/lipgloss"

func BaseTitle() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFDF5")).
		MarginLeft(2).
		Padding(0, 1)
}
