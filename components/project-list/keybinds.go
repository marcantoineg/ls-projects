package projectlist

import (
	"fmt"
	projectform "list-my-projects/components/project-form"
	searchinput "list-my-projects/components/search-input"
	"list-my-projects/fileutil"
	"list-my-projects/models/project"
	"os/exec"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
)

// handleKeyMsg handles the keybinding part of the Update function.
func handleListkeybinds(m *Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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
				m.Update(projectform.ProjectUpdateErrorMsg(err))
				return m, nil
			}

			m.items = castToListItem(projects)
			m.list.SetItems(m.items)

			disableMovingMode(m)
		}

	case "a":
		if !m.movingModeActive {
			f := projectform.NewProjectForm(m, nil)
			m.projectForm = &f

			return m.projectForm.Update(nil)
		}

	case "e":
		if !m.movingModeActive {
			if p, ok := m.items[m.list.Index()].(project.Project); ok {
				f := projectform.NewProjectForm(m, &p)
				m.projectForm = &f
			}

			return m.projectForm.Update(nil)
		}

	case "d":
		if !m.movingModeActive {
			if p, ok := m.list.SelectedItem().(project.Project); ok {
				projects, err := project.Delete(m.list.Index(), p)
				if err != nil {
					m.list.Styles.Title = Style.ErrorTitleStyle
					m.list.Title = fmt.Sprintf("error deleting project '%s'", p.Name)
				}

				m.items = castToListItem(projects)
				cmd := m.list.SetItems(m.items)

				m.list.Styles.Title = Style.SuccessTitleStyle
				m.list.Title = fmt.Sprintf("project '%s' deleted", p.Name)

				return m, cmd
			}
		}

	case "y":
		if !clipboard.Unsupported && !m.movingModeActive {
			if p, ok := m.list.SelectedItem().(project.Project); ok {
				clipboard.WriteAll(p.Path)

				m.list.Styles.Title = Style.SuccessTitleStyle
				m.list.Title = fmt.Sprintf("path for project '%s' copied", p.Name)

				return m, nil
			}
		}

	case "m":
		m.movingModeInitialIndex = m.list.Index()
		m.list.SetDelegate(itemDelegate{movingModeInitialIndex: m.movingModeInitialIndex})
		m.movingModeActive = true

		m.list.Styles.Title = Style.MovingModeTitleStyle
		m.list.Title = "select another project to swap position"

	case "f", "/":
		if m.searchInput == nil {
			projectNames := make([]string, len(m.items))
			for i, p := range m.items {
				projectNames[i] = p.(project.Project).Name
			}

			s := searchinput.NewSearchInput(projectNames)
			m.searchInput = &s
			m.searchInput.Focus()
			m.typingSearchTerm = true
			m.searchInput.Update(nil)
		} else {
			m.typingSearchTerm = true
			m.searchInput.Focus()
		}

	case "up":
		if m.searchInput != nil && m.list.Paginator.Page == 0 && m.list.Cursor() == 0 {
			m.typingSearchTerm = true
			m.searchInput.Focus()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
