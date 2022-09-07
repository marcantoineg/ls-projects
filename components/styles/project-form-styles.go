package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

type ProjectFormStyles struct {
	NewProjectTitleStyle  lipgloss.Style
	EditProjectTitleStyle lipgloss.Style
	ErrorTitleStyle       lipgloss.Style
	NewFocusedStyle       lipgloss.Style
	EditFocusedStyle      lipgloss.Style
	BlurredStyle          lipgloss.Style
	CursorModeHelpStyle   lipgloss.Style
	FormHelpStyle         lipgloss.Style
	MarginStyle           lipgloss.Style
	NoStyle               lipgloss.Style
}

func (s ProjectFormStyles) BlurredButton() string {
	return fmt.Sprintf("[ %s ]", s.BlurredStyle.Render("Submit"))
}

var Form = ProjectFormStyles{
	NewProjectTitleStyle:  baseTitle().MarginLeft(0).Background(lipgloss.Color("#6C91BF")),
	EditProjectTitleStyle: baseTitle().MarginLeft(0).Background(lipgloss.Color("#DDB771")).Foreground(lipgloss.Color("#000000")),
	ErrorTitleStyle:       baseTitle().MarginLeft(0).Background(lipgloss.Color("#E84855")),
	NewFocusedStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF")),
	EditFocusedStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("#DDB771")),
	BlurredStyle:          lipgloss.NewStyle().Foreground(lipgloss.Color("240")),
	CursorModeHelpStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF")),
	FormHelpStyle:         lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true).Faint(true),
	MarginStyle:           lipgloss.NewStyle().MarginLeft(4),
	NoStyle:               lipgloss.NewStyle(),
}
