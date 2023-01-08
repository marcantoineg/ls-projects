package projectlist

import (
	"list-my-projects/models/project"

	"github.com/charmbracelet/bubbles/list"
)

// castToListItem takes a list of 'Project's and returns it as a casted list of tea's interface 'list.Item'.
func castToListItem(projects []project.Project) []list.Item {
	castedItems := make([]list.Item, len(projects))
	for i, p := range projects {
		castedItems[i] = p
	}
	return castedItems
}

// resetListTitle resets the initial style and text of the list's title.
func resetListTitle(m *Model) {
	m.list.Styles.Title = Style.TitleStyle
	m.list.Title = listInitialTitle
}

// disableMovingMode resets required value to disable the moving mode.
func disableMovingMode(m *Model) {
	m.movingModeInitialIndex = -1
	m.movingModeActive = false
	m.list.SetDelegate(itemDelegate{movingModeInitialIndex: -1})
	resetListTitle(m)
}

// filterList filters the items in the list (m.list) given a list of indices.
func (m *Model) filterList(filteredIndices []int) {
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
