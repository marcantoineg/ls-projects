package styles

import "github.com/charmbracelet/lipgloss"

type SearchInputStyles struct {
	ContainerStyle        lipgloss.Style
	FocusedContainerStyle lipgloss.Style
	InputCursorStyle      lipgloss.Style
}

var SearchInput = SearchInputStyles{
	ContainerStyle:        lipgloss.NewStyle().Padding(0, 1).MarginLeft(4).Border(lipgloss.NormalBorder()),
	FocusedContainerStyle: lipgloss.NewStyle().Padding(0, 1).MarginLeft(4).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#6C91BF")),
	InputCursorStyle:      Form.NewFocusedStyle.Copy(),
}
