package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var HooksRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: fmt.Sprintf("remove git hooks for %s", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks called")
	},
}

func init() {
	RootCmd.AddCommand(HooksCmd)
	HooksCmd.AddCommand(HooksInitCmd)
	HooksCmd.AddCommand(HooksRemoveCmd)
}
