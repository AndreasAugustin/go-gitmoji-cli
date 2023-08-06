package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var updateGitmojisCmd = &cobra.Command{
	Use:   "gitmojis",
	Short: fmt.Sprintf("update the local gitmoji database %s", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("update gitmojis called")
		gitmojis := pkg.GetGitmojis()
		pkg.CacheGitmojis(gitmojis)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("Manage %s updates", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("update called")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(updateGitmojisCmd)
}
