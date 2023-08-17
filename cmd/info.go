package cmd

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get some relevant information",
	Long:  `Get some information like config or cache directory on your OS where the cli is writing configuration or cache.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := pkg.GetCurrentConfig()
		if err != nil {
			log.Fatalf("error while getting config %s", err)
		}
		globalConfigDir, _ := utils.GetGlobalConfigDir(pkg.ProgramName)
		cacheDir, _ := utils.GetCacheDir(pkg.ProgramName)
		log.Debug("info called")
		log.Infof("program name: %s", pkg.ProgramName)
		log.Infof("version: %s", pkg.Version)
		log.Info("It is possible to store the configuration within the repository or globally")
		log.Infof("The global config path: %s", globalConfigDir)
		log.Infof("The gitmoji information is queried from the internet at the gitmoji defined in the config %s", config.GitmojisUrl)
		log.Info("and cached")
		log.Infof("Gitmoji cache dir: %s", cacheDir)
	},
}

func init() {
	RootCmd.AddCommand(InfoCmd)
}
