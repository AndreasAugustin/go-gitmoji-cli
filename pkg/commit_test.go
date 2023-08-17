package pkg_test

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func buildCommitTestConfig(emojiFormat pkg.EmojiCommitFormats) pkg.Config {
	return pkg.Config{
		AutoAdd:         false,
		AutoSign:        false,
		EmojiFormat:     emojiFormat,
		ScopePrompt:     false,
		BodyPrompt:      false,
		CapitalizeTitle: false,
		GitmojisUrl:     "",
	}
}

var gitmoji = pkg.Gitmoji{
	Emoji:  "🍻",
	Entity: "\u0026#x1f37b;",
	Code:   ":beers:",
	Desc:   "Testing stuff.",
	Name:   "beers",
	Semver: "",
}

func TestBuildCommitTitleCodeFormatIsNotBreakingNoScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := ""
	isBreaking := false
	desc := "test description"
	config := buildCommitTestConfig(pkg.CODE)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := "\"feat: :beers: test description\""
	assert.Equal(t, exp, title)
}

func TestBuildCommitTitleEmojiFormatIsNotBreakingNoScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := ""
	isBreaking := false
	desc := "test description"
	config := buildCommitTestConfig(pkg.EMOJI)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := fmt.Sprintf("\"feat: %s test description\"", "🍻")
	assert.Equal(t, exp, title)
}

func TestBuildCommitTitleCodeFormatIsNotBreakingWithScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "test"
	isBreaking := false
	desc := "test description"
	config := buildCommitTestConfig(pkg.CODE)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := fmt.Sprintf("\"feat(test): %s test description\"", ":beers:")
	assert.Equal(t, exp, title)
}

func TestBuildCommitTitleEmojiFormatIsNotBreakingWithScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "test"
	isBreaking := false
	desc := "test description"
	config := buildCommitTestConfig(pkg.EMOJI)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := fmt.Sprintf("\"feat(test): %s test description\"", "🍻")
	assert.Equal(t, exp, title)
}

func TestBuildCommitTitleCodeFormatIsBreakingWithScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "test"
	isBreaking := true
	desc := "test description"
	config := buildCommitTestConfig(pkg.CODE)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := fmt.Sprintf("\"feat(test)!: %s test description\"", ":beers:")
	assert.Equal(t, exp, title)
}

func TestBuildCommitTitleEmojiFormatIsBreakingWithScopeEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "test"
	isBreaking := true
	desc := "test description"
	config := buildCommitTestConfig(pkg.EMOJI)
	title := pkg.BuildCommitTitle(_type, scope, isBreaking, desc, gitmoji, config)
	exp := fmt.Sprintf("\"feat(test)!: %s test description\"", "🍻")
	assert.Equal(t, exp, title)
}

func TestParseMessagesLenWrongHasError(t *testing.T) {
	var empty []string
	var tooHuge = []string{"head", "body", "footer", "not right"}
	var casesArgs = [][]string{empty, tooHuge}

	var testFkt = func(messages []string) func(t *testing.T) {
		return func(t *testing.T) {
			_, err := pkg.ParseCommitMessages(messages)
			assert.Error(t, err, "the amount of messages is to low or to high")
		}
	}

	for i, arg := range casesArgs {
		t.Run(strconv.Itoa(i), testFkt(arg))
	}
}

func TestParseMessagesHasHeaderDescEqualsExp(t *testing.T) {
	desc := "only have description"
	exp := pkg.ParsedMessages{Desc: desc}
	res, err := pkg.ParseCommitMessages([]string{desc})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}

func TestParseMessagesHasHeaderDescBodyFooterEqualsExp(t *testing.T) {
	desc := "only have description"
	body := "I am a body"
	footer := "I am a footer"
	exp := pkg.ParsedMessages{Desc: desc, Body: body, Footer: footer}
	res, err := pkg.ParseCommitMessages([]string{desc, body, footer})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}

func TestParseMessagesHasHeaderTypeDescBodyFooterEqualsExp(t *testing.T) {
	_type := "feat"
	desc := "feat and description"
	header := fmt.Sprintf("%s:%s", _type, desc)
	body := "I am a body"
	footer := "I am a footer"
	exp := pkg.ParsedMessages{Desc: desc, Body: body, Footer: footer, Type: _type}
	res, err := pkg.ParseCommitMessages([]string{header, body, footer})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}

func TestParseMessagesHasHeaderTypeDescIsBreakingBodyFooterEqualsExp(t *testing.T) {
	_type := "feat"
	desc := "feat and description"
	header := fmt.Sprintf("%s!:%s", _type, desc)
	body := "I am a body"
	footer := "I am a footer"
	exp := pkg.ParsedMessages{IsBreaking: true, Desc: desc, Body: body, Footer: footer, Type: _type}
	res, err := pkg.ParseCommitMessages([]string{header, body, footer})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}

func TestParseMessagesHasHeaderTypeDescIsBreakingScopeBodyFooterEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "api"
	desc := "feat and description"
	header := fmt.Sprintf("%s(%s)!:%s", _type, scope, desc)
	body := "I am a body"
	footer := "I am a footer"
	exp := pkg.ParsedMessages{Scope: scope, IsBreaking: true, Desc: desc, Body: body, Footer: footer, Type: _type}
	res, err := pkg.ParseCommitMessages([]string{header, body, footer})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}

func TestParseMessagesHasHeaderTypeDescEmojiIsBreakingScopeBodyFooterEqualsExp(t *testing.T) {
	_type := "feat"
	scope := "api"
	descEmoji := ":rocket:"
	desc := "(#18) feat and description"
	descCombined := fmt.Sprintf("%s %s", descEmoji, desc)
	header := fmt.Sprintf("%s(%s):%s", _type, scope, descCombined)
	body := "I am a body"
	footer := "I am a footer"
	expEmoji := pkg.Gitmoji{Emoji: "🚀", Entity: "&#x1f680;", Code: ":rocket:", Desc: "Deploy stuff.", Name: "rocket", Semver: ""}
	exp := pkg.ParsedMessages{Scope: scope, IsBreaking: false, Desc: desc, Body: body, Footer: footer, Type: _type, Gitmoji: expEmoji}
	res, err := pkg.ParseCommitMessages([]string{header, body, footer})
	assert.NoError(t, err)
	assert.Equal(t, exp, *res)
}
