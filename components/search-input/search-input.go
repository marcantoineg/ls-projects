package searchinput

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sahilm/fuzzy"
)

type Model struct {
	input           textinput.Model
	unfilteredItems []string
}

func NewSearchInput(items []string) Model {
	t := textinput.New()
	t.Placeholder = "Search..."
	t.CursorStyle = Style.InputCursorStyle

	return Model{
		input:           t,
		unfilteredItems: items,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "esc":
			return m, func() tea.Msg { return CancelSearch{} }

		case "enter", "down":
			m.input.Blur()
			return m, func() tea.Msg { return SubmitSearch{m.getFilteredItems()} }
		}
	}

	input_m, _ := m.input.Update(msg)
	m.input = input_m
	return m, func() tea.Msg { return UpdateSearch{m.getFilteredItems()} }
}

func (m Model) View() string {
	containerStyle := Style.ContainerStyle
	if m.input.Focused() {
		containerStyle = Style.FocusedContainerStyle
	}

	return containerStyle.Render(m.input.View())
}

func (m *Model) Focus() tea.Cmd {
	return m.input.Focus()
}

// getFilteredItems returns the indices in the unfiltered list of items that are a fuzzy match
// with the search term entered in the text input.
//
// Returns nil if the search term is empty and an empty list if not match is found.
func (m Model) getFilteredItems() []int {
	if m.input.Value() == "" {
		return nil
	}

	matches := fuzzy.Find(m.input.Value(), m.unfilteredItems)
	itemIndices := make([]int, matches.Len())
	for i, match := range matches {
		itemIndices[i] = match.Index
	}

	return itemIndices
}
