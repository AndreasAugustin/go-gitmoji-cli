package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
	"os"
)

type ListSettings struct {
	Title              string
	IsFilteringEnabled bool
	IsShowStatusBar    bool
}

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

func ListRun[K interface{ FilterValue() string }](settings ListSettings, input []K) K {
	mapped := make([]list.Item, len(input))

	for i, e := range input {
		mapped[i] = list.Item(e)
	}

	_list := list.New(mapped, list.NewDefaultDelegate(), 0, 0)
	_list.SetFilteringEnabled(settings.IsFilteringEnabled)
	_list.SetShowStatusBar(settings.IsShowStatusBar)
	m := listModel[K]{list: _list}
	m.list.Title = settings.Title

	p := tea.NewProgram(&m, tea.WithAltScreen())

	_, err := p.Run()

	if err != nil {
		log.Errorf("Error running program %s", err)
		os.Exit(1)
	}

	if m.quitting {
		log.Warn("ctrl + c pressed -> quitting")
		os.Exit(0)
	}

	return m.choice
}
