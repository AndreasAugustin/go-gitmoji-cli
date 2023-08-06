package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var isConfigGlobal bool

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
		emojiFormat := runEmojiSelectionPrompt("Select how emojis should be used in commits. For a comparison please visit https://gitmoji.dev/specification")
		scopePrompt := runConfigConfirmationPrompt("Enable scope prompt", config.ScopePrompt)
		messagePrompt := runConfigConfirmationPrompt("Enable message prompt", config.MessagePrompt)
		capitalizeTitle := runConfigConfirmationPrompt("Capitalize title", config.CapitalizeTitle)
		gitmojisApiUrl := runGitmojiUrlInputPrompt("Set gitmojis api url", "https://gitmoji.dev/api/gitmojis")
		config = pkg.Config{Autoadd: autoAdd, EmojiFormat: emojiFormat, ScopePrompt: scopePrompt, CapitalizeTitle: capitalizeTitle, GitmojisUrl: gitmojisApiUrl, MessagePrompt: messagePrompt}
		pkg.UpdateConfig(config, isConfigGlobal)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().BoolVarP(&isConfigGlobal, "global", "g", false, "set configuration values globally within ")

}

func runEmojiSelectionPrompt(title string) pkg.EmojiCommitFormats {

	listSettings := ui.ListSettings{Title: title, IsFilteringEnabled: false, IsShowStatusBar: true}
	res := ui.ListRun(listSettings, []pkg.EmojiCommitFormats{pkg.CODE, pkg.EMOJI})

	return res
}

func runGitmojiUrlInputPrompt(title string, initialValue string) string {
	input := ui.TextInputRun(title, initialValue)

	return input
}

func runConfigConfirmationPrompt(title string, isCurrentlyEnabled bool) bool {
	res := ui.ConfirmationRun(fmt.Sprintf("%s , is currently enabled : %t", title, isCurrentlyEnabled))
	return res
}
