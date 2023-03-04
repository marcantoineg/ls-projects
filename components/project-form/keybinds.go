package projectform

import (
	"errors"
	"ls-projects/models/project"

	tea "github.com/charmbracelet/bubbletea"
)

func handleFormKeybinds(m *Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c":
		return m, tea.Quit

	case "esc":
		return m.Model.Update(NoProjectCreatedMsg{})

	// Set focus to next input
	case "tab", "shift+tab", "enter", "up", "down":
		s := msg.String()

		// Did the user press enter or space while the submit button was focused?
		// If so, create the project and exit to the list-selector component.
		if (s == "enter" || s == "space") && m.focusIndex == len(m.inputs) {
			for i := range m.inputs {
				err := m.inputs[i].Validate(m.inputs[i].Value())
				if err != nil {
					return m.Update(ProjectCreationErrorMsg(err))
				}
			}

			p := &project.Project{
				Name: m.inputs[0].Value(),
				Path: m.inputs[1].Value(),
			}

			if valid := p.ValidatePath(); valid {
				var msg tea.Msg
				if m.isEditMode {
					msg = ProjectUpdatedMsg{*p}
				} else {
					msg = ProjectCreatedMsg{*p}
				}
				return m.Model.Update(msg)
			} else {
				return m.Update(ProjectCreationErrorMsg(errors.New("project's path is invalid")))
			}
		}

		// Cycle indexes
		if s == "up" || s == "shift+tab" {
			m.focusIndex--
		} else {
			m.focusIndex++
		}

		if m.focusIndex > len(m.inputs) {
			m.focusIndex = 0
		} else if m.focusIndex < 0 {
			m.focusIndex = len(m.inputs)
		}

		cmds := make([]tea.Cmd, len(m.inputs))
		for i := 0; i <= len(m.inputs)-1; i++ {
			if i == m.focusIndex {
				cmds[i] = m.inputs[i].Focus()
				m.inputs[i].PromptStyle = focusedStyle(*m)
				m.inputs[i].TextStyle = focusedStyle(*m)
				continue
			}
			// Remove focused state
			m.inputs[i].Blur()
			m.inputs[i].PromptStyle = Style.NoStyle
			m.inputs[i].TextStyle = Style.NoStyle
		}

		return m, tea.Batch(cmds...)
	}

	return nil, nil
}
