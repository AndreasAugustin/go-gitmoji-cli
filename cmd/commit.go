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
	PreRun: func(cmd *cobra.Command, args []string) {
		programNameFigure()
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("commit called")
		log.Debug(commitMsg)
		spin := ui.NewSpinner()
		spin.Run()
		defer func() { spin.Stop() }()
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
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("get current config issue, %s", err)
		}
		gitmojis := pkg.GetGitmojis(config)
		defaultTypes := pkg.DefaultCommitTypes()
		initialCommitValues := pkg.BuildInitialCommitValues(
			_type,
			scope,
			desc,
			body,
			commitMsg,
			gitmojis.Gitmojis,
		)
		listSettingsGitmojis := ui.ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
		listSettingsCommitTypes := ui.ListSettings{Title: "Commit types", IsShowStatusBar: true, IsFilteringEnabled: true}
		spin.Stop()
		selectedGitmoji := ui.ListRun(listSettingsGitmojis, gitmojis.Gitmojis)
		log.Debugf("selected gitmoji %s", selectedGitmoji)
		selectedDefaultType := ui.ListRun(listSettingsCommitTypes, defaultTypes)
		log.Debugf("selected %s", selectedDefaultType)
		initialCommitValues.Type = selectedDefaultType.Type
		textInputsData := initialCommitValues.BuildTextInputsData(config)
		inputsRes := ui.TextInputsRun("please add", textInputsData)

		commitValues := pkg.CreateMessage(inputsRes, selectedGitmoji, initialCommitValues, config, isBreaking)

		log.Debugf("complete title: %s", commitValues.Title)
		if isDryRun {
			log.Infof("The commit title: %s", commitValues.Title)
			log.Infof("The commit body: %s", commitValues.Body)
		} else {
			pkg.ExecuteCommit(commitValues.Title, commitValues.Body, config)
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

	bindAndAddBoolFlagP(CommitCmd, string(pkg.AUTO_ADD), "auto-add", "a", "call git add . before commit")
	bindAndAddBoolFlagP(CommitCmd, string(pkg.AUTO_SIGN), "auto-sign", "S", "add -S flag to git commit")
	bindAndAddBoolFlag(CommitCmd, string(pkg.CAPITALIZE_TITLE), "capitalize-title", "capitalize the title")
	bindAndAddBoolFlag(CommitCmd, string(pkg.SCOPE_PROMPT), "scope-prompt", "enable scope prompt")
	bindAndAddBoolFlag(CommitCmd, string(pkg.BODY_PROMPT), "body-prompt", "enable body prompt")
}
