package components

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type itemDelegate struct {
	movingInitialIndex int
}

var (
	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6C91BF")).
				PaddingLeft(2)

	selectedItemForMovingStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("#4d4d4d")).
					PaddingLeft(2)
)

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	str := fmt.Sprintf("%d. %s", index+1, listItem.FilterValue())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	} else if index == d.movingInitialIndex {
		fn = func(s string) string {
			return selectedItemForMovingStyle.Render("* " + s)
		}
	}

	fmt.Fprintf(w, fn(str))
}
