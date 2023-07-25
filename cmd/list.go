package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
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
		fmt.Println("list called")
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
		s.Start()                                                   // Start the spinner
		//time.Sleep(4 * time.Second)                                 // Run for some time to simulate work
		res, err := http.Get(pkg.ConfigInstance.GitmojisUrl)
		if err != nil {
			fmt.Println("error", err)
		}

		if res.Body != nil {
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {

				}
			}(res.Body)
		}

		body, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Fatal(readErr)
		}
		fmt.Println(body)
		gitmojis := pkg.Gitmojis{}
		jsonErr := json.Unmarshal(body, &gitmojis)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		printEmojis(gitmojis)
		s.Stop()
	},
}

func printEmojis(gitmojis pkg.Gitmojis) {
	for _, gitmoji := range gitmojis.Gitmojis {
		blue := color.New(color.FgBlue).SprintFunc()
		fmt.Println(gitmoji.Emoji, blue(gitmoji.Code), gitmoji.Description)
	}
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
