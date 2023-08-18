package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
	"os"
)

type errMsgSpinner error

type spinnerModel struct {
	spinner  spinner.Model
	err      error
	quitting bool
}

func initialSpinnerModel() spinnerModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return spinnerModel{spinner: s}
}

func (m *spinnerModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *spinnerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			log.Warn("ctrl + c -> quitting")
			os.Exit(0)
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsgSpinner:
		m.err = msg
		return m, nil

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m *spinnerModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Loading ...press q to quit\n\n", m.spinner.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

type Spinner struct {
	model   *spinnerModel
	program *tea.Program
}

func NewSpinner() Spinner {
	model := initialSpinnerModel()
	return Spinner{
		model:   &model,
		program: tea.NewProgram(&model),
	}
}

func (s *Spinner) Run() {
	go func() {
		_, err := s.program.Run()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

}

func (s *Spinner) Stop() {
	s.model.quitting = true
	s.program.Quit()
}
