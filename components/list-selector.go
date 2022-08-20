package components

import (
	"fmt"
	"os/exec"

	models "list-my-projects/models"
	fileutils "list-my-projects/utils"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listHeight   = 14
	listWidth    = 20
	initialTitle = "Navigate to ...?"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#6C91BF")).
			MarginLeft(2).
			Padding(0, 1)

	successTitleStyle = titleStyle.Copy().
				Background(lipgloss.Color("#25A065"))

	errorTitleStyle = titleStyle.Copy().
			Background(lipgloss.Color("#E84855"))

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	selectedItemStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6C91BF")).
				PaddingLeft(2)

	noItemsStyle = list.DefaultStyles().NoItems.
			MarginLeft(4)

	paginationStyle = list.DefaultStyles().PaginationStyle.
			PaddingLeft(4)

	helpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1)

	quitTextStyle = lipgloss.NewStyle().
			Margin(1, 2)

	pathTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("170"))
)

type listSelectorModel struct {
	list        list.Model
	items       []list.Item
	choice      *models.Project
	projectForm *projectFormModel
	quitting    bool
}

func NewListSelector() tea.Model {
	l := list.New([]list.Item{}, itemDelegate{}, listWidth, listHeight)

	l.Title = initialTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.Styles.Title = titleStyle
	l.Styles.NoItems = noItemsStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add a project")),
			key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete selected project")),
		}
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("enter", "space"), key.WithHelp("enter/space", "select a project")),
		}
	}

	m := listSelectorModel{list: l}
	return m
}

type initMsg struct{ items []list.Item }

func (m listSelectorModel) Init() tea.Cmd {
	projects, err := models.GetProjects()
	if err != nil {
		panic(err)
	}

	return func() tea.Msg { return initMsg{castToListItem(projects)} }
}

func (m listSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case initMsg:
		m.items = msg.items
		m.list.SetItems(m.items)

	case projectCreatedMsg:
		_, err := models.SaveProject(msg.project)
		if err != nil {
			m.Update(projectCreationErrorMsg(err))
			return m, nil
		}

		m.items = append(m.items, msg.project)
		m.list.SetItems(m.items)

		m.list.Styles.Title = successTitleStyle
		m.list.Title = fmt.Sprintf("project '%s' added!", msg.project.Name)

		m.projectForm = nil

	case noProjectCreatedMsg:
		m.list.Styles.Title = titleStyle
		m.list.Title = initialTitle

		m.projectForm = nil

	case projectCreationErrorMsg:
		m.list.Styles.Title = errorTitleStyle
		m.list.Title = msg.Error()

		m.projectForm = nil

	// Keybinding
	case tea.KeyMsg:
		return handleKeyMsg(&m, msg)
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// handleKeyMsg handles the keybinding part of the Update function.
func handleKeyMsg(m *listSelectorModel, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch keypress := msg.String(); keypress {
	case "ctrl+c", "q", "esc":
		m.quitting = true
		return m, tea.Quit

	case "enter", "space":
		selectedItem := m.list.SelectedItem().(models.Project)
		m.choice = &selectedItem
		m.quitting = true
		return m, tea.Quit

	case "a":
		f := NewProjectForm(m)
		m.projectForm = &f

		model, cmd := m.projectForm.Update(nil)
		return model, cmd

	case "d":
		p, ok := m.list.SelectedItem().(models.Project)
		if ok {
			selectedIndex := m.list.Index()
			projects, err := models.DeleteProject(selectedIndex, p)
			if err != nil {
				m.list.Styles.Title = errorTitleStyle
				m.list.Title = fmt.Sprintf("error deleting project '%s'", p.Name)
			}

			cmd := m.list.SetItems(castToListItem(projects))

			m.list.Styles.Title = successTitleStyle
			m.list.Title = fmt.Sprintf("project '%s' deleted", p.Name)

			return m, cmd
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listSelectorModel) View() string {
	if m.choice != nil {
		cmd := exec.Command("code", "-n", ".")

		cmd.Dir = fileutils.ReplaceTilde(m.choice.Path)
		err := cmd.Run()
		if err != nil {
			return quitTextStyle.Render(err.Error())
		}
		return quitTextStyle.Render(fmt.Sprintf("Opening %s 💻", pathTextStyle.Render(m.choice.Path)))
	}

	if m.quitting {
		return quitTextStyle.Render("mmmhhhh-kay.")
	}

	if m.projectForm != nil {
		return m.projectForm.View()
	}

	return "\n" + m.list.View()
}

// castToListItem takes a list of 'Project's and returns it as a casted list of tea's interface 'list.Item'.
func castToListItem(projects []models.Project) []list.Item {
	castedItems := make([]list.Item, len(projects))
	for i, p := range projects {
		castedItems[i] = p
	}
	return castedItems
}
