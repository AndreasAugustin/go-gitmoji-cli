package pkg

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"regexp"
	"strings"
)

var gitHooks = [...]string{
	"prepare-commit-msg",
}

var hookFileScript = `#!/bin/sh
# go-gitmoji-cli
# version ` + Version + `

hookName=` + "`basename \"$0\"`" + `
gitParams="$*"

IS_AMEND=$(ps -ocommand= -p $PPID | grep -e '--amend');

if [ -n "$IS_AMEND" ]; then
  echo "Using amend, skipping use of go-gitmoji-cli"
  exit 0;
fi

if command -v go-gitmoji-cli >/dev/null 2>&1; then
  go-gitmoji-cli hooks --hook $gitParams
else
  echo "Can't find go-gitmoji-cli, skipping $hookName hook"
  echo "You can reinstall it using 'go get -u github.com/AndreasAugustin/go-gitmoji-cli' or delete this hook"
fi`

var hookFileContents = []byte(hookFileScript)
var ErrInvalidGitHooksDirectoryPath = errors.New("invalid git hooks directory path")

func ReadAndParseCommitEditMsg(filePath string, gitmojis []Gitmoji) (*ParsedMessages, error) {
	file, err := utils.ReadFile(filePath)
	log.Debugf("file content of %s", filePath)
	log.Debugf("%s", string(file))
	if err != nil {
		return nil, err
	}
	fileStr := string(file)
	lines := regexp.MustCompile("\r?\n").Split(fileStr, -1)
	var messages []string
	for _, line := range lines {
		if line != "" && !strings.HasPrefix(line, "#") {
			messages = append(messages, line)
		}
	}
	log.Debugf("messages: %v", messages)
	if messages == nil {
		return &ParsedMessages{}, nil
	}
	return ParseCommitMessages(messages, gitmojis)
}

func CreateAllHookFiles() error {
	hooksDir, hooksErr := utils.GetGitRepoHooksDirectory()
	if hooksErr != nil {
		return ErrInvalidGitHooksDirectoryPath
	}

	var notCreatedHooks []string

	for _, hook := range gitHooks {
		hookPath := filepath.Join(hooksDir, hook)
		err := utils.WriteFile(hookPath, hookFileContents)
		if err != nil {
			notCreatedHooks = append(notCreatedHooks, hook)
		}
	}

	if len(notCreatedHooks) > 0 {
		return fmt.Errorf("encountered an error while attempting to create one or more hook files. did not create hooks: %v", notCreatedHooks)
	}

	return nil
}

func RemoveAllHookFiles() error {
	hooksDir, hooksErr := utils.GetGitRepoHooksDirectory()
	if hooksErr != nil {
		return ErrInvalidGitHooksDirectoryPath
	}

	var notRemovedHooks []string

	for _, hook := range gitHooks {
		hookPath := filepath.Join(hooksDir, hook)
		err := utils.RemoveFile(hookPath)
		if err != nil {
			notRemovedHooks = append(notRemovedHooks, hook)
		}
	}

	if len(notRemovedHooks) > 0 {
		return fmt.Errorf("encountered an error while attempting to create one or more hook files. did not create hooks: %v", notRemovedHooks)
	}

	return nil
}

func HookFilesExistent() ([]string, error) {
	hooksDir, hooksErr := utils.GetGitRepoHooksDirectory()
	if hooksErr != nil {
		return []string{}, ErrInvalidGitHooksDirectoryPath
	}

	var existentHooks []string

	for _, hook := range gitHooks {
		hookPath := filepath.Join(hooksDir, hook)
		exists := utils.FileExists(hookPath)
		if exists {
			existentHooks = append(existentHooks, hook)
		}
	}

	return existentHooks, nil
}
