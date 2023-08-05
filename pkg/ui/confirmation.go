package ui

import (
	"bytes"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erikgeiser/promptkit"
	"github.com/muesli/termenv"
	"html/template"
	"io"
	"os"
)

const TemplateArrow = `
{{- Bold .Prompt -}}
{{ if .YesSelected -}}
	{{- print (Bold " ▸Yes ") " No" -}}
{{- else if .NoSelected -}}
	{{- print "  Yes " (Bold "▸No") -}}
{{- else -}}
	{{- "  Yes  No" -}}
{{- end -}}
`

// ResultTemplateArrow is the ResultTemplate that matches TemplateArrow.
const ResultTemplateArrow = `
{{- print .Prompt " " -}}
{{- if .FinalValue -}}
	{{- Foreground "32" "Yes" -}}
{{- else -}}
	{{- Foreground "32" "No" -}}
{{- end }}
`

const (
	// DefaultTemplate defines the default appearance of the text input and can
	// be copied as a starting point for a custom template.
	DefaultTemplate = TemplateArrow

	// DefaultResultTemplate defines the default appearance with which the
	// finale result of the prompt is presented.
	DefaultResultTemplate = ResultTemplateArrow
)

// Value is the value of the confirmation prompt which can be Undecided, Yes or
// No.
type Value *bool

// NewValue creates a Value from a bool.
func NewValue(v bool) Value {
	if v {
		return Yes
	}

	return No
}

// NewDefaultKeyMap returns a KeyMap with sensible default key mappings that can
// also be used as a starting point for customization.
func NewDefaultKeyMap() *KeyMap {
	return &KeyMap{
		Yes:       []string{"y", "Y"},
		No:        []string{"n", "N"},
		SelectYes: []string{"left"},
		SelectNo:  []string{"right"},
		Toggle:    []string{"tab"},
		Submit:    []string{"enter"},
		Abort:     []string{"ctrl+c"},
	}
}

// KeyMap defines the keys that trigger certain actions.
type KeyMap struct {
	Yes       []string
	No        []string
	SelectYes []string
	SelectNo  []string
	Toggle    []string
	Submit    []string
	Abort     []string
}

func keyMatches(key tea.KeyMsg, mapping []string) bool {
	for _, m := range mapping {
		if m == key.String() {
			return true
		}
	}

	return false
}

// validateKeyMap returns true if the given key map contains at
// least the bare minimum set of key bindings for the functional
// prompt and false otherwise.
func validateKeyMap(km *KeyMap) error {
	if len(km.Yes) == 0 && len(km.No) == 0 && len(km.Submit) == 0 {
		return fmt.Errorf("no submit key")
	}

	if !(len(km.Yes) > 0 && len(km.No) > 0) &&
		len(km.Toggle) == 0 &&
		!(len(km.SelectYes) > 0 && len(km.SelectNo) > 0) {
		return fmt.Errorf("missing keys to select a value")
	}

	return nil
}

var (
	yes = true
	no  = false

	// Yes is a possible value of the confirmation prompt that corresponds to
	// true.
	Yes = Value(&yes)
	// No is a possible value of the confirmation prompt that corresponds to
	// false.
	No = Value(&no)
	// Undecided is a possible value of the confirmation prompt that is used
	// when neither Yes nor No are selected.
	Undecided = Value(nil)
)

