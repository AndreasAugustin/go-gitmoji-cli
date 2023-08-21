package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var hook bool

var HooksRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: fmt.Sprintf("remove git hooks for %s", pkg.ProgramName),
	Long:  `Delete the commit hooks which are created by the cli`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks rm called")
		err := pkg.RemoveAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var HooksInitCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("initialize git hooks for %s", pkg.ProgramName),
	Long:  `Install the commit hooks into the local .git/hooks/ directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks init called")
		spin := ui.NewSpinner()
		spin.Run()
		defer func() { spin.Stop() }()
		err := pkg.CreateAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var HooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: fmt.Sprintf("Manage %s commit hooks", pkg.ProgramName),
	Long:  `Manage git hooks for the cli`,
	PreRun: func(cmd *cobra.Command, args []string) {
		programNameFigure()
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("hooks called")
		log.Infof("run: %v", args)
		if hook {
			hookCommitMessageFile := args[0]
			hookCommit(hookCommitMessageFile)
		}
	},
}

func init() {
	RootCmd.AddCommand(HooksCmd)
	HooksCmd.AddCommand(HooksInitCmd)
	HooksCmd.AddCommand(HooksRemoveCmd)
	HooksCmd.PersistentFlags().BoolVar(&hook, "hook", false, "used when the git hook is installed")
}

func hookCommit(commitMsgFile string) {
	log.Debug("hook --hooks called")
	log.Debug(commitMsg)
	spin := ui.NewSpinner()
	spin.Run()
	defer func() { spin.Stop() }()
	existentHookFiles, err := pkg.HookFilesExistent()
	if err != nil {
		log.Fatalf("Error checking if hook files existent")
	}
	if len(existentHookFiles) == 0 {
		log.Infof("There are no hook files existent for %s", existentHookFiles)
		log.Infof("Please use commit command or create hooks with %s hooks init", pkg.ProgramName)
		spin.Stop()
		return
	}
	config, err := pkg.GetCurrentConfig()
	if err != nil {
		log.Fatalf("get current config issue, %s", err)
	}
	parsedMessages, err := pkg.ReadAndParseCommitEditMsg(commitMsgFile)
	if err != nil {
		log.Fatalf("issue reading and parsing the commit msg file %s", err)
	}

	gitmojis := pkg.GetGitmojis(config)
	defaultTypes := pkg.DefaultCommitTypes()

	initialCommitValues := pkg.InitialCommitValues{
		Type:  parsedMessages.Type,
		Scope: parsedMessages.Scope,
		Desc:  parsedMessages.Desc,
		Body:  parsedMessages.Body,
	}
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

	commitMsg := fmt.Sprintf("%s \n \n %s", commitValues.Title, commitValues.Body)
	err = utils.WriteFile(commitMsgFile, []byte(commitMsg))
	if err != nil {
		log.Fatalf("error writing commit hook file %s", err)
	}
}
