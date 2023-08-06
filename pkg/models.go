package pkg

import "fmt"

type Gitmoji struct {
	Emoji  string `mapstructure:"emoji" json:"emoji"`
	Entity string `mapstructure:"entity" json:"entity"`
	Code   string `mapstructure:"code" json:"code"`
	Desc   string `mapstructure:"description" json:"description"`
	Name   string `mapstructure:"name" json:"name"`
	Semver string `mapstructure:"semver" json:"semver"`
}

func (i Gitmoji) FilterValue() string { return i.Name + i.Desc }
func (i Gitmoji) Title() string       { return fmt.Sprintf("%s %s", i.Emoji, i.Code) }
func (i Gitmoji) Description() string { return i.Desc }

type Gitmojis struct {
	Gitmojis []Gitmoji `json:"gitmojis"`
}

type YesNo string

const (
	YES YesNo = "Yes"
	NO  YesNo = "NO"
)

func (i YesNo) FilterValue() string { return string(i) }
func (i YesNo) Title() string       { return string(i) }
func (i YesNo) Description() string {
	if i == YES {
		return "Accept"
	} else {
		return "Deny"
	}
}

type ConfigEnum string

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

func (i EmojiCommitFormats) FilterValue() string {
	if i == CODE {
		return "shortcode"
	} else {
		return "unicode"
	}
}
func (i EmojiCommitFormats) Title() string {
	if i == CODE {
		return "shortcode"
	} else {
		return "unicode"
	}
}
func (i EmojiCommitFormats) Description() string {
	if i == CODE {
		return "shortcode format e.g. :smile:"
	} else {
		return "unicode format e.g. 😄"
	}
}

type Config struct {
	Autoadd         bool               `mapstructure:"AUTO_ADD" json:"AUTO_ADD"`
	EmojiFormat     EmojiCommitFormats `mapstructure:"EMOJI_FORMAT" json:"EMOJI_FORMAT"`
	ScopePrompt     bool               `mapstructure:"SCOPE_PROMPT" json:"SCOPE_PROMPT"`
	MessagePrompt   bool               `mapstructure:"MESSAGE_PROMPT" json:"MESSAGE_PROMPT"`
	CapitalizeTitle bool               `mapstructure:"CAPITALIZE_TITLE" json:"CAPITALIZE_TITLE"`
	GitmojisUrl     string             `mapstructure:"GITMOJIS_URL" json:"GITMOJIS_URL"`
}
