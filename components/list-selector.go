package components

import (
	"fmt"
	"list-my-projects/models"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight = 14
	listWidth  = 20
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			MarginLeft(2).
			Padding(0, 1)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#25A065")).
				PaddingLeft(2)

	paginationStyle = list.DefaultStyles().PaginationStyle.
			PaddingLeft(4)

	helpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1)

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 0, 2, 4)

	pathTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("170"))
)

type model struct {
	list     list.Model
	items    []models.Project
	choice   *models.Project
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			selectedItem := m.list.SelectedItem().(models.Project)
			m.choice = &selectedItem
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != nil {
		cmd := exec.Command("code", "-n", ".")
		cmd.Dir = m.choice.Path
		err := cmd.Run()
		if err != nil {
			return quitTextStyle.Render(err.Error())
		}
		return quitTextStyle.Render(fmt.Sprintf("Opening %s 💻", pathTextStyle.Render(m.choice.Path)))
	}
	if m.quitting {
		return quitTextStyle.Render("mmmhhhh-kay.")
	}
	return "\n" + m.list.View()
}

func NewListSelector(items []models.Project) tea.Model {
	castedItems := make([]list.Item, len(items))
	for i := range items {
		castedItems[i] = items[i]
	}

	l := list.New(castedItems, itemDelegate{}, listWidth, listHeight)
	l.Title = "Navigate to ...?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := model{list: l}
	return m
}
