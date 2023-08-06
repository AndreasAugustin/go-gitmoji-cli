package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/spf13/cobra"
	"os"
)

var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     pkg.ProgramName,
	Version: pkg.Version,
	Short:   "Cli to help managing gitmoji commit messages",
	Long:    `See https://gitmoji.dev/ for more information about Gitmoji`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(pkg.InitConfig)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "verbose logging")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		pkg.ToggleDebug(debug)
		return nil
	}

}
