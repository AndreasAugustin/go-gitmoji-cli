package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ListCommitTypesCmd = &cobra.Command{
	Use:   "commit-types",
	Short: "List all the available commit types",
	Long:  "The list from conventional commits is used",
	PreRun: func(cmd *cobra.Command, args []string) {
		programNameFigure()
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list commit-types called")
		spin := ui.NewSpinner()
		spin.Run()
		defer func() { spin.Stop() }()
		defaultTypes := pkg.DefaultCommitTypes()
		spin.Stop()
		listSettings := ui.ListSettings{Title: "Commit types", IsShowStatusBar: true, IsFilteringEnabled: true}
		selectedDefaultType := ui.ListRun(listSettings, defaultTypes)
		log.Debugf("selected %s", selectedDefaultType)
	},
}

var ListGitmojisCmd = &cobra.Command{
	Use:   "gitmojis",
	Short: "List all the available gitmojis",
	Long:  fmt.Sprintf(`The list is queried from the api %s.`, pkg.DefaultGitmojiApiUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list gitmojis called")
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

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the available gitmojis",
	Long:  fmt.Sprintf(`The list is queried from the api %s.`, pkg.DefaultGitmojiApiUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("list called")
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
	ListCmd.AddCommand(ListGitmojisCmd)
	ListCmd.AddCommand(ListCommitTypesCmd)
}
