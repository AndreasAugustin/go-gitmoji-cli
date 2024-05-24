package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var isConfigGlobal bool

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: fmt.Sprintf("Setup %s preferences.", pkg.ProgramName),
	Long: `Configure the cli.
			There are default options available which are overwritten
			by the local configuration file or a global configuration file within your OS config folder (use the info command to get the information where it is stored)`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("config called")
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("get current config issue, %s", err)
		}
		autoAdd := runConfigConfirmationPrompt("Enable automatic 'git add .'", config.AutoAdd)
		autoSign := runConfigConfirmationPrompt("Automatically sign commits (add '-S' flag)", config.AutoSign)
		autoSignature := runConfigConfirmationPrompt("Automatically add signature to commits (add '-s' flag)", config.AutoSignature)
		emojiFormat := runEmojiSelectionPrompt("Select how emojis should be used in commits. For a comparison please visit https://gitmoji.dev/specification")
		scopePrompt := runConfigConfirmationPrompt("Enable scope prompt", config.ScopePrompt)
		bodyPrompt := runConfigConfirmationPrompt("Enable body prompt", config.BodyPrompt)
		useDefaultGitMessages := runConfigConfirmationPrompt("Use default git messages (merge, squash,..)", config.UseDefaultGitMessages)
		debug := runConfigConfirmationPrompt("debug mode", config.Debug)
		capitalizeTitle := runConfigConfirmationPrompt("Capitalize title", config.CapitalizeTitle)
		gitmojisApiUrl := runGitmojiUrlInputPrompt("Set gitmojis api url", "https://gitmoji.dev/api/gitmojis")
		config = pkg.Config{
			AutoAdd:               autoAdd,
			AutoSign:              autoSign,
			AutoSignature:         autoSignature,
			EmojiFormat:           emojiFormat,
			ScopePrompt:           scopePrompt,
			CapitalizeTitle:       capitalizeTitle,
			GitmojisUrl:           gitmojisApiUrl,
			BodyPrompt:            bodyPrompt,
			UseDefaultGitMessages: useDefaultGitMessages,
			Debug:                 debug,
		}
		pkg.UpdateConfig(config, isConfigGlobal)
	},
}

func init() {
	RootCmd.AddCommand(ConfigCmd)

	ConfigCmd.PersistentFlags().BoolVarP(&isConfigGlobal, "global", "g", false, "set configuration values globally within ")
}

func runEmojiSelectionPrompt(title string) pkg.EmojiCommitFormats {

	listSettings := ui.ListSettings{Title: title, IsFilteringEnabled: false, IsShowStatusBar: true}
	res := ui.ListRun(listSettings, []pkg.EmojiCommitFormats{pkg.CODE, pkg.EMOJI})

	return res
}

func runGitmojiUrlInputPrompt(title string, initialValue string) string {
	input := ui.TextInputsRun("Gitmoji API url", []pkg.TextInputData{{Placeholder: title, InitialValue: initialValue, Charlimit: 156}})

	return input[0].Value
}

func runConfigConfirmationPrompt(title string, isCurrentlyEnabled bool) bool {
	res := ui.ConfirmationRun(fmt.Sprintf("%s , is currently enabled : %t", title, isCurrentlyEnabled))
	return res
}
