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
var body string

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
			log.Infof("There are hook files existent for %s", existentHookFiles)
			log.Infof("Please use git commit command or remove the hooks with %s hooks rm", pkg.ProgramName)
			spin.Stop()
			return
		}
		gitmojis := pkg.GetGitmojis()
		spin.Stop()
		initialCommitValues := pkg.InitialCommitValues{Type: _type, Scope: scope, Desc: desc, Body: body}
		listSettings := ui.ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
		selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
		log.Debugf("selected gitmoji %s", selectedGitmoji)
		textInputsData := initialCommitValues.BuildTextInputsData(pkg.ConfigInstance)
		inputsRes := ui.TextInputsRun("please add", textInputsData)

		commitValues := pkg.CreateMessage(inputsRes, selectedGitmoji, initialCommitValues, pkg.ConfigInstance, isBreaking)

		log.Debugf("complete title: %s", commitValues.Title)
		if isDryRun {
			log.Infof("The commit title: %s", commitValues.Title)
			log.Infof("The commit body: %s", commitValues.Body)
		} else {
			pkg.ExecuteCommit(commitValues.Title, commitValues.Body, pkg.ConfigInstance)
		}
	},
}

func init() {
	RootCmd.AddCommand(CommitCmd)
	var a []string
	CommitCmd.PersistentFlags().StringSliceVarP(&commitMsg, "message", "m", a, "The commit message. Can be repeated")
	CommitCmd.PersistentFlags().BoolVarP(&isDryRun, "dry-run", "n", false, "dry run: just output the commit message without doing a commit")
	CommitCmd.PersistentFlags().StringVar(&_type, string(pkg.TYPE), "", "add the type")
	CommitCmd.PersistentFlags().StringVar(&scope, string(pkg.SCOPE), "", "add a scope")
	CommitCmd.PersistentFlags().StringVar(&desc, string(pkg.DESC), "", "add a description")
	CommitCmd.PersistentFlags().StringVar(&body, string(pkg.BODY), "", "add the commit message body")
	CommitCmd.PersistentFlags().BoolVar(&isBreaking, string(pkg.IS_BREAKING), false, "set if the commit is a breaking change")
}
