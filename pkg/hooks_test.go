package pkg_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

func hooksTestGitmojis(t *testing.T) []pkg.Gitmoji {
	gitmojis := pkg.Gitmojis{}

	err := json.Unmarshal([]byte(testGitmojisStr), &gitmojis)
	assert.NoError(t, err)
	return gitmojis.Gitmojis
}

var expHookFileScript = `#!/bin/sh
# go-gitmoji-cli
# version ` + pkg.Version + `

hookName=` + "`basename \"$0\"`" + `
gitParams="$*"

GIT_CMD=$(ps -ocommand= -p $PPID);
IS_AMEND=$(echo "${GIT_CMD}" | grep -e '--amend');
IS_REBASE=$(echo "${GIT_CMD}" | grep -e 'rebase');

if [ -n "$IS_AMEND" ]; then
  echo "Using amend, skipping use of go-gitmoji-cli"
  exit 0;
elif [ -n "$IS_REBASE" ]; then
  echo "Using rebase, skipping use of go-gitmoji-cli"
  exit 0;
fi

if command -v go-gitmoji-cli >/dev/null 2>&1; then
  go-gitmoji-cli hooks --hook $gitParams
else
  echo "Can't find go-gitmoji-cli, skipping $hookName hook"
  echo "You can reinstall it using 'go get -u github.com/AndreasAugustin/go-gitmoji-cli' or delete this hook"
fi`

var expGitHooks = [...]string{
	"prepare-commit-msg",
}

var expHookFileContents = []byte(expHookFileScript)

const gitHooksPath = "/usr/foo/repos/bar/.git/hooks"

func TestCreateAllHookFilesReturnsCorrectErrorOnHooksDirError(t *testing.T) {
	origGitFunc := utils.GetGitRepoHooksDirectory
	defer func() { utils.GetGitRepoHooksDirectory = origGitFunc }()
	utils.GetGitRepoHooksDirectory = func() (string, error) {
		return "", errors.New("")
	}

	assert.Equal(t, pkg.ErrInvalidGitHooksDirectoryPath, pkg.CreateAllHookFiles())
}

func TestCreateAllHookFilesReturnsCorrectErrorWhenSomeHooksNotCreated(t *testing.T) {
	expErrorHooks := [2]string{"prepare-commit-msg"}
	expErrMsg := fmt.Sprintf("encountered an error while attempting to create one or more hook files. did not create hooks: %v", expErrorHooks)
	origGitFunc := utils.GetGitRepoHooksDirectory
	defer func() { utils.GetGitRepoHooksDirectory = origGitFunc }()
	utils.GetGitRepoHooksDirectory = func() (string, error) {
		return gitHooksPath, nil
	}
	originalWriteFile := utils.WriteFile
	defer func() { utils.WriteFile = originalWriteFile }()
	utils.WriteFile = func(filePath string, contents []byte) error {
		hook := strings.TrimPrefix(filePath, filepath.Join(gitHooksPath))
		hook = hook[1:]

		if hook == "prepare-commit-msg" {
			return errors.New("")
		}

		return nil
	}

	assert.Error(t, pkg.CreateAllHookFiles(), expErrMsg)
}

func TestCreateAllHookFilesCreatesCorrectHooks(t *testing.T) {
	var actHookPaths []string
	origGitFunc := utils.GetGitRepoHooksDirectory
	defer func() { utils.GetGitRepoHooksDirectory = origGitFunc }()
	utils.GetGitRepoHooksDirectory = func() (string, error) {
		return gitHooksPath, nil
	}
	originalWriteFile := utils.WriteFile
	defer func() { utils.WriteFile = originalWriteFile }()
	utils.WriteFile = func(filePath string, contents []byte) error {
		if !bytes.Equal(contents, expHookFileContents) {
			hook := strings.TrimPrefix(filePath, filepath.Join(gitHooksPath))
			hook = hook[1:]
			t.Errorf("Incorrect script contents used for hook '%s'. Expected: %s, but got: %s", hook, string(expHookFileContents), string(contents))
		}
		actHookPaths = append(actHookPaths, filePath)
		return nil
	}
	err := pkg.CreateAllHookFiles()

	assert.NoError(t, err)
	assert.Equal(t, len(actHookPaths), len(expGitHooks))
	for i, actHookPath := range actHookPaths {
		expHookPath := filepath.Join(gitHooksPath, expGitHooks[i])
		assert.Equal(t, expHookPath, actHookPath)
	}
}

func TestReadCommitEditMsgWithTitleBodyFooterEquExpected(t *testing.T) {
	gitmojis := hooksTestGitmojis(t)
	parsedMessages, err := pkg.ReadAndParseCommitEditMsg("./test_data/COMMIT_EDITMSG_title_header_footer", gitmojis)
	assert.NoError(t, err)

	exp := pkg.ParsedMessages{
		Body:       "here is the body",
		Footer:     "here is the footer",
		Type:       "feat",
		Scope:      "hooks",
		IsBreaking: true,
		Desc:       "this is my message"}
	assert.Equal(t, exp, *parsedMessages)
}

func TestReadCommitEditMsgNoMsgEquExpected(t *testing.T) {
	gitmojis := hooksTestGitmojis(t)
	parsedMessages, err := pkg.ReadAndParseCommitEditMsg("./test_data/COMMIT_EDITMSG_no_msg_provided", gitmojis)
	assert.NoError(t, err)

	exp := pkg.ParsedMessages{
		Body:       "",
		Footer:     "",
		Type:       "",
		Scope:      "",
		IsBreaking: false,
		Desc:       "",
	}
	assert.Equal(t, exp, *parsedMessages)
}
