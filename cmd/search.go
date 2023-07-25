package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/briandowns/spinner"
	"github.com/ktr0731/go-fuzzyfinder"
	"time"

	"github.com/spf13/cobra"
)

func search(gitmojis []pkg.Gitmoji) pkg.Gitmoji {
	//if query == "" {
	//	return gitmojis
	//}
	idx, err := fuzzyfinder.Find(gitmojis, func(i int) string {
		return fmt.Sprintf("%s - %s - %s", gitmojis[i].Emoji, gitmojis[i].Name, gitmojis[i].Description)
	},
		fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
			if i == -1 {
				return ""
			}
			return pkg.ColoredRepr(gitmojis[i])
		}))
	if err != nil {
		fmt.Println(err)
	}
	return gitmojis[idx]
}

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search gitmojis",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("search called")
		fmt.Println("list called")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
		s.Start()                                                   // Start the spinner
		gitmojis := pkg.GetGitmojis()
		gitmoji := search(gitmojis.Gitmojis)
		fmt.Println("selected", pkg.ColoredRepr(gitmoji))
		s.Stop()
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
