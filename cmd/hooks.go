package cmd

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var hooksRemoveCmd = &cobra.Command{
	Use:   "rm",
	Short: fmt.Sprintf("remove git hooks for %s", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks init called")
		err := pkg.RemoveAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var hooksInitCmd = &cobra.Command{
	Use:   "init",
	Short: fmt.Sprintf("initialize git hooks for %s", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks init called")
		err := pkg.CreateAllHookFiles()
		if err != nil {
			log.Error(err)
			return
		}
	},
}

var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: fmt.Sprintf("Manage %s commit hooks", pkg.ProgramName),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hooks called")
		//var b = []byte("#!/usr/bin/env bash\n# gitmoji as a commit hook\nif npx -v >&/dev/null\nthen\nexec < /dev/tty\n  npx -c \"gitmoji --hook $1 $2\"\nelse\nexec < /dev/tty\n  gitmoji --hook $1 $2\nfi")
		//err := os.WriteFile("info.txt", b, 0644)
		//if err != nil {
		//	log.Fatal(err)
		//}
	},
}

func init() {
	rootCmd.AddCommand(hooksCmd)
	hooksCmd.AddCommand(hooksInitCmd)
	hooksCmd.AddCommand(hooksRemoveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
