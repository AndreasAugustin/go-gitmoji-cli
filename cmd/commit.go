package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type CommitFlagName string

const (
	SCOPE       CommitFlagName = "scope"
	DESC        CommitFlagName = "desc"
	TYPE        CommitFlagName = "type"
	BODY        CommitFlagName = "body"
	IS_BREAKING CommitFlagName = "is-breaking"
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
		title := pkg.BuildCommitTitle(
			extractMessageForFlagName(TYPE, inputsRes),
			extractMessageForFlagName(SCOPE, inputsRes),
			isBreaking,
			extractMessageForFlagName(DESC, inputsRes),
			selectedGitmoji,
			pkg.ConfigInstance)
		_body := extractMessageForFlagName(BODY, inputsRes)
		log.Debugf("complete title: %s", title)
		if isDryRun {
			log.Infof("The commit title: %s", title)
			log.Infof("The commit body: %s", _body)
		} else {
			pkg.ExecuteCommit(title, _body, pkg.ConfigInstance)
		}

	},
}

func init() {
	RootCmd.AddCommand(CommitCmd)
	var a []string
	CommitCmd.PersistentFlags().StringSliceVarP(&commitMsg, "message", "m", a, "The commit message. Can be repeated")
	CommitCmd.PersistentFlags().BoolVarP(&isDryRun, "dry-run", "n", false, "dry run: just output the commit message without doing a commit")
	CommitCmd.PersistentFlags().StringVar(&_type, string(TYPE), "", "add the type")
	CommitCmd.PersistentFlags().StringVar(&scope, string(SCOPE), "", "add a scope")
	CommitCmd.PersistentFlags().StringVar(&desc, string(DESC), "", "add a description")
	CommitCmd.PersistentFlags().StringVar(&body, string(BODY), "", "add the commit message body")
	CommitCmd.PersistentFlags().BoolVar(&isBreaking, string(IS_BREAKING), false, "set if the commit is a breaking change")
}

func extractMessageForFlagName(flagName CommitFlagName, inputsRes []ui.TextInputRes) string {
	for _, res := range inputsRes {
		if res.Label == string(flagName) {
			return res.Value
		}
	}
	return ""
	//return CommitCmd.PersistentFlags().Lookup(string(flagName)).Value.String()
}

func buildTextInputsData() []ui.TextInputData {
	var textInputsData = []ui.TextInputData{{Placeholder: "type", Charlimit: 64, Label: string(TYPE), InitialValue: _type}}

	if pkg.ConfigInstance.ScopePrompt {
		textInputsData = append(textInputsData, ui.TextInputData{Placeholder: "scope", Charlimit: 64, Label: string(SCOPE), InitialValue: scope})
	}

	textInputsData = append(textInputsData, ui.TextInputData{Placeholder: "description", Charlimit: 64, Label: string(DESC), InitialValue: desc})

	if pkg.ConfigInstance.BodyPrompt {
		textInputsData = append(textInputsData, ui.TextInputData{Placeholder: "body", Charlimit: 250, Label: string(BODY), InitialValue: body})
	}

	return textInputsData
}
