package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/spf13/cobra"
)

var commitMsg []string
var isDryRun bool

// CommitCmd represents the commit command
var CommitCmd = &cobra.Command{
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
		textInputsData := buildTextInputsData()
		inputsRes := ui.TextInputsRun("please add", textInputsData)
		scopeAndTitle := buildTitleEventualWithScope(inputsRes)
		log.Debugf("scope and title: %s ", scopeAndTitle)
		var selectedEmoji string
		if pkg.ConfigInstance.EmojiFormat == pkg.CODE {
			selectedEmoji = selectedGitmoji.Code
		} else {
			selectedEmoji = selectedGitmoji.Emoji
		}
		completeMessage := fmt.Sprintf("%s %s", selectedEmoji, scopeAndTitle)
		log.Debugf("complete message: %s", completeMessage)
		if isDryRun {
			log.Infof("The commit message: %s", completeMessage)
		}
		//longMessage := ui.TextAreaRun()
		// TODO(anau) check if message prompt enabled
		//log.Debugf("long message %s", longMessage)
		//TODO(anau) do the commit
	},
}

func init() {
	RootCmd.AddCommand(CommitCmd)
	var a []string
	CommitCmd.PersistentFlags().StringSliceVarP(&commitMsg, "message", "m", a, "The commit message. Can be repeated")
	CommitCmd.PersistentFlags().BoolVarP(&isDryRun, "dry-run", "n", false, "dry run: just output the commit message without doing a commit")
}

func buildTitleEventualWithScope(inputsRes []string) string {
	if len(inputsRes) == 0 {
		return ""
	} else if len(inputsRes) == 1 {
		return eventualCapitalizeTitle(inputsRes[0])
	} else {
		return fmt.Sprintf("(%s) %s", inputsRes[0], eventualCapitalizeTitle(inputsRes[1]))
	}
}

func eventualCapitalizeTitle(title string) string {
	if pkg.ConfigInstance.CapitalizeTitle {
		return strings.ToUpper(title)
	}
	return title
}

func buildTextInputsData() []ui.TextInputData {
	var textInputsData []ui.TextInputData

	if pkg.ConfigInstance.ScopePrompt {
		textInputsData = append(textInputsData, ui.TextInputData{Placeholder: "scope", Charlimit: 64, Label: "scope"})
	}

	return append(textInputsData, ui.TextInputData{Placeholder: "title", Charlimit: 64, Label: "title"})
}
