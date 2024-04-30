package boards

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss/table"
)

type model struct {
	choices  []string         
	cursor   int             
	selected map[int]struct{}
}

func InitialModel() model {
	return model{
		choices: []string{"Task 1", "Task 2", "Task 3"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Select tasks to complete:\n\n"

	tbl := table.New()

	tbl.Row("Task 1", "Task 2", "Task 3")	

	s += tbl.Render() + "\n"

	for i, choice := range m.choices {

		cursor := " " 
		if m.cursor == i {
			cursor = ">" 
		}

		checked := " " 
		if _, ok := m.selected[i]; ok {
			checked = "x" 
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}