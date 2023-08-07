package pkg_test

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strconv"
	"testing"
)

func getTestConfigPath(T *testing.T) string {
	wDir, err := os.Getwd()
	assert.NoError(T, err)
	return path.Join(wDir, "tmp", "test")
}

func TestConfigDefaultValuesEqualsExpected(t *testing.T) {
	config, _ := pkg.LoadConfig([]string{})
	expected := pkg.Config{EmojiFormat: pkg.CODE, Autoadd: false, ScopePrompt: false, MessagePrompt: true, CapitalizeTitle: true, GitmojisUrl: "https://gitmoji.dev/api/gitmojis"}
	assert.Equal(t, expected, config)
}

func TestConfigEvnVariablesEqualsExpected(t *testing.T) {
	var autoadd = true
	var emojiFormat = pkg.EMOJI
	var scopePrompt = true
	var messagePrompt = false
	var capitalizeTitle = false
	var gitmojisUrl = "http://foo.bar"

	t.Setenv(pkg.AddEnvPrefix("AUTO_ADD"), strconv.FormatBool(autoadd))
	t.Setenv(pkg.AddEnvPrefix("EMOJI_FORMAT"), string(emojiFormat))
	t.Setenv(pkg.AddEnvPrefix("SCOPE_PROMPT"), strconv.FormatBool(scopePrompt))
	t.Setenv(pkg.AddEnvPrefix("MESSAGE_PROMPT"), strconv.FormatBool(messagePrompt))
	t.Setenv(pkg.AddEnvPrefix("CAPITALIZE_TITLE"), strconv.FormatBool(capitalizeTitle))
	t.Setenv(pkg.AddEnvPrefix("GITMOJIS_URL"), gitmojisUrl)
	config, _ := pkg.LoadConfig([]string{})
	expected := pkg.Config{EmojiFormat: emojiFormat, Autoadd: autoadd, ScopePrompt: scopePrompt, MessagePrompt: messagePrompt, CapitalizeTitle: capitalizeTitle, GitmojisUrl: gitmojisUrl}
	assert.Equal(t, expected, config)
}

func TestConfigConfigFileEqualsExpected(t *testing.T) {
	config, _ := pkg.LoadConfig([]string{"./test_data"})
	expected := pkg.Config{EmojiFormat: pkg.EMOJI, Autoadd: true, ScopePrompt: true, MessagePrompt: false, CapitalizeTitle: false, GitmojisUrl: "http://from.file"}
	assert.Equal(t, expected, config)
}

func TestWriteGlobalConfigAndReadEqualsExpected(t *testing.T) {
	config, err := pkg.LoadConfig([]string{"./test_data"})

	assert.NoError(t, err)
	getGlobalConfigDirPersist := utils.GetGlobalConfigDir

	utils.GetGlobalConfigDir = func(programName string) (string, error) {
		testConfigPath := getTestConfigPath(t)
		return path.Join(testConfigPath, programName), nil
	}

	pkg.UpdateConfig(config, true)
	globalConfigPath := path.Join(getTestConfigPath(t), pkg.ProgramName)
	configFromGlobal, err := pkg.LoadConfig([]string{globalConfigPath})
	defer func() {
		utils.GetGlobalConfigDir = getGlobalConfigDirPersist
		os.RemoveAll(getTestCacheProgramPath(t))
	}()
	assert.NoError(t, err)
	assert.Equal(t, config, configFromGlobal)
}
