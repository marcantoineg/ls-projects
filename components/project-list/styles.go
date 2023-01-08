package projectlist

import (
	"list-my-projects/components/styles"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight       = 14
	listWidth        = 20
	listInitialTitle = "Navigate to ...?"
)

var Style = struct {
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
	TitleStyle:           styles.BaseTitle().Background(lipgloss.Color("#6C91BF")),
	SuccessTitleStyle:    styles.BaseTitle().Background(lipgloss.Color("#25A065")),
	ErrorTitleStyle:      styles.BaseTitle().Background(lipgloss.Color("#E84855")),
	MovingModeTitleStyle: styles.BaseTitle().Background(lipgloss.Color("#4d4d4d")),
	NoItemsStyle:         list.DefaultStyles().NoItems.MarginLeft(4),
	PaginationStyle:      list.DefaultStyles().PaginationStyle.PaddingLeft(4),
	HelpStyle:            list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1),
	QuitTextStyle:        lipgloss.NewStyle().Margin(1, 2),
	FatalErrorStyle:      lipgloss.NewStyle().Margin(1, 2).Foreground(lipgloss.Color("#E84855")),
	PathTextStyle:        lipgloss.NewStyle().Foreground(lipgloss.Color("170")),
}
