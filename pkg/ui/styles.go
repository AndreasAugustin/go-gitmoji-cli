package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var (
	textInputFocusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	textInputsBlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	textInputsCursorStyle  = textInputFocusedStyle.Copy()
	noStyle                = lipgloss.NewStyle()
	helpStyle              = textInputsBlurredStyle.Copy()
	cursorModeHelpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = textInputFocusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", textInputsBlurredStyle.Render("Submit"))
)
