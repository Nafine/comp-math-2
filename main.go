package main

import (
	"comp-math-2/internal/model"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if _, err := tea.NewProgram(model.InitialModel()).Run(); err != nil {
		fmt.Printf("Fatal: %v", err)
	}
}
