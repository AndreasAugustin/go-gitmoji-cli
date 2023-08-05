package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var commitMsg []string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Interactively commit using prompts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("commit called")
		log.Debug(commitMsg)
		spin := ui.NewSpinner()
		spin.Run()

		gitmojis := pkg.GetGitmojis()
		spin.Stop()
		listSettings := ui.ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
		selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
		log.Debugf("selected gitmoji %s", selectedGitmoji)
		scopeAndMessage := ui.TextInputsRun()
		scope := scopeAndMessage[0]
		message := scopeAndMessage[1]
		log.Debugf("scope: %s and message: %s", scope, message)
		completeMessage := fmt.Sprintf("%s: %s %s", scope, message, selectedGitmoji.Code)
		log.Debugf("complete message: %s", completeMessage)
		longMessage := ui.TextAreaRun()
		log.Debugf("long message %s", longMessage)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	var a []string
	commitCmd.PersistentFlags().StringSliceVarP(&commitMsg, "message", "m", a, "The commit message. Can be repeated")
}
