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
	expected := pkg.Config{EmojiFormat: pkg.CODE, AutoAdd: false, ScopePrompt: false, BodyPrompt: false, CapitalizeTitle: false, GitmojisUrl: "https://gitmoji.dev/api/gitmojis"}
	assert.Equal(t, expected, config)
}

func TestConfigEvnVariablesEqualsExpected(t *testing.T) {
	var autoadd = true
	var emojiFormat = pkg.EMOJI
	var scopePrompt = true
	var bodyPrompt = true
	var capitalizeTitle = true
	var gitmojisUrl = "http://foo.bar"
	var autoSign = true

	t.Setenv(pkg.AddEnvPrefix(string(pkg.AUTO_ADD)), strconv.FormatBool(autoadd))
	t.Setenv(pkg.AddEnvPrefix(string(pkg.EMOJI_FORMAT)), string(emojiFormat))
	t.Setenv(pkg.AddEnvPrefix(string(pkg.SCOPE_PROMPT)), strconv.FormatBool(scopePrompt))
	t.Setenv(pkg.AddEnvPrefix(string(pkg.BODY_PROMPT)), strconv.FormatBool(bodyPrompt))
	t.Setenv(pkg.AddEnvPrefix(string(pkg.CAPITALIZE_TITLE)), strconv.FormatBool(capitalizeTitle))
	t.Setenv(pkg.AddEnvPrefix(string(pkg.GITMOJIS_URL)), gitmojisUrl)
	t.Setenv(pkg.AddEnvPrefix(string(pkg.AUTO_SIGN)), strconv.FormatBool(autoSign))
	config, _ := pkg.LoadConfig([]string{})
	expected := pkg.Config{EmojiFormat: emojiFormat, AutoAdd: autoadd, AutoSign: autoSign, ScopePrompt: scopePrompt, BodyPrompt: bodyPrompt, CapitalizeTitle: capitalizeTitle, GitmojisUrl: gitmojisUrl}
	assert.Equal(t, expected, config)
}

func TestConfigConfigFileEqualsExpected(t *testing.T) {
	config, _ := pkg.LoadConfig([]string{"./test_data"})
	expected := pkg.Config{EmojiFormat: pkg.EMOJI, AutoAdd: true, ScopePrompt: true, BodyPrompt: false, CapitalizeTitle: false, GitmojisUrl: "http://from.file"}
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
