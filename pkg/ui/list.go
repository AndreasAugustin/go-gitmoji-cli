package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type listModel[K interface{ FilterValue() string }] struct {
	list     list.Model
	choice   K
	quitting bool
}

func (m *listModel[K]) Init() tea.Cmd {
	return nil
}

func (m *listModel[K]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(K)
			if ok {
				m.choice = i
			}
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m *listModel[K]) View() string {
	return docStyle.Render(m.list.View())
}

func ListRun[K interface{ FilterValue() string }](listTitle string, input []K) K {
	mapped := make([]list.Item, len(input))

	for i, e := range input {
		mapped[i] = list.Item(e)
	}

	m := listModel[K]{list: list.New(mapped, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = listTitle

	p := tea.NewProgram(&m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	return m.choice
}
