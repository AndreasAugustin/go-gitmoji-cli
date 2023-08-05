package pkg

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
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
	AUTO_ADD         ConfigEnum = "AUTO_ADD"
	EMOJI_FORMAT     ConfigEnum = "EMOJI_FORMAT"
	SCOPE_PROMPT     ConfigEnum = "SCOPE_PROMPT"
	GITMOJIS_URL     ConfigEnum = "GITMOJIS_URL"
	MESSAGE_PROMPT   ConfigEnum = "MESSAGE_PROMPT"
	CAPITALIZE_TITLE ConfigEnum = "CAPITALIZE_TITLE"
)

type EmojiCommitFormats string

const (
	CODE  EmojiCommitFormats = "code"
	EMOJI EmojiCommitFormats = "emoji"
)

type Config struct {
	Autoadd         bool               `mapstructure:"AUTO_ADD" json:"AUTO_ADD"`
	EmojiFormat     EmojiCommitFormats `mapstructure:"EMOJI_FORMAT" json:"EMOJI_FORMAT"`
	ScopePrompt     bool               `mapstructure:"SCOPE_PROMPT" json:"SCOPE_PROMPT"`
	MessagePrompt   bool               `mapstructure:"MESSAGE_PROMPT" json:"MESSAGE_PROMPT"`
	CapitalizeTitle bool               `mapstructure:"CAPITALIZE_TITLE" json:"CAPITALIZE_TITLE"`
	GitmojisUrl     string             `mapstructure:"GITMOJIS_URL" json:"GITMOJIS_URL"`
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
	viper.SetDefault(string(AUTO_ADD), false)
	viper.SetDefault(string(EMOJI_FORMAT), string(CODE))
	viper.SetDefault(string(SCOPE_PROMPT), false)
	viper.SetDefault(string(MESSAGE_PROMPT), true)
	viper.SetDefault(string(CAPITALIZE_TITLE), true)
	viper.SetDefault(string(GITMOJIS_URL), "https://gitmoji.dev/api/gitmojis")

	viper.SetEnvPrefix(EnvPrefix)
	//viper.BindEnv(AddEnvPrefix(string(AUTO_ADD)))
	//viper.BindEnv(AddEnvPrefix(string(EMOJI_FORMAT)))
	//viper.BindEnv(AddEnvPrefix())

	//viper.SetConfigType("env")
	viper.SetConfigType("json")
	if configPath != "" {
		viper.AddConfigPath(configPath)
		viper.SetConfigName(configName)
		if err = viper.ReadInConfig(); err != nil {
			// It's okay if there isn't a config file
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if !errors.As(err, &configFileNotFoundError) {
				return
			}
		}
	}

	//viper.EnvKeyReplacer(strings.NewReplacer("_", ""))
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return config, err
}

func UpdateConfig(config Config) {
	viper.Set(string(AUTO_ADD), config.Autoadd)
	viper.Set(string(EMOJI_FORMAT), string(config.EmojiFormat))
	viper.Set(string(SCOPE_PROMPT), config.ScopePrompt)
	viper.Set(string(MESSAGE_PROMPT), config.MessagePrompt)
	viper.Set(string(CAPITALIZE_TITLE), config.CapitalizeTitle)
	viper.Set(string(GITMOJIS_URL), config.GitmojisUrl)

	err := viper.WriteConfig()
	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Info("config file not present yet. Writting config file")
			err = viper.SafeWriteConfig()
			if err == nil {
				log.Info("Config file written")
				return
			}
		}
		log.Fatalf("writting config did not work %s", err)
	}
}
