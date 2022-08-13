package components

import (
	"errors"
	"fmt"
	"list-my-projects/models"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	formTitleStyle      = titleStyle.Copy().MarginLeft(0)
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF"))
	formHelpStyle       = blurredStyle.Copy().Italic(true).Faint(true)
	marginStyle         = lipgloss.NewStyle().MarginLeft(4)

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type projectFormModel struct {
	focusIndex        int
	inputs            []textinput.Model
	listSelectorModel *listSelectorModel
}

func NewProjectForm(listSelectorModel *listSelectorModel) projectFormModel {
	m := projectFormModel{
		inputs:            make([]textinput.Model, 2),
		listSelectorModel: listSelectorModel,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = cursorStyle
		t.CharLimit = 64

		switch i {
		case 0:
			t.Placeholder = "Name [*]"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Validate = validateTextField
		case 1:
			t.Placeholder = "Path [*]"
			t.Validate = validateTextField
		}

		m.inputs[i] = t
	}

	return m
}

func (m projectFormModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m projectFormModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			return m.listSelectorModel.Update(noProjectCreatedMsg{})

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter or space while the submit button was focused?
			// If so, create the project and exit to the list-selector component.
			if (s == "enter" || s == "space") && m.focusIndex == len(m.inputs) {
				for i := range m.inputs {
					err := m.inputs[i].Validate(m.inputs[i].Value())
					if err != nil {
						return m.listSelectorModel.Update(projectCreationErrorMsg(err))
					}
				}

				p := &models.Project{
					Name: m.inputs[0].Value(),
					Path: m.inputs[1].Value(),
				}

				if valid := p.ValidatePath(); valid {
					return m.listSelectorModel.Update(projectCreatedMsg{*p})
				} else {
					return m.listSelectorModel.Update(projectCreationErrorMsg(errors.New("project's path is invalid")))
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
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *projectFormModel) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m projectFormModel) View() string {
	var b strings.Builder

	fmt.Fprintf(&b, "\n%s\n\n", formTitleStyle.Render("Add new project"))

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	fmt.Fprintf(&b, "\n%s", formHelpStyle.Render("[*] marks required fields"))

	return marginStyle.Render(b.String())
}

func validateTextField(v string) error {
	if v == "" {
		return errors.New("fields can't be empty")
	}
	return nil
}

type noProjectCreatedMsg struct{}
type projectCreationErrorMsg error
type projectCreatedMsg struct {
	project models.Project
}
