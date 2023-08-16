package pkg_test

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
)

var expHookFileScript = `#!/bin/sh
# go-gitmoji-cli
# version ` + pkg.Version + `

hookName=` + "`basename \"$0\"`" + `
gitParams="$*"

if command -v go-gitmoji-cli >/dev/null 2>&1; then
  go-gitmoji-cli hooks --hook "$gitParams"
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

func TestRemoveAllHookFilesReturnsCorrectError(t *testing.T) {
	t.Skip("need to be checked")
	var expErr error
	assert.Equal(t, expErr, pkg.RemoveAllHookFiles())
}
