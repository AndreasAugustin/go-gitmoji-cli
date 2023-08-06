package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	log "github.com/sirupsen/logrus"
)

type (
	errMsgTextInput error
)

type textInputModel struct {
	textInput textinput.Model
	label     string
	err       error
}

func initialTextInputModel(label string, initialValue string) textInputModel {
	ti := textinput.New()
	ti.SetValue(initialValue)
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 100

	return textInputModel{
		textInput: ti,
		label:     label,
		err:       nil,
	}
}

func (m *textInputModel) Init() tea.Cmd {
	return nil
}

func (m *textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsgTextInput:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *textInputModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n",
		m.label,
		m.textInput.View(),
		"(esc to quit)",
	)
}

func TextInputRun(label string, initialValue string) string {
	model := initialTextInputModel(label, initialValue)
	p := tea.NewProgram(&model)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

	return model.textInput.Value()
}
