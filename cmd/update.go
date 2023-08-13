package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var UpdateGitmojisCmd = &cobra.Command{
	Use:   "gitmojis",
	Short: fmt.Sprintf("update the local gitmoji database %s", pkg.ProgramName),
	Long:  fmt.Sprintf(`Update the gitmojis local cache from %s.`, pkg.DefaultGitmojiApiUrl),
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("update gitmojis called")
		gitmojis := pkg.GetGitmojis()
		pkg.CacheGitmojis(gitmojis)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: fmt.Sprintf("Manage %s updates", pkg.ProgramName),
	Long:  `Update command for the cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("update called")
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(UpdateGitmojisCmd)
}
