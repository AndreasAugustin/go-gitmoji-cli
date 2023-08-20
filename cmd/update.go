package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var UpdateGitmojisCmd = &cobra.Command{
	Use: "gitmojis",
	PreRun: func(cmd *cobra.Command, args []string) {
		programNameFigure()
	},
	Short: fmt.Sprintf("update the local gitmoji database %s", pkg.ProgramName),
	Long:  fmt.Sprintf(`Update the gitmojis local cache from %s.`, pkg.DefaultGitmojiApiUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("update gitmojis called")
		spin := ui.NewSpinner()
		spin.Run()
		defer func() {
			spin.Stop()
		}()
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("get current config issue, %s", err)
		}
		gitmojis := pkg.GetGitmojis(config)
		pkg.CacheGitmojis(gitmojis)
		log.Info("done updating the cached gitmojis file")
	},
}

var updateCmd = &cobra.Command{
	Use: "update",
	PreRun: func(cmd *cobra.Command, args []string) {
		programNameFigure()
	},
	Short: fmt.Sprintf("Manage %s updates", pkg.ProgramName),
	Long:  `Update command for the cli.`,
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(UpdateGitmojisCmd)
}
