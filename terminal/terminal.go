package terminal

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
);

func RenderModel(initialModel tea.Model) {
	p := tea.NewProgram(initialModel)

    if _, err := p.Run(); err != nil {
        fmt.Printf("[ERROR] %v", err)
        os.Exit(1)
    }
}