// Confirmation represents a configurable confirmation prompt.
type Confirmation struct {
	// Prompt holds the question.
	Prompt string

	// DefaultValue decides if a value should already be selected at startup. By
	// default it is Undecided but it can be set to Yes (corresponds to true)
	// and No (corresponds to false).
	DefaultValue Value

	// Template holds the display template. A custom template can be used to
	// completely customize the appearance of the text input. If empty, the
	// DefaultTemplate is used. The following variables and functions are
	// available:
	//
	//  * Prompt string: The configured prompt.
	//  * YesSelected bool: Whether or not Yes is the currently selected
	//    value.
	//  * NoSelected bool: Whether or not No is the currently selected value.
	//  * Undecided bool: Whether or not Undecided is the currently selected
	//    value.
	//  * DefaultYes bool: Whether or not Yes is confiured as default value.
	//  * DefaultNo bool: Whether or not No is confiured as default value.
	//  * DefaultUndecided bool: Whether or not Undecided is confiured as
	//    default value.
	//  * TerminalWidth int: The width of the terminal.
	//  * promptkit.UtilFuncMap: Handy helper functions.
	//  * termenv TemplateFuncs (see https://github.com/muesli/termenv).
	//  * The functions specified in ExtendedTemplateFuncs.
	Template string

	// ResultTemplate is rendered as soon as a input has been confirmed.
	// It is intended to permanently indicate the result of the prompt when the
	// input itself has disappeared. This template is only rendered in the Run()
	// method and NOT when the text input is used as a textInputModel. The following
	// variables and functions are available:
	//
	//  * FinalValue bool: The final value of the confirmation.
	//  * FinalValue string: The final value's string representation ("true"
	//    or "false").
	//  * Prompt string: The configured prompt.
	//  * DefaultYes bool: Whether or not Yes is confiured as default value.
	//  * DefaultNo bool: Whether or not No is confiured as default value.
	//  * DefaultUndecided bool: Whether or not Undecided is confiured as
	//    default value.
	//  * TerminalWidth int: The width of the terminal.
	//  * promptkit.UtilFuncMap: Handy helper functions.
	//  * termenv TemplateFuncs (see https://github.com/muesli/termenv).
	//  * The functions specified in ExtendedTemplateFuncs.
	ResultTemplate string

	// ExtendedTemplateFuncs can be used to add additional functions to the
	// evaluation scope of the templates.
	ExtendedTemplateFuncs template.FuncMap

	// KeyMap determines with which keys the confirmation prompt is controlled.
	// By default, DefaultKeyMap is used.
	KeyMap *KeyMap

	// WrapMode decides which way the prompt view is wrapped if it does not fit
	// the terminal. It can be a WrapMode provided by promptkit or a custom
	// function. By default it is promptkit.WordWrap. It can also be nil which
	// disables wrapping and likely causes output glitches.
	WrapMode promptkit.WrapMode

	// Output is the output writer, by default os.Stdout is used.
	Output io.Writer
	// Input is the input reader, by default, os.Stdin is used.
	Input io.Reader

	// ColorProfile determines how colors are rendered. By default, the terminal
	// is queried.
	ColorProfile termenv.Profile
}

// Model implements the bubbletea.Model for a confirmation prompt.
type Model struct {
	*Confirmation

	// Err holds errors that may occur during the execution of
	// the confirmation prompt.
	Err error

	// MaxWidth limits the width of the view using the Confirmation's WrapMode.
	MaxWidth int

	tmpl       *template.Template
	resultTmpl *template.Template

	value Value

	quitting bool

	width int
}

// ensure that the Model interface is implemented.
var _ tea.Model = &Model{}

// NewModel returns a new textInputModel based on the provided confirmation prompt.
func NewModel(confirmation *Confirmation) *Model {
	return &Model{
		Confirmation: confirmation,
		value:        confirmation.DefaultValue,
	}
}

// Init initializes the confirmation prompt textInputModel.
func (m *Model) Init() tea.Cmd {
	m.tmpl, m.Err = m.initTemplate()
	if m.Err != nil {
		return tea.Quit
	}

	m.resultTmpl, m.Err = m.initResultTemplate()
	if m.Err != nil {
		return tea.Quit
	}

	return textinput.Blink
}

func (m *Model) initTemplate() (*template.Template, error) {
	tmpl := template.New("view")
	tmpl.Funcs(termenv.TemplateFuncs(m.ColorProfile))
	tmpl.Funcs(promptkit.UtilFuncMap())
	tmpl.Funcs(m.ExtendedTemplateFuncs)

	return tmpl.Parse(m.Template)
}

func (m *Model) initResultTemplate() (*template.Template, error) {
	if m.ResultTemplate == "" {
		return nil, nil
	}

	tmpl := template.New("result")
	tmpl.Funcs(termenv.TemplateFuncs(m.ColorProfile))
	tmpl.Funcs(promptkit.UtilFuncMap())
	tmpl.Funcs(m.ExtendedTemplateFuncs)

	return tmpl.Parse(m.ResultTemplate)
}

