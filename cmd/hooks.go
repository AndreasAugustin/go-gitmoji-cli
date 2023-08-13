package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
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
		log.Infof("args: %v", args)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks called")
	},
}

func init() {
	RootCmd.AddCommand(HooksCmd)
	HooksCmd.AddCommand(HooksInitCmd)
	HooksCmd.AddCommand(HooksRemoveCmd)
	HooksCmd.PersistentFlags().BoolVar(&hook, "hook", false, "used when the git hook is installed")
}
