package pkg_test

import (
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"strconv"
	"testing"
)

func addEnvPrefix(suffix string) string {
	return fmt.Sprintf("%s_%s", pkg.EnvPrefix, suffix)
}

func getTestConfigPath(T *testing.T) string {
	wDir, err := os.Getwd()
	assert.NoError(T, err)
	return path.Join(wDir, "tmp", "test")
}

func TestConfigDefaultValuesEqualsExpected(t *testing.T) {
	var configIsInitPers = pkg.ConfigIsInit
	pkg.ConfigIsInit = true
	defer func() {
		pkg.ConfigIsInit = configIsInitPers
	}()
	err := pkg.LoadConfig([]string{})

	assert.NoError(t, err)
	config, err := pkg.GetCurrentConfig()
	assert.NoError(t, err)
	expected := pkg.Config{
		EmojiFormat:           pkg.CODE,
		AutoAdd:               false,
		ScopePrompt:           false,
		BodyPrompt:            false,
		CapitalizeTitle:       false,
		GitmojisUrl:           "https://gitmoji.dev/api/gitmojis",
		UseDefaultGitMessages: true,
	}
	assert.Equal(t, expected, config)
}

func TestConfigEnvVariablesEqualsExpected(t *testing.T) {
	var autoadd = true
	var emojiFormat = pkg.EMOJI
	var scopePrompt = true
	var bodyPrompt = true
	var capitalizeTitle = true
	var gitmojisUrl = "http://foo.bar"
	var autoSign = true
	var useDefaultGitMessages = false
	var debug = true

	var configIsInitPers = pkg.ConfigIsInit
	pkg.ConfigIsInit = true
	defer func() {
		pkg.ConfigIsInit = configIsInitPers
	}()

	t.Setenv(addEnvPrefix(string(pkg.AUTO_ADD)), strconv.FormatBool(autoadd))
	t.Setenv(addEnvPrefix(string(pkg.EMOJI_FORMAT)), string(emojiFormat))
	t.Setenv(addEnvPrefix(string(pkg.SCOPE_PROMPT)), strconv.FormatBool(scopePrompt))
	t.Setenv(addEnvPrefix(string(pkg.BODY_PROMPT)), strconv.FormatBool(bodyPrompt))
	t.Setenv(addEnvPrefix(string(pkg.CAPITALIZE_TITLE)), strconv.FormatBool(capitalizeTitle))
	t.Setenv(addEnvPrefix(string(pkg.GITMOJIS_URL)), gitmojisUrl)
	t.Setenv(addEnvPrefix(string(pkg.AUTO_SIGN)), strconv.FormatBool(autoSign))
	t.Setenv(addEnvPrefix(string(pkg.USE_DEFAULT_GIT_MESSAGES)), strconv.FormatBool(useDefaultGitMessages))
	t.Setenv(addEnvPrefix(string(pkg.DEBUG)), strconv.FormatBool(debug))
	err1 := pkg.LoadConfig([]string{})
	assert.NoError(t, err1)
	config, err := pkg.GetCurrentConfig()
	assert.NoError(t, err)
	expected := pkg.Config{
		EmojiFormat:           emojiFormat,
		AutoAdd:               autoadd,
		AutoSign:              autoSign,
		ScopePrompt:           scopePrompt,
		BodyPrompt:            bodyPrompt,
		CapitalizeTitle:       capitalizeTitle,
		GitmojisUrl:           gitmojisUrl,
		UseDefaultGitMessages: useDefaultGitMessages,
		Debug:                 debug,
	}
	assert.Equal(t, expected, config)
}

func TestConfigConfigFileEqualsExpected(t *testing.T) {
	var configIsInitPers = pkg.ConfigIsInit
	pkg.ConfigIsInit = true
	defer func() {
		pkg.ConfigIsInit = configIsInitPers
	}()
	err1 := pkg.LoadConfig([]string{"./test_data"})
	assert.NoError(t, err1)

	config, err := pkg.GetCurrentConfig()
	assert.NoError(t, err)
	expected := pkg.Config{
		EmojiFormat:           pkg.EMOJI,
		AutoAdd:               true,
		ScopePrompt:           true,
		BodyPrompt:            true,
		CapitalizeTitle:       true,
		GitmojisUrl:           "http://from.file",
		UseDefaultGitMessages: false,
		Debug:                 true,
	}
	assert.Equal(t, expected, config)
}

func TestWriteGlobalConfigAndReadEqualsExpected(t *testing.T) {
	var configIsInitPers = pkg.ConfigIsInit
	pkg.ConfigIsInit = true
	defer func() {
		pkg.ConfigIsInit = configIsInitPers
	}()
	err1 := pkg.LoadConfig([]string{"./test_data"})
	assert.NoError(t, err1)
	config, err := pkg.GetCurrentConfig()
	assert.NoError(t, err)
	getGlobalConfigDirPersist := utils.GetGlobalConfigDir

	utils.GetGlobalConfigDir = func(programName string) (string, error) {
		testConfigPath := getTestConfigPath(t)
		return path.Join(testConfigPath, programName), nil
	}

	pkg.UpdateConfig(config, true)
	globalConfigPath := path.Join(getTestConfigPath(t), pkg.ProgramName)
	err1 = pkg.LoadConfig([]string{globalConfigPath})
	assert.NoError(t, err1)
	configFromGlobal, err := pkg.GetCurrentConfig()
	assert.NoError(t, err)
	defer func() {
		utils.GetGlobalConfigDir = getGlobalConfigDirPersist
		os.RemoveAll(getTestCacheProgramPath(t))
	}()
	assert.NoError(t, err)
	assert.Equal(t, config, configFromGlobal)
}
