package projectlist

import (
	"fmt"
	"strings"

	projectform "list-my-projects/components/project-form"
	searchinput "list-my-projects/components/search-input"
	"list-my-projects/models/project"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list                   list.Model
	items                  []list.Item
	choice                 *project.Project
	projectForm            *projectform.Model
	fatalError             error
	movingModeActive       bool
	movingModeInitialIndex int
	quitting               bool
	searchInput            *searchinput.Model
	typingSearchTerm       bool
}

func NewProjectList() tea.Model {
	l := list.New([]list.Item{}, itemDelegate{movingModeInitialIndex: -1}, listWidth, listHeight)

	l.Title = listInitialTitle
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	l.Styles.Title = Style.TitleStyle
	l.Styles.NoItems = Style.NoItemsStyle
	l.Styles.PaginationStyle = Style.PaginationStyle
	l.Styles.HelpStyle = Style.HelpStyle

	l.KeyMap.NextPage = key.NewBinding()
	l.KeyMap.PrevPage = key.NewBinding()
	l.AdditionalShortHelpKeys = keybinds.defineShort
	l.AdditionalFullHelpKeys = keybinds.defineLong

	m := Model{list: l}
	return m
}

func (m Model) Init() tea.Cmd {
	projects, err := project.GetAll()
	if err != nil {
		return func() tea.Msg { return fatalErrorMsg{err} }
	}

	return func() tea.Msg { return initMsg{castToListItem(projects)} }
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

	case projectform.ProjectCreatedMsg:
		projects, err := project.Save(m.list.Index(), msg.Project)
		if err != nil {
			m.Update(projectform.ProjectCreationErrorMsg(err))
			return m, nil
		}

		m.items = castToListItem(projects)
		m.list.SetItems(m.items)

		m.list.Styles.Title = Style.SuccessTitleStyle
		m.list.Title = fmt.Sprintf("project '%s' added!", msg.Project.Name)

		m.projectForm = nil

	case projectform.ProjectUpdatedMsg:
		projects, err := project.Update(m.list.Index(), msg.Project)
		if err != nil {
			m.Update(projectform.ProjectUpdateErrorMsg(err))
			return m, nil
		}

		m.items = castToListItem(projects)
		m.list.SetItems(m.items)

		m.list.Styles.Title = Style.SuccessTitleStyle
		m.list.Title = fmt.Sprintf("project '%s' updated!", msg.Project.Name)

		m.projectForm = nil

	case projectform.NoProjectCreatedMsg:
		resetListTitle(&m)
		m.projectForm = nil

	case searchinput.CancelSearch:
		m.typingSearchTerm = false
		m.searchInput = nil
		m.filterList(nil)

	case searchinput.UpdateSearch:
		m.filterList(msg.FilteredItemsIndices)

	case searchinput.SubmitSearch:
		m.typingSearchTerm = false
		m.filterList(msg.FilteredItemsIndices)

	// Keybinding
	case tea.KeyMsg:
		if m.typingSearchTerm && m.searchInput != nil {
			model, cmd := m.searchInput.Update(msg)
			searchModel := model.(searchinput.Model)
			m.searchInput = &searchModel
			return m, cmd
		} else {
			return keybinds.handle(&m, msg)
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if m.fatalError != nil {
		return Style.FatalErrorStyle.Render(
			fmt.Sprintf("fatal error:\n\n%s", m.fatalError),
		)
	}

	if m.choice != nil {
		return Style.QuitTextStyle.Render(
			fmt.Sprintf("Opening %s 💻", Style.PathTextStyle.Render(m.choice.Path)),
		)
	}

	if m.quitting {
		var sb strings.Builder
		sb.WriteString(Style.QuitTextStyle.Render(quitMessage()))
		sb.WriteString(Style.QuitTextStyleSub.Render("— presented to you by ChatGPT"))
		return sb.String()
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
