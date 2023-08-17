package pkg

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strings"
)

type ParsedMessages struct {
	Type       string
	Scope      string
	Desc       string
	IsBreaking bool
	Gitmoji    Gitmoji
	Body       string
	Footer     string
}

type InitialCommitValues struct {
	Type  string
	Scope string
	Desc  string
	Body  string
}

type CommitValues struct {
	Title string
	Body  string
}

type TextInputRes struct {
	Value string
	Label string
}

type CommitFlagName string

const (
	SCOPE       CommitFlagName = "scope"
	DESC        CommitFlagName = "desc"
	TYPE        CommitFlagName = "type"
	BODY        CommitFlagName = "body"
	IS_BREAKING CommitFlagName = "is-breaking"
)

func BuildInitialCommitValues(_type string, scope string, desc string, body string, commitMsg []string) InitialCommitValues {
	var stringEmptyOrOption = func(input string, option string) string {
		if input != "" {
			return input
		}
		return option
	}

	parsedMessages, err := ParseCommitMessages(commitMsg)
	if err != nil {
		log.Fatalf("parsing the messages did not work %s", err)
	}
	return InitialCommitValues{Type: stringEmptyOrOption(_type, parsedMessages.Type), Scope: stringEmptyOrOption(scope, parsedMessages.Scope), Desc: stringEmptyOrOption(desc, parsedMessages.Desc), Body: stringEmptyOrOption(body, parsedMessages.Body)}
}

func ParseCommitMessages(messages []string) (*ParsedMessages, error) {
	if len(messages) == 0 || len(messages) > 3 {
		return nil, errors.New("the amount of messages is to low or to high")
	}

	var body = ""
	if len(messages) >= 2 {
		body = messages[1]
	}

	var footer = ""
	if len(messages) >= 3 {
		footer = messages[2]
	}

	splitDesc := strings.SplitN(messages[0], ":", 2)
	if len(splitDesc) == 1 {
		return &ParsedMessages{Desc: messages[0], Body: body, Footer: footer}, nil
	}

	typeScopeEventualBreaking := splitDesc[0]
	typeScope := strings.ReplaceAll(typeScopeEventualBreaking, "!", "")

	reScope := regexp.MustCompile(`\((.*?)\)`)
	matchScope := reScope.FindAllString(typeScope, 1)
	var _type string
	var scope string
	if matchScope == nil {
		_type = typeScope
	} else {
		scopeWithBraces := matchScope[0]
		_type = strings.Replace(typeScope, scopeWithBraces, "", 1)
		scopeWithBraces = strings.TrimLeft(scopeWithBraces, "(")
		scope = strings.TrimRight(scopeWithBraces, ")")
	}

	descEventualEmoji := splitDesc[1]
	reEmoji := regexp.MustCompile(`:(.*?):`)
	matchEmoji := reEmoji.FindAllString(descEventualEmoji, 1)
	var parsedMessage = ParsedMessages{
		IsBreaking: strings.ContainsAny(typeScopeEventualBreaking, "!"),
		Type:       _type,
		Scope:      scope,
		Body:       body,
		Footer:     footer,
	}
	if matchEmoji == nil {
		parsedMessage.Desc = descEventualEmoji
	} else {
		matchedCode := matchEmoji[0]
		parsedMessage.Desc = strings.TrimLeft(strings.ReplaceAll(descEventualEmoji, matchedCode, ""), " ")
		config, err := GetCurrentConfig()
		if err != nil {
			log.Warnf("error while getting config %s", err)
			return &parsedMessage, nil
		}
		gitmojis := GetGitmojis(config)
		foundGitmoji := FindGitmoji(matchedCode, gitmojis.Gitmojis)
		if foundGitmoji == nil {
			log.Warnf("no gitmoji for %s has been found", matchedCode)
			return &parsedMessage, nil
		}
		parsedMessage.Gitmoji = *foundGitmoji
	}

	return &parsedMessage, nil
}

func CreateMessage(inputsRes []TextInputRes, selectedGitmoji Gitmoji, initialCommitValues InitialCommitValues, config Config, isBreaking bool) CommitValues {
	title := BuildCommitTitle(
		extractMessageForFlagName(TYPE, inputsRes, initialCommitValues),
		extractMessageForFlagName(SCOPE, inputsRes, initialCommitValues),
		isBreaking,
		extractMessageForFlagName(DESC, inputsRes, initialCommitValues),
		selectedGitmoji,
		config)
	_body := extractMessageForFlagName(BODY, inputsRes, initialCommitValues)
	return CommitValues{Title: title, Body: _body}
}

func ExecuteCommit(title string, body string, config Config) {
	commitCmd, err := utils.BuildGitCommitCommandStr(config.AutoAdd, config.AutoSign, title, body)
	if err != nil {
		log.Fatalf("error building commit message: %s", err)
	}
	log.Debugf("commit cmd: %s", commitCmd)
	utils.Commit(commitCmd)
}

func BuildCommitTitle(_type string, scope string, isBreaking bool, desc string, gitmoji Gitmoji, config Config) string {

	var s strings.Builder
	s.WriteString("\"")
	s.WriteString(_type)

	if scope != "" {
		s.WriteString(fmt.Sprintf("(%s)", scope))
	}

	if isBreaking {
		s.WriteString("!")
	}

	s.WriteString(fmt.Sprintf(": %s ", gitmojiToString(gitmoji, config)))
	s.WriteString(eventualCapitalizeTitle(desc, config))
	s.WriteString("\"")
	return s.String()
}

func gitmojiToString(gitmoji Gitmoji, config Config) string {
	if config.EmojiFormat == CODE {
		return gitmoji.Code
	} else {
		return gitmoji.Emoji
	}
}

func eventualCapitalizeTitle(title string, config Config) string {
	if config.CapitalizeTitle {
		return strings.ToUpper(title)
	}
	return title
}

func extractMessageForFlagName(flagName CommitFlagName, inputsRes []TextInputRes, initialCommitValues InitialCommitValues) string {
	for _, res := range inputsRes {
		if res.Label == string(flagName) {
			return res.Value
		}
	}
	return initialCommitValues.extractInitialValue(flagName)
}

func (i InitialCommitValues) BuildTextInputsData(config Config) []TextInputData {
	var textInputsData = []TextInputData{{Placeholder: "type", Charlimit: 64, Label: string(TYPE), InitialValue: i.Type}}

	if config.ScopePrompt {
		textInputsData = append(textInputsData, TextInputData{Placeholder: "scope", Charlimit: 64, Label: string(SCOPE), InitialValue: i.Scope})
	}

	textInputsData = append(textInputsData, TextInputData{Placeholder: "description", Charlimit: 64, Label: string(DESC), InitialValue: i.Desc})

	if config.BodyPrompt {
		textInputsData = append(textInputsData, TextInputData{Placeholder: "body", Charlimit: 250, Label: string(BODY), InitialValue: i.Body})
	}

	return textInputsData
}

func (i InitialCommitValues) extractInitialValue(flagName CommitFlagName) string {
	switch flagName {
	case TYPE:
		return i.Type
	case BODY:
		return i.Body
	case DESC:
		return i.Desc
	case SCOPE:
		return i.Scope
	default:
		log.Debugf("flagname %s not found", string(flagName))
		return ""

	}
}
