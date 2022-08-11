package main

import (
	"fmt"
	"list-my-projects/components"
	"list-my-projects/models"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	projects, err := models.GetProjects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	p := tea.NewProgram(components.NewListSelector(projects))
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
