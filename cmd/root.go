package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var verbose string

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

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setupLogs(os.Stdout, verbose); err != nil {
			return err
		}
		return nil
	}

	rootCmd.PersistentFlags().StringVarP(&verbose, "verbosity", "v", logrus.WarnLevel.String(), "Log level (debug, info, warn, error, fatal, panic)")
}

func setupLogs(out io.Writer, level string) error {
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	logrus.Debugf("Set log level to %s", lvl)
	return nil
}
