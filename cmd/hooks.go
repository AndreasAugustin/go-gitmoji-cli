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
var hookCommitMessageFile string

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
	Args: func(cmd *cobra.Command, args []string) error {
		log.Infof("args: %+v", args)
		if hook {
			hookCommitMessageFile = args[0]
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks called")
		if hook {
			hookCommit()
		}
	},
}

func init() {
	RootCmd.AddCommand(HooksCmd)
	HooksCmd.AddCommand(HooksInitCmd)
	HooksCmd.AddCommand(HooksRemoveCmd)
	HooksCmd.PersistentFlags().BoolVar(&hook, "hook", false, "used when the git hook is installed")
}

// TODO(anau) add messages
func hookCommit() {
	log.Debug("hook --hooks called")
	log.Debug(commitMsg)
	spin := ui.NewSpinner()
	spin.Run()
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
	gitmojis := pkg.GetGitmojis(config)
	spin.Stop()
	initialCommitValues := pkg.InitialCommitValues{}
	listSettings := ui.ListSettings{IsShowStatusBar: true, IsFilteringEnabled: true, Title: "Gitmojis"}
	selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
	log.Debugf("selected gitmoji %s", selectedGitmoji)
	textInputsData := initialCommitValues.BuildTextInputsData(config)
	inputsRes := ui.TextInputsRun("please add", textInputsData)

	commitValues := pkg.CreateMessage(inputsRes, selectedGitmoji, initialCommitValues, config, isBreaking)

	log.Debugf("complete title: %s", commitValues.Title)

	commitMsg := fmt.Sprintf("%s \n \n %s", commitValues.Title, commitValues.Body)
	err = utils.WriteFile(hookCommitMessageFile, []byte(commitMsg))
	if err != nil {
		log.Fatalf("error writing commit hook file %s", err)
	}
}
