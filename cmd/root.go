package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/spf13/cobra"
	"os"
)

var debug bool

var RootCmd = &cobra.Command{
	Use: pkg.ProgramName,
	//Version: pkg.Version,
	Short: "Cli to help managing gitmoji commit messages",
	Long:  fmt.Sprintf(`See %s for more information about Gitmoji`, pkg.DefaultGitmojiUrl),
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	if //goland:noinspection GoBoolExpressions
	len(pkg.CommitSHA) >= 7 {
		vt := RootCmd.VersionTemplate()
		RootCmd.SetVersionTemplate(vt[:len(vt)-1] + " (" + pkg.CommitSHA[0:7] + ")\n")
	}
	if pkg.Version == "" {
		pkg.Version = "unknown (built from source)"
	}
	RootCmd.Version = pkg.Version

	pkg.ProgramNameFigure()
	cobra.OnInitialize(pkg.InitConfig)

	RootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "verbose logging")

	RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		pkg.ToggleDebug(debug)
		return nil
	}

}
