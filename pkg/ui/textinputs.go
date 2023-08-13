package ui

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textInputsModel struct {
	focusIndex int
	title      string
	inputs     []textinput.Model
	cursorMode cursor.Mode
}

type TextInputData struct {
	Placeholder  string
	Charlimit    int
	InitialValue string
	Label        string
}

type TextInputRes struct {
	Value string
	Label string
}

func initialTextInputsModel(title string, textInputsData []TextInputData) textInputsModel {
	m := textInputsModel{
		title:  title,
		inputs: make([]textinput.Model, len(textInputsData)),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = textInputsCursorStyle
		t.Placeholder = textInputsData[i].Placeholder
		t.CharLimit = textInputsData[i].Charlimit
		t.Prompt = fmt.Sprintf("%s >", textInputsData[i].Label)

		if textInputsData[i].InitialValue != "" {
			t.SetValue(textInputsData[i].InitialValue)
		}

		if i == 0 {
			t.Focus()
			t.PromptStyle = textInputFocusedStyle
			t.TextStyle = textInputFocusedStyle
		}

		m.inputs[i] = t
	}

	return m
}

func (m *textInputsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *textInputsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = textInputFocusedStyle
					m.inputs[i].TextStyle = textInputFocusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *textInputsModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *textInputsModel) View() string {
	var b strings.Builder
	b.WriteString(m.title)
	b.WriteRune('\n')
	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(helpStyle.Render("cursor mode is "))
	b.WriteString(cursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(helpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func TextInputsRun(title string, textInputsData []TextInputData) []TextInputRes {
	if len(textInputsData) == 0 {
		return []TextInputRes{}
	}
	model := initialTextInputsModel(title, textInputsData)
	if _, err := tea.NewProgram(&model).Run(); err != nil {
		log.Errorf("could not start program: %s\n", err)
		os.Exit(1)
	}
	mapped := make([]TextInputRes, len(textInputsData))

	for i, e := range textInputsData {
		mapped[i].Value = model.inputs[i].Value()
		mapped[i].Label = e.Label
	}
	return mapped
}
