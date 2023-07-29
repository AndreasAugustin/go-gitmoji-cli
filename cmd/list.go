package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/ui"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
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
		time.Sleep(200 * time.Millisecond)
		//s := spinner.New()
		//s.Spinner = spinner.Dot
		//s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
		//s.View()
		gitmojis := pkg.GetGitmojis()
		spin.Stop()
		//pkg.PrintEmojis(gitmojis)
		ui.ListRun("Gitmojis", gitmojis.Gitmojis)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
