package components

import (
	"fmt"
	"os/exec"
	"strings"

	"list-my-projects/components/styles"
	"list-my-projects/fileutil"
	"list-my-projects/models/project"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type listSelectorModel struct {
	list                   list.Model
	items                  []list.Item
	choice                 *project.Project
	projectForm            *projectFormModel
	fatalError             error
	movingModeActive       bool
	movingModeInitialIndex int
	quitting               bool
	searchInput            *searchInputModel
	typingSearchTerm       bool
}

func NewListSelector() tea.Model {
	l := list.New([]list.Item{}, itemDelegate{movingModeInitialIndex: -1}, styles.ListWidth, styles.ListHeight)

	l.Title = styles.ListInitialTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.Styles.Title = styles.ListSelector.TitleStyle
	l.Styles.NoItems = styles.ListSelector.NoItemsStyle
	l.Styles.PaginationStyle = styles.ListSelector.PaginationStyle
	l.Styles.HelpStyle = styles.ListSelector.HelpStyle

	l.KeyMap.NextPage = key.NewBinding()
	l.KeyMap.PrevPage = key.NewBinding()
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("enter", "space"), key.WithHelp("enter/space", "select a project")),
		}
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "add a project")),
			key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit selected project")),
			key.NewBinding(key.WithKeys("d"), key.WithHelp("d", "delete selected project")),
			key.NewBinding(key.WithKeys("f"), key.WithHelp("f", "search projects")),
			key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "yank selected project's path")),
			key.NewBinding(key.WithKeys("m"), key.WithHelp("m", "enter moving mode")),
		}
	}

	m := listSelectorModel{list: l}
	return m
}

type initMsg struct{ items []list.Item }

func (m listSelectorModel) Init() tea.Cmd {
	projects, err := project.GetAll()
	if err != nil {
		return func() tea.Msg { return fatalErrorMsg{err} }
	}

	return func() tea.Msg { return initMsg{castToListItem(projects)} }
}

