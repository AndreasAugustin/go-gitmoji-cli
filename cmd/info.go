package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get some relevant information",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		globalConfigDir, _ := utils.GetGlobalConfigDir(pkg.ProgramName)
		cacheDir, _ := utils.GetCacheDir(pkg.ProgramName)
		log.Debug("info called")
		log.Infof("program name: %s", pkg.ProgramName)
		log.Infof("version: %s", pkg.Version)
		log.Info("It is possible to store the configuration within the repository or globally")
		log.Infof("The global config path: %s", globalConfigDir)
		log.Infof("The gitmoji infomration is queried from the internet at the gitmoji defined in the config %s", pkg.ConfigInstance.GitmojisUrl)
		log.Info("and cached")
		log.Infof("Gitmoji cache dir: %s", cacheDir)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
