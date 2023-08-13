package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var commitMsg []string
var isDryRun bool
var scope string
var desc string
var isBreaking bool
var _type string

// CommitCmd represents the commit command
var CommitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Interactively commit using prompts",
	Long:  `Do the commit. This command is disabled when you are using commit hooks`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("commit called")
		log.Debug(commitMsg)
		spin := ui.NewSpinner()
		spin.Run()
		existentHookFiles, err := pkg.HookFilesExistent()
		if err != nil {
			log.Fatalf("Error checking if hook files existent")
		}
		if len(existentHookFiles) > 0 {
			log.Infof("There are hook files existen for %s", existentHookFiles)
			log.Infof("Please use git commit command or remove the hooks with %s hooks rm", pkg.ProgramName)
			spin.Stop()
			return
		}
		gitmojis := pkg.GetGitmojis()
		spin.Stop()
		listSettings := ui.ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
		selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
		log.Debugf("selected gitmoji %s", selectedGitmoji)
		textInputsData := buildTextInputsData()
		inputsRes := ui.TextInputsRun("please add", textInputsData)
		msg := pkg.BuildCommitTitle(extractType(inputsRes), extractScope(inputsRes), isBreaking, extractDesc(inputsRes), selectedGitmoji, pkg.ConfigInstance)
		log.Debugf("complete message: %s", msg)
		if isDryRun {
			log.Infof("The commit message: %s", msg)
		}
		// TODO(anau) autosign and autoadd
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
	CommitCmd.PersistentFlags().StringVar(&_type, "type", "", "add the type")
	CommitCmd.PersistentFlags().StringVar(&scope, "scope", "", "add a scope")
	CommitCmd.PersistentFlags().StringVar(&desc, "desc", "", "add a description")
	CommitCmd.PersistentFlags().BoolVar(&isBreaking, "is-breaking", false, "set if the commit is a breaking change")
}

func extractType(inputsRes []string) string {
	return inputsRes[0]
}

func extractScope(inputsRes []string) string {
	if len(inputsRes) > 2 {
		return inputsRes[1]
	} else {
		return scope
	}
}

func extractDesc(inputsRes []string) string {
	if len(inputsRes) > 2 {
		return inputsRes[2]
	} else {
		return inputsRes[1]
	}
}

func buildTextInputsData() []ui.TextInputData {
	var textInputsData = []ui.TextInputData{{Placeholder: "type", Charlimit: 64, Label: "type", InitialValue: _type}}

	if pkg.ConfigInstance.ScopePrompt {
		textInputsData = append(textInputsData, ui.TextInputData{Placeholder: "scope", Charlimit: 64, Label: "scope", InitialValue: scope})
	}

	return append(textInputsData, ui.TextInputData{Placeholder: "description", Charlimit: 64, Label: "description", InitialValue: desc})
}