// Update updates the textInputModel based on the received message.
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.Err != nil {
		return m, tea.Quit
	}

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case keyMatches(msg, m.KeyMap.Submit):
			if m.value != Undecided {
				m.quitting = true

				return m, tea.Quit
			}
		case keyMatches(msg, m.KeyMap.Abort):
			m.Err = promptkit.ErrAborted
			m.quitting = true

			return m, tea.Quit
		case keyMatches(msg, m.KeyMap.Yes):
			m.value = Yes
			m.quitting = true

			return m, tea.Quit
		case keyMatches(msg, m.KeyMap.No):
			m.value = No
			m.quitting = true

			return m, tea.Quit
		case keyMatches(msg, m.KeyMap.SelectYes):
			m.value = Yes
		case keyMatches(msg, m.KeyMap.SelectNo):
			m.value = No
		case keyMatches(msg, m.KeyMap.Toggle):
			switch m.value {
			case Yes:
				m.value = No
			case No, Undecided:
				m.value = Yes
			}
		}
	case tea.WindowSizeMsg:
		m.width = zeroAwareMin(msg.Width, m.MaxWidth)
	case error:
		m.Err = msg

		return m, tea.Quit
	}

	return m, cmd
}

// View renders the confirmation prompt.
func (m *Model) View() string {
	// avoid panics if Quit is sent during Init
	if m.quitting {
		view, err := m.resultView()
		if err != nil {
			m.Err = err

			return ""
		}

		return m.wrap(view)
	}

	// avoid panics if Quit is sent during Init
	if m.tmpl == nil {
		return ""
	}

	viewBuffer := &bytes.Buffer{}

	err := m.tmpl.Execute(viewBuffer, map[string]interface{}{
		"Prompt":           m.Prompt,
		"YesSelected":      m.value == Yes,
		"NoSelected":       m.value == No,
		"Undecided":        m.value == Undecided,
		"DefaultYes":       m.DefaultValue == Yes,
		"DefaultNo":        m.DefaultValue == No,
		"DefaultUndecided": m.DefaultValue == Undecided,
		"TerminalWidth":    m.width,
	})
	if err != nil {
		m.Err = err

		return "Template Error: " + err.Error()
	}

	return m.wrap(viewBuffer.String())
}

func (m *Model) resultView() (string, error) {
	viewBuffer := &bytes.Buffer{}

	if m.ResultTemplate == "" {
		return "", nil
	}

	if m.resultTmpl == nil {
		return "", fmt.Errorf("rendering confirmation without loaded template")
	}

	value, err := m.Value()
	if err != nil {
		return "", err
	}

	err = m.resultTmpl.Execute(viewBuffer, map[string]interface{}{
		"FinalValue":       value,
		"FinalValueString": fmt.Sprintf("%v", value),
		"Prompt":           m.Prompt,
		"DefaultYes":       m.DefaultValue == Yes,
		"DefaultNo":        m.DefaultValue == No,
		"DefaultUndecided": m.DefaultValue == Undecided,
		"TerminalWidth":    m.width,
	})
	if err != nil {
		return "", fmt.Errorf("execute confirmation template: %w", err)
	}

	return viewBuffer.String(), nil
}

func (m *Model) wrap(text string) string {
	if m.WrapMode == nil {
		return text
	}

	return m.WrapMode(text, m.width)
}

// Value returns the current value and error.
func (m *Model) Value() (bool, error) {
	if m.Err != nil {
		return false, m.Err
	}

	if m.value == Undecided {
		return false, fmt.Errorf("no decision was made")
	}

	return *m.value, m.Err
}

func zeroAwareMin(a int, b int) int {
	switch {
	case a == 0:
		return b
	case b == 0:
		return a
	case a > b:
		return b
	default:
		return a
	}
}

// New creates a new text input. If the default value is nil it is equivalent to
// Undecided. See the Confirmation properties for more documentation.
func New(prompt string, defaultValue Value) *Confirmation {
	return &Confirmation{
		Prompt:                prompt,
		DefaultValue:          defaultValue,
		Template:              DefaultTemplate,
		ResultTemplate:        DefaultResultTemplate,
		KeyMap:                NewDefaultKeyMap(),
		ExtendedTemplateFuncs: template.FuncMap{},
		WrapMode:              promptkit.Truncate,
		Output:                os.Stdout,
		Input:                 os.Stdin,
	}
}

// RunPrompt executes the confirmation prompt.
func (c *Confirmation) RunPrompt() (bool, error) {
	err := validateKeyMap(c.KeyMap)
	if err != nil {
		return false, fmt.Errorf("insufficient key map: %w", err)
	}

	m := NewModel(c)

	p := tea.NewProgram(m, tea.WithOutput(c.Output), tea.WithInput(c.Input))

	_, err = p.Run()
	if err != nil {
		return false, fmt.Errorf("running prompt: %w", err)
	}

	return m.Value()
}
