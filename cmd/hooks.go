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
		spin := ui.NewSpinner()
		spin.Run()
		defer func() { spin.Stop() }()
		err := pkg.RemoveAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("The hook is now removed")
		log.Infof("happy coding %s", "🚀")
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
		defer func() {
			spin.Stop()
		}()
		err := pkg.CreateAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("The hook is now initialized")
		log.Infof("happy coding %s", "🚀")
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
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("get current config issue, %s", err)
		}
		log.Debug("hooks called")
		log.Debugf("run: %v", args)
		if hook {
			if len(args) == 0 {
				log.Fatalf("len(args) must not be 0 when using pre-commit-message hook https://git-scm.com/docs/githooks#_prepare_commit_msg")
			}
			hookCommitMessageFile := args[0]
			log.Debugf("hook message file %s", hookCommitMessageFile)
			if len(args) > 1 {
				hookCommitMsgSource := args[1]
				log.Debugf("hook message source %s", hookCommitMsgSource)
				if config.UseDefaultGitMessages {
					log.Debugf("Default git messages used")
					switch hookCommitMsgSource {
					case string(pkg.MERGE):
						log.Infof("Merge commit, using default git message")
						return
					case string(pkg.SQUASH):
						log.Infof("Squash commit, using default git message")
						return
					}
				}
			}

			hookCommit(hookCommitMessageFile, config)
		} else {
			err := cmd.Help()
			if err != nil {
				log.Fatalf("issue with the help command %s", err)
				return
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(HooksCmd)
	HooksCmd.AddCommand(HooksInitCmd)
	HooksCmd.AddCommand(HooksRemoveCmd)
	HooksCmd.PersistentFlags().BoolVar(&hook, "hook", false, "used when the git hook is installed")
}

func hookCommit(commitMsgFile string, config pkg.Config) {
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
		return
	}

	gitmojis := pkg.GetGitmojis(config)

	parsedMessages, err := pkg.ReadAndParseCommitEditMsg(commitMsgFile, gitmojis.Gitmojis)

	if err != nil {
		log.Fatalf("issue reading and parsing the commit msg file %s", err)
	}

	initialCommitValues := pkg.InitialCommitValues{
		Type:  parsedMessages.Type,
		Scope: parsedMessages.Scope,
		Desc:  parsedMessages.Desc,
		Body:  parsedMessages.Body,
	}
	spin.Stop()
	commitValues := ui.CommitPrompt(config, gitmojis.Gitmojis, initialCommitValues, isBreaking)
	commitMsg := fmt.Sprintf("%s \n \n %s", commitValues.Title, commitValues.Body)
	err = utils.WriteFile(commitMsgFile, []byte(commitMsg))
	if err != nil {
		log.Fatalf("error writing commit hook file %s", err)
	}
}
