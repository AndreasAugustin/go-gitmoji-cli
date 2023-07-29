package pkg

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

const ProgramName = "go-gitmoji-cli"
const Version = "v0.1.0"

type ConfigEnum string

const EnvPrefix = "GO_GITMOJI_CLI"
const configName = ".gitmojirc"
const configPath = "."

const (
	AUTO_ADD         ConfigEnum = "autoAdd"
	EMOJI_FORMAT     ConfigEnum = "emojiFormat"
	SCOPE_PROMPT     ConfigEnum = "scopePrompt"
	GITMOJIS_URL     ConfigEnum = "gitmojisUrl"
	MESSAGE_PROMPT   ConfigEnum = "messagePrompt"
	CAPITALIZE_TITLE ConfigEnum = "capitalizeTitle"
)

type EmojiCommitFormats string

const (
	CODE  EmojiCommitFormats = "code"
	EMOJI EmojiCommitFormats = "emoji"
)

type Config struct {
	Autoadd         bool               `mapstructure:"autoAdd" json:"autoadd"`
	EmojiFormat     EmojiCommitFormats `mapstructure:"emojiFormat" json:"emojiFormat"`
	ScopePrompt     bool               `mapstructure:"scopePrompt" json:"scopePrompt"`
	MessagePrompt   bool               `mapstructure:"messagePrompt" json:"messagePrompt"`
	CapitalizeTitle bool               `mapstructure:"capitalizeTitle" json:"capitalizeTitle"`
	GitmojisUrl     string             `mapstructure:"gitmojisUrl" json:"gitmojisUrl"`
}

var ConfigInstance Config

func InitConfig() {
	var err error
	ConfigInstance, err = LoadConfig(configPath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func AddEnvPrefix(suffix string) string {
	return fmt.Sprintf("%s_%s", EnvPrefix, suffix)
}

func LoadConfig(configPath string) (config Config, err error) {
	viper.SetDefault("autoAdd", false)
	viper.SetDefault("emojiFormat", CODE)
	viper.SetDefault("scopePrompt", false)
	viper.SetDefault("messagePrompt", true)
	viper.SetDefault("capitalizeTitle", true)
	viper.SetDefault("gitmojisUrl", "https://gitmoji.dev/api/gitmojis")

	//viper.SetEnvPrefix(EnvPrefix)
	//viper.BindEnv(AddEnvPrefix("AUTO_ADD"))
	//viper.BindEnv(AddEnvPrefix("EMOJI_FORMAT"))

	//viper.SetConfigType("env")
	if configPath != "" {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(configName)
		if err = viper.ReadInConfig(); err != nil {
			// It's okay if there isn't a config file
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				return
			}
		}
	}

	//viper.EnvKeyReplacer(strings.NewReplacer("_", ""))
	//viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return config, err
}
