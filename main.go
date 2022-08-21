package main

import (
	"fmt"
	"list-my-projects/components"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	i := 3
	fmt.Printf("something something %s\n", i)

	p := tea.NewProgram(components.NewListSelector())
	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
