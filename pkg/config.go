package pkg

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
)

var ConfigIsInit = false

const ProgramName = "go-gitmoji-cli"

// Version as provided by goreleaser.
var Version = ""

// CommitSHA as provided by goreleaser.
var CommitSHA = ""

const EnvPrefix = "GO_GITMOJI_CLI"
const configName = ".gitmojirc"
const configPath = "."

func InitConfig() {
	if ConfigIsInit {
		return
	}

	globalConfigPath, err := getGlobalConfigPath()
	if err != nil {
		log.Fatalf("No global config path found")
	}
	err = LoadConfig([]string{configPath, globalConfigPath})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ConfigIsInit = true
}

func getGlobalConfigPath() (string, error) {
	return utils.GetUserConfigDirCreateIfNotExists(ProgramName)
}

func LoadConfig(configPaths []string) (err error) {
	log.Debug("Load config")
	viper.SetDefault(string(AUTO_ADD), false)
	viper.SetDefault(string(AUTO_SIGN), false)
	viper.SetDefault(string(EMOJI_FORMAT), string(CODE))
	viper.SetDefault(string(SCOPE_PROMPT), false)
	viper.SetDefault(string(BODY_PROMPT), false)
	viper.SetDefault(string(CAPITALIZE_TITLE), false)
	viper.SetDefault(string(GITMOJIS_URL), DefaultGitmojiApiUrl)
	viper.SetDefault(string(IS_DEFAULT_MERGE_MESSAGE), true)
	viper.SetDefault(string(DEBUG), false)

	viper.SetEnvPrefix(EnvPrefix)

	viper.SetConfigType("json")
	if len(configPaths) != 0 {
		for _, val := range configPaths {
			viper.AddConfigPath(val)
		}

		viper.SetConfigName(configName)
		if expErr := viper.ReadInConfig(); expErr != nil {
			log.Debug("issue reading config")
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if !errors.As(expErr, &configFileNotFoundError) {
				return
			}
		}
	}

	viper.AutomaticEnv()

	return
}

func GetCurrentConfig() (config Config, err error) {
	log.Debug("Get current config")
	if !ConfigIsInit {
		InitConfig()
	}
	err = viper.Unmarshal(&config)
	log.Debugf("Config %+v", config)
	return
}

func UpdateConfig(config Config, isGlobalConfig bool) {
	viper.Set(string(AUTO_ADD), config.AutoAdd)
	viper.Set(string(AUTO_SIGN), config.AutoSign)
	viper.Set(string(EMOJI_FORMAT), string(config.EmojiFormat))
	viper.Set(string(SCOPE_PROMPT), config.ScopePrompt)
	viper.Set(string(BODY_PROMPT), config.BodyPrompt)
	viper.Set(string(CAPITALIZE_TITLE), config.CapitalizeTitle)
	viper.Set(string(GITMOJIS_URL), config.GitmojisUrl)
	viper.Set(string(IS_DEFAULT_MERGE_MESSAGE), config.IsDefaultMergeMessage)
	viper.Set(string(DEBUG), config.Debug)

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

func DefaultCommitTypes() []CommitType {
	return []CommitType{
		{Type: "feat", Desc: "A  new feature"},
		{Type: "fix", Desc: "A bug fix"},
		{Type: "docs", Desc: "Documentation only changes"},
		{Type: "style", Desc: "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)"},
		{Type: "refactor", Desc: "A code change that neither fixes a bug nor adds a feature"},
		{Type: "perf", Desc: "A code change that improves performance"},
		{Type: "test", Desc: "Adding missing tests or correcting existing tests"},
		{Type: "build", Desc: "Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)"},
		{Type: "ci", Desc: "Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)"},
		{Type: "chore", Desc: "Other changes that don't modify src or test files"},
		{Type: "revert", Desc: "Reverts a previous commit"},
	}
}
