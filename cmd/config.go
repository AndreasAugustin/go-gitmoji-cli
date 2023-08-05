package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/charmbracelet/lipgloss"
	"github.com/erikgeiser/promptkit/confirmation"
	"github.com/erikgeiser/promptkit/selection"
	"github.com/erikgeiser/promptkit/textinput"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle.Copy()
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle.Copy()
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: fmt.Sprintf("Setup %s preferences.", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("config called")
		config := pkg.ConfigInstance
		autoAdd := runConfigConfirmationPrompt("Enable automatic 'git add .'", config.Autoadd)
		emojiStyle := runConfigSelectionPrompt("Select how emojis should be used in commits", []string{":smile:", "😄"})
		scopePrompt := runConfigConfirmationPrompt("Enable scope prompt", config.ScopePrompt)
		messagePrompt := runConfigConfirmationPrompt("Enable message prompt", config.MessagePrompt)
		capitalizeTitle := runConfigConfirmationPrompt("Capitalize title", config.CapitalizeTitle)
		gitmojisApiUrl := runTextInputPrompt("Set gitmojis api url", "https://gitmoji.dev/api/gitmojis")
		config = pkg.Config{Autoadd: autoAdd, EmojiFormat: getEmojiFormat(emojiStyle), ScopePrompt: scopePrompt, CapitalizeTitle: capitalizeTitle, GitmojisUrl: gitmojisApiUrl, MessagePrompt: messagePrompt}
		pkg.UpdateConfig(config)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func getEmojiFormat(input string) pkg.EmojiCommitFormats {
	if input == ":smile:" {
		return pkg.CODE
	} else {
		return pkg.EMOJI
	}
}

func runConfigSelectionPrompt[T any](title string, input []T) T {
	sp := selection.New(title, input)
	sp.PageSize = 3

	choice, err := sp.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}
	return choice
}

func runTextInputPrompt(title string, initialValue string) string {
	input := textinput.New(title)
	input.InitialValue = initialValue
	input.Placeholder = "cannot be empty"
	input.InputCursorStyle = cursorStyle
	input.InputTextStyle = focusedStyle

	val, err := input.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}
	return val
}

func runConfigConfirmationPrompt(title string, isCurrentlyEnabled bool) bool {
	prompt := confirmation.New(title, confirmation.NewValue(isCurrentlyEnabled))

	ready, err := prompt.RunPrompt()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

		os.Exit(1)
	}
	return ready
}
