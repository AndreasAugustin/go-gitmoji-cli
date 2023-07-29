package pkg

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"path/filepath"
)

var gitHooks = [...]string{
	"prepare-commit-msg",
}

var hookFileScript = `#!/bin/sh
# go-gitmoji-cli
# version ` + Version + `

hookName=` + "`basename \"$0\"`" + `
gitParams="$*"

if command -v go-gitmoji-cli >/dev/null 2>&1; then
  go-gitmoji-cli commit "$gitParams"
else
  echo "Can't find go-gitmoji-cli, skipping $hookName hook"
  echo "You can reinstall it using 'go get -u github.com/AndreasAugustin/go-gitmoji-cli' or delete this hook"
fi`

var hookFileContents = []byte(hookFileScript)
var ErrInvalidGitHooksDirectoryPath = errors.New("invalid git hooks directory path")

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
	return nil
}
