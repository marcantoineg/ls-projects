package components

import (
	"list-my-projects/components/styles"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sahilm/fuzzy"
)

type searchInputModel struct {
	input           textinput.Model
	unfilteredItems []string
}

func newSearchInput(items []string) searchInputModel {
	t := textinput.New()
	t.Placeholder = "Search..."
	t.CursorStyle = styles.SearchInput.InputCursorStyle

	return searchInputModel{
		input:           t,
		unfilteredItems: items,
	}
}

func (m searchInputModel) Init() tea.Cmd {
	return nil
}

func (m searchInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c", "esc":
			return m, func() tea.Msg { return cancelSearch{} }

		case "enter", "down":
			m.input.Blur()
			return m, func() tea.Msg { return submitSearch{m.GetFilteredItems()} }
		}
	}

	input_m, _ := m.input.Update(msg)
	m.input = input_m
	return m, func() tea.Msg { return updateSearch{m.GetFilteredItems()} }
}

func (m searchInputModel) View() string {
	containerStyle := styles.SearchInput.ContainerStyle
	if m.input.Focused() {
		containerStyle = styles.SearchInput.FocusedContainerStyle
	}

	return containerStyle.Render(m.input.View())
}

func (m *searchInputModel) Focus() tea.Cmd {
	return m.input.Focus()
}

func (m searchInputModel) GetFilteredItems() []int {
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

type cancelSearch struct{}
type updateSearch struct {
	filteredItemsIndices []int
}
type submitSearch struct {
	filteredItemsIndices []int
}