func (m listSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd = nil

	switch msg := msg.(type) {
	case fatalErrorMsg:
		m.fatalError = msg.err
		m.projectForm = nil
		m.quitting = true

		return m, tea.Quit

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case initMsg:
		m.items = msg.items
		m.list.SetItems(m.items)

	case projectCreatedMsg:
		projects, err := project.Save(m.list.Index(), msg.project)
		if err != nil {
			m.Update(projectCreationErrorMsg(err))
			return m, nil
		}

		m.items = castToListItem(projects)
		m.list.SetItems(m.items)

		m.list.Styles.Title = styles.ListSelector.SuccessTitleStyle
		m.list.Title = fmt.Sprintf("project '%s' added!", msg.project.Name)

		m.projectForm = nil

	case projectUpdatedMsg:
		projects, err := project.Update(m.list.Index(), msg.project)
		if err != nil {
			m.Update(projectUpdateErrorMsg(err))
			return m, nil
		}

		m.items = castToListItem(projects)
		m.list.SetItems(m.items)

		m.list.Styles.Title = styles.ListSelector.SuccessTitleStyle
		m.list.Title = fmt.Sprintf("project '%s' updated!", msg.project.Name)

		m.projectForm = nil

	case noProjectCreatedMsg:
		resetListTitle(&m)
		m.projectForm = nil

	case cancelSearch:
		m.typingSearchTerm = false
		m.searchInput = nil
		m.filterList(nil)

	case updateSearch:
		m.filterList(msg.filteredItemsIndices)

	case submitSearch:
		m.typingSearchTerm = false
		m.filterList(msg.filteredItemsIndices)

	// Keybinding
	case tea.KeyMsg:
		if m.typingSearchTerm && m.searchInput != nil {
			model, cmd := m.searchInput.Update(msg)
			searchModel := model.(searchInputModel)
			m.searchInput = &searchModel
			return m, cmd
		} else {
			return handleListkeybinds(&m, msg)
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// handleKeyMsg handles the keybinding part of the Update function.
func handleListkeybinds(m *listSelectorModel, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch keypress := msg.String(); keypress {
	case "ctrl+c", "q", "esc":
		if m.movingModeActive {
			disableMovingMode(m)
			return m, nil
		}

		m.quitting = true
		return m, tea.Quit

	case "enter", "space":
		if len(m.list.Items()) == 0 {
			return m, nil
		}
		if !m.movingModeActive {
			selectedItem := m.list.SelectedItem().(project.Project)
			m.choice = &selectedItem

			cmd := exec.Command("code", "-n", ".")
			cmd.Dir = fileutil.ReplaceTilde(m.choice.Path)

			err := cmd.Run()
			if err != nil {
				return m, func() tea.Msg { return fatalErrorMsg{err} }
			}

			return m, tea.Quit
		} else {
			projects, err := project.SwapIndex(m.movingModeInitialIndex, m.list.Index())
			if err != nil {
				m.Update(projectUpdateErrorMsg(err))
				return m, nil
			}

			m.items = castToListItem(projects)
			m.list.SetItems(m.items)

			disableMovingMode(m)
		}

	case "a":
		if !m.movingModeActive {
			f := NewProjectForm(m, nil)
			m.projectForm = &f

			return m.projectForm.Update(nil)
		}

	case "e":
		if !m.movingModeActive {
			if p, ok := m.items[m.list.Index()].(project.Project); ok {
				f := NewProjectForm(m, &p)
				m.projectForm = &f
			}

			return m.projectForm.Update(nil)
		}

	case "d":
		if !m.movingModeActive {
			if p, ok := m.list.SelectedItem().(project.Project); ok {
				projects, err := project.Delete(m.list.Index(), p)
				if err != nil {
					m.list.Styles.Title = styles.ListSelector.ErrorTitleStyle
					m.list.Title = fmt.Sprintf("error deleting project '%s'", p.Name)
				}

				m.items = castToListItem(projects)
				cmd := m.list.SetItems(m.items)

				m.list.Styles.Title = styles.ListSelector.SuccessTitleStyle
				m.list.Title = fmt.Sprintf("project '%s' deleted", p.Name)

				return m, cmd
			}
		}

	case "y":
		if !clipboard.Unsupported && !m.movingModeActive {
			if p, ok := m.list.SelectedItem().(project.Project); ok {
				clipboard.WriteAll(p.Path)

				m.list.Styles.Title = styles.ListSelector.SuccessTitleStyle
				m.list.Title = fmt.Sprintf("path for project '%s' copied", p.Name)

				return m, nil
			}
		}

	case "m":
		m.movingModeInitialIndex = m.list.Index()
		m.list.SetDelegate(itemDelegate{movingModeInitialIndex: m.movingModeInitialIndex})
		m.movingModeActive = true

		m.list.Styles.Title = styles.ListSelector.MovingModeTitleStyle
		m.list.Title = "select another project to swap position"

	case "f":
		if m.searchInput == nil {
			projectNames := make([]string, len(m.items))
			for i, p := range m.items {
				projectNames[i] = p.(project.Project).Name
			}

			s := newSearchInput(projectNames)
			m.searchInput = &s
			m.searchInput.Focus()
			m.typingSearchTerm = true
			m.searchInput.Update(nil)
		} else {
			m.typingSearchTerm = true
			m.searchInput.Focus()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m listSelectorModel) View() string {
	if m.fatalError != nil {
		return styles.ListSelector.FatalErrorStyle.Render(
			fmt.Sprintf("fatal error:\n\n%s", m.fatalError),
		)
	}

	if m.choice != nil {
		return styles.ListSelector.QuitTextStyle.Render(
			fmt.Sprintf("Opening %s 💻", styles.ListSelector.PathTextStyle.Render(m.choice.Path)),
		)
	}

	if m.quitting {
		return styles.ListSelector.QuitTextStyle.Render("mmmhhhh-kay.")
	}

	if m.projectForm != nil {
		return m.projectForm.View()
	}

	var sb strings.Builder

	if m.searchInput != nil {
		sb.WriteString(m.searchInput.View() + "\n")
	}

	sb.WriteString("\n" + m.list.View())

	return sb.String()
}

// castToListItem takes a list of 'Project's and returns it as a casted list of tea's interface 'list.Item'.
func castToListItem(projects []project.Project) []list.Item {
	castedItems := make([]list.Item, len(projects))
	for i, p := range projects {
		castedItems[i] = p
	}
	return castedItems
}

// resetListTitle resets the initial style and text of the list's title.
func resetListTitle(m *listSelectorModel) {
	m.list.Styles.Title = styles.ListSelector.TitleStyle
	m.list.Title = styles.ListInitialTitle
}

// disableMovingMode resets required value to disable the moving mode.
func disableMovingMode(m *listSelectorModel) {
	m.movingModeInitialIndex = -1
	m.movingModeActive = false
	m.list.SetDelegate(itemDelegate{movingModeInitialIndex: -1})
	resetListTitle(m)
}

// filterList filters the items in the list (m.list) given a list of indices.
func (m *listSelectorModel) filterList(filteredIndices []int) {
	if filteredIndices == nil {
		m.list.SetItems(m.items)
		return
	}

	filteredItems := make([]list.Item, len(filteredIndices))
	for i, itemIndex := range filteredIndices {
		if itemIndex <= len(m.items) {
			filteredItems[i] = m.items[itemIndex]
		}
	}
	m.list.SetItems(filteredItems)
}

type fatalErrorMsg struct {
	err error
}
