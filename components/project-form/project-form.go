package projectform

import (
	"errors"
	"fmt"
	"ls-projects/models/project"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func focusedStyle(m Model) lipgloss.Style {
	if m.isEditMode {
		return Style.EditFocusedStyle
	} else {
		return Style.NewFocusedStyle
	}
}

func focusedButton(m Model) string {
	return focusedStyle(m).Render("[ submit ]")
}

type Model struct {
	focusIndex int
	inputs     []textinput.Model
	Model      tea.Model
	isEditMode bool
	err        error
}

func NewProjectForm(l tea.Model, project *project.Project) Model {
	m := Model{
		inputs:     make([]textinput.Model, 2),
		Model:      l,
		isEditMode: project != nil,
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = focusedStyle(m)
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

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ProjectCreationErrorMsg:
		m.err = msg

	case tea.KeyMsg:
		tmpModel, tempCmd := handleFormKeybinds(&m, msg)
		if tmpModel != nil {
			return tmpModel, tempCmd
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	var cmds = make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Model) View() string {
	var b strings.Builder

	var title string
	var s lipgloss.Style
	if m.err != nil {
		title = m.err.Error()
		s = Style.ErrorTitleStyle
	} else if m.isEditMode {
		title = "Edit project"
		s = Style.EditProjectTitleStyle
	} else {
		title = "Add new project"
		s = Style.NewProjectTitleStyle
	}
	fmt.Fprintf(&b, "\n%s\n\n", s.Render(title))

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := Style.BlurredButton()
	if m.focusIndex == len(m.inputs) {
		button = focusedButton(m)
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", button)

	fmt.Fprintf(&b, "\n%s", Style.FormHelpStyle.Render("[*] marks required fields"))

	return Style.MarginStyle.Render(b.String())
}

func validateTextField(v string) error {
	if v == "" {
		return errors.New("fields can't be empty")
	}
	return nil
}
