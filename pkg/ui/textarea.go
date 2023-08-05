package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

type textAreaErrMsg error

type textAreaModel struct {
	textarea textarea.Model
	err      error
}

func initialTextAreaModel() textAreaModel {
	ti := textarea.New()
	ti.Placeholder = "Once upon a time..."
	ti.Focus()

	return textAreaModel{
		textarea: ti,
		err:      nil,
	}
}

func (m *textAreaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m *textAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case textAreaErrMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *textAreaModel) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func TextAreaRun() string {
	model := initialTextAreaModel()
	p := tea.NewProgram(&model)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	return model.textarea.View()
}
