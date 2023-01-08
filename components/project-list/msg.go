package projectlist

import "github.com/charmbracelet/bubbles/list"

type fatalErrorMsg struct {
	err error
}
type initMsg struct{ items []list.Item }
