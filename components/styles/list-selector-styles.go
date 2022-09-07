package styles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	ListHeight       = 14
	ListWidth        = 20
	ListInitialTitle = "Navigate to ...?"
)

var ListSelector = struct {
	TitleStyle           lipgloss.Style
	SuccessTitleStyle    lipgloss.Style
	ErrorTitleStyle      lipgloss.Style
	MovingModeTitleStyle lipgloss.Style
	NoItemsStyle         lipgloss.Style
	PaginationStyle      lipgloss.Style
	HelpStyle            lipgloss.Style
	QuitTextStyle        lipgloss.Style
	FatalErrorStyle      lipgloss.Style
	PathTextStyle        lipgloss.Style
}{
	TitleStyle:           baseTitle().Background(lipgloss.Color("#6C91BF")),
	SuccessTitleStyle:    baseTitle().Background(lipgloss.Color("#25A065")),
	ErrorTitleStyle:      baseTitle().Background(lipgloss.Color("#E84855")),
	MovingModeTitleStyle: baseTitle().Background(lipgloss.Color("#4d4d4d")),
	NoItemsStyle:         list.DefaultStyles().NoItems.MarginLeft(4),
	PaginationStyle:      list.DefaultStyles().PaginationStyle.PaddingLeft(4),
	HelpStyle:            list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1),
	QuitTextStyle:        lipgloss.NewStyle().Margin(1, 2),
	FatalErrorStyle:      lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("#E84855")),
	PathTextStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("170")),
}
