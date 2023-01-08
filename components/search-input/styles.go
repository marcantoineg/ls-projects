package searchinput

import (
	projectform "list-my-projects/components/project-form"

	"github.com/charmbracelet/lipgloss"
)

type SearchInputStyles struct {
	ContainerStyle        lipgloss.Style
	FocusedContainerStyle lipgloss.Style
	InputCursorStyle      lipgloss.Style
}

var Style = SearchInputStyles{
	ContainerStyle:        lipgloss.NewStyle().Padding(0, 1).MarginLeft(4).Border(lipgloss.NormalBorder()),
	FocusedContainerStyle: lipgloss.NewStyle().Padding(0, 1).MarginLeft(4).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#6C91BF")),
	InputCursorStyle:      projectform.Style.NewFocusedStyle.Copy(),
}
