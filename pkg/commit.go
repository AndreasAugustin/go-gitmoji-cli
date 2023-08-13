package pkg

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"strings"
)

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
