package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ListCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available gitmojis",
	Long:  fmt.Sprintf(`The list is queried from the api %s.`, pkg.DefaultGitmojiApiUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list called")
		spin := ui.NewSpinner()
		spin.Run()
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("get current config issue, %s", err)
		}
		gitmojis := pkg.GetGitmojis(config)
		spin.Stop()
		listSettings := ui.ListSettings{Title: "Gitmojis", IsShowStatusBar: true, IsFilteringEnabled: true}
		selectedGitmoji := ui.ListRun(listSettings, gitmojis.Gitmojis)
		log.Debugf("selected %s", selectedGitmoji)
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
