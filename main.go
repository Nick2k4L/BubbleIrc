package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/Nick2k4L/BubbleIrc/models"
)

func main() {
	p := tea.NewProgram(models.InitiateApp())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
