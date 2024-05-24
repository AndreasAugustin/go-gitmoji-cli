package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)
var (
	textInputFocusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	textInputsBlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	textInputsCursorStyle  = textInputFocusedStyle
	noStyle                = lipgloss.NewStyle()
	helpStyle              = textInputsBlurredStyle
	cursorModeHelpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = textInputFocusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", textInputsBlurredStyle.Render("Submit"))
)
