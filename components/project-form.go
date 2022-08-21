package components

import (
	"errors"
	"fmt"
	models "list-my-projects/models"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	newProjectTitleStyle  = titleStyle.Copy().MarginLeft(0)
	editProjectTitleStyle = newProjectTitleStyle.Copy().Background(lipgloss.Color("#DDB771")).Foreground(lipgloss.Color("#000000"))
	newFocusedStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF"))
	editFocusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#DDB771"))
	blurredStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle               = lipgloss.NewStyle()
	cursorModeHelpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#6C91BF"))
	formHelpStyle         = blurredStyle.Copy().Italic(true).Faint(true)
	marginStyle           = lipgloss.NewStyle().MarginLeft(4)

	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

func focusedButton(m projectFormModel) string {
	return focusedStyle(m).Render("[ submit ]")
}

func focusedStyle(m projectFormModel) lipgloss.Style {
	if m.isEditMode {
		return editFocusedStyle
	} else {
		return newFocusedStyle
	}
}

type projectFormModel struct {
	focusIndex        int
	inputs            []textinput.Model
	listSelectorModel *listSelectorModel
	isEditMode        bool
}

func NewProjectForm(l *listSelectorModel, project *models.Project) projectFormModel {
	m := projectFormModel{
		inputs:            make([]textinput.Model, 2),
		listSelectorModel: l,
		isEditMode:        project != nil,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = focusedStyle(m)
		t.CharLimit = 0

		switch i {
		case 0:
			t.Placeholder = "Name [*]"
			t.Focus()
			t.PromptStyle = focusedStyle(m)
			t.TextStyle = focusedStyle(m)
			t.Validate = validateTextField
			if m.isEditMode {
				t.SetValue(project.Name)
			}

		case 1:
			t.Placeholder = "Path [*]"
			t.Validate = validateTextField
			if m.isEditMode {
				t.SetValue(project.Path)
			}
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
					var msg tea.Msg
					if m.isEditMode {
						msg = projectUpdatedMsg{*p}
					} else {
						msg = projectCreatedMsg{*p}
					}
					return m.listSelectorModel.Update(msg)
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
					m.inputs[i].PromptStyle = focusedStyle(m)
					m.inputs[i].TextStyle = focusedStyle(m)
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

	if m.isEditMode {
		fmt.Fprintf(&b, "\n%s\n\n", editProjectTitleStyle.Render("Edit project"))
	} else {
		fmt.Fprintf(&b, "\n%s\n\n", newProjectTitleStyle.Render("Add new project"))
	}

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := blurredButton
	if m.focusIndex == len(m.inputs) {
		button = focusedButton(m)
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", button)

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
type projectCreatedMsg struct {
	project models.Project
}
type projectCreationErrorMsg error
type projectUpdatedMsg struct {
	project models.Project
}
type projectUpdateErrorMsg error
