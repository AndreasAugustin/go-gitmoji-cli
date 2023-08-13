package pkg_test

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/stretchr/testify/assert"
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
