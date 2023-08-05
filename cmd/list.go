package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available gitmojis",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list called")
		spin := ui.NewSpinner()
		spin.Run()
		gitmojis := pkg.GetGitmojis()
		spin.Stop()
		listSettings := ui.ListSettings{Title: "Gitmojis", IsShowStatusBar: true, IsFilteringEnabled: true}
		selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
		log.Debugf("selected %s", selectedGitmoji)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
