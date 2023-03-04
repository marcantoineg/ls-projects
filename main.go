package main

import (
	"fmt"
	projectlist "ls-projects/components/project-list"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(projectlist.NewProjectList())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
