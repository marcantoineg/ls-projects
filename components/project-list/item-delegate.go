package projectlist

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type itemDelegate struct {
	movingModeInitialIndex int
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
		fn = func(strs ...string) string {
			strs = append([]string{">"}, strs...)
			return selectedItemStyle.Render(strs...)
		}
	} else if index == d.movingModeInitialIndex {
		fn = func(strs ...string) string {
			strs = append([]string{"*"}, strs...)
			return selectedItemForMovingStyle.Render(strs...)
		}
	}

	//lint:ignore SA1006 expected behavior
	fmt.Fprintf(w, fn(str))
}
