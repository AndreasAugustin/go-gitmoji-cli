package pkg

import "fmt"

// DefaultGitmojiUrl The URL to reach the gitmoji website
const DefaultGitmojiUrl = "https://gitmoji.dev/"

// DefaultGitmojiApiUrl The Url to reach the grimoji API
const DefaultGitmojiApiUrl = "https://gitmoji.dev/api/gitmojis"

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
	AUTO_ADD                 ConfigEnum = "AUTO_ADD"
	AUTO_SIGN                ConfigEnum = "AUTO_SIGN"
	AUTO_SIGNATURE           ConfigEnum = "AUTO_SIGNATURE"
	EMOJI_FORMAT             ConfigEnum = "EMOJI_FORMAT"
	SCOPE_PROMPT             ConfigEnum = "SCOPE_PROMPT"
	GITMOJIS_URL             ConfigEnum = "GITMOJIS_URL"
	BODY_PROMPT              ConfigEnum = "BODY_PROMPT"
	CAPITALIZE_TITLE         ConfigEnum = "CAPITALIZE_TITLE"
	USE_DEFAULT_GIT_MESSAGES ConfigEnum = "USE_DEFAULT_GIT_MESSAGES"
	DEBUG                    ConfigEnum = "DEBUG"
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
	AutoAdd               bool               `mapstructure:"AUTO_ADD" json:"auto_add"`
	AutoSign              bool               `mapstructure:"AUTO_SIGN" json:"auto_sign"`
	AutoSignature         bool               `mapstructure:"AUTO_SIGNATURE" json:"auto_signature"`
	EmojiFormat           EmojiCommitFormats `mapstructure:"EMOJI_FORMAT" json:"emoji_format"`
	ScopePrompt           bool               `mapstructure:"SCOPE_PROMPT" json:"scope_prompt"`
	BodyPrompt            bool               `mapstructure:"BODY_PROMPT" json:"body_prompt"`
	CapitalizeTitle       bool               `mapstructure:"CAPITALIZE_TITLE" json:"capitalize_title"`
	GitmojisUrl           string             `mapstructure:"GITMOJIS_URL" json:"gitmojis_url"`
	UseDefaultGitMessages bool               `mapstructure:"USE_DEFAULT_GIT_MESSAGES" json:"use_default_git_messages"`
	Debug                 bool               `mapstructure:"DEBUG" json:"debug"`
}

type TextInputData struct {
	Placeholder  string
	Charlimit    int
	InitialValue string
	Label        string
}

type CommitType struct {
	Type string `mapstructure:"TYPE" json:"type"`
	Desc string `mapstructure:"DESC" json:"desc"`
}

func (i CommitType) FilterValue() string {
	return i.Type + i.Desc
}
func (i CommitType) Title() string {
	return i.Type
}
func (i CommitType) Description() string {
	return i.Desc
}

type CommitMessageSource string

const (
	MESSAGE  CommitMessageSource = "message"
	TEMPLATE CommitMessageSource = "template"
	MERGE    CommitMessageSource = "merge"
	SQUASH   CommitMessageSource = "squash"
	COMMIT   CommitMessageSource = "commit"
)
