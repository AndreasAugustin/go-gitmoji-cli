package pkg

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/common-nighthawk/go-figure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

const ProgramName = "go-gitmoji-cli"

// Version as provided by goreleaser.
var Version = ""

// CommitSHA as provided by goreleaser.
var CommitSHA = ""

func ProgramNameFigure() {
	programNameFigure := figure.NewColorFigure(ProgramName, "cybermedium", "purple", true)

	programNameFigure.Print()
}

const EnvPrefix = "GO_GITMOJI_CLI"
const configName = ".gitmojirc"
const configPath = "."

var ConfigInstance Config

func getGlobalConfigPath() (string, error) {
	return utils.GetUserConfigDirCreateIfNotExists(ProgramName)
}

func InitConfig() {

	globalConfigPath, err := getGlobalConfigPath()
	if err != nil {
		log.Fatalf("No global config path found")
	}
	ConfigInstance, err = LoadConfig([]string{configPath, globalConfigPath})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func AddEnvPrefix(suffix string) string {
	return fmt.Sprintf("%s_%s", EnvPrefix, suffix)
}

func LoadConfig(configPaths []string) (config Config, err error) {
	viper.SetDefault(string(AUTO_ADD), false)
	viper.SetDefault(string(AUTO_SIGN), false)
	viper.SetDefault(string(EMOJI_FORMAT), string(CODE))
	viper.SetDefault(string(SCOPE_PROMPT), false)
	viper.SetDefault(string(BODY_PROMPT), false)
	viper.SetDefault(string(CAPITALIZE_TITLE), false)
	viper.SetDefault(string(GITMOJIS_URL), DefaultGitmojiApiUrl)

	viper.SetEnvPrefix(EnvPrefix)

	viper.SetConfigType("json")
	if len(configPaths) != 0 {
		for _, val := range configPaths {
			viper.AddConfigPath(val)
		}

		viper.SetConfigName(configName)
		if err = viper.ReadInConfig(); err != nil {
			// It's okay if there isn't a config file
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if !errors.As(err, &configFileNotFoundError) {
				return
			}
		}
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}

func UpdateConfig(config Config, isGlobalConfig bool) {
	viper.Set(string(AUTO_ADD), config.AutoAdd)
	viper.Set(string(EMOJI_FORMAT), string(config.EmojiFormat))
	viper.Set(string(SCOPE_PROMPT), config.ScopePrompt)
	viper.Set(string(BODY_PROMPT), config.BodyPrompt)
	viper.Set(string(CAPITALIZE_TITLE), config.CapitalizeTitle)
	viper.Set(string(GITMOJIS_URL), config.GitmojisUrl)

	pathToWrite := configFilePath(isGlobalConfig)

	err := viper.WriteConfigAs(pathToWrite)
	if err != nil {
		log.Fatalf("writting config did not work %s", err)
	}
	log.Debugf("Write to path %s finished", pathToWrite)
}

func configFilePath(isGlobalConfig bool) string {
	name := fmt.Sprintf("%s.json", configName)
	if isGlobalConfig {
		globalConfigPath, err := getGlobalConfigPath()
		if err != nil {
			log.Fatalf("No global config path found")
		}
		return path.Join(globalConfigPath, name)
	} else {
		return path.Join(configPath, name)
	}
}
