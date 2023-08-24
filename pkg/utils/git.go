package utils

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"strings"
)

var GetGitRepoRootDirectoryPath = getRootDirectoryPath
var GetGitRepoHooksDirectory = getHooksDirectory

func getRootDirectoryPath() (gitDirPath string, err error) {
	result, err := RunCommand("git rev-parse --show-toplevel")

	if err != nil {
		baseErr := err.Error()
		errMsg := "unexpected error encountered while trying to determine the git repo root directory. Error details: " + baseErr
		return result, errors.New(errMsg)
	} else if len(result) == 0 {
		errMsg := "got an unexpected result for the git repo root directory"
		return result, errors.New(errMsg)
	}
	gitDirPath = strings.TrimSuffix(result, "\n")
	return gitDirPath, err
}

func getHooksDirectory() (string, error) {
	hooksDir, err := RunCommand("git rev-parse --git-path hooks")
	if err != nil {
		baseErr := err.Error()
		errMsg := "unexpected error encountered while trying to determine the git repo hooks directory. Error details: " + baseErr
		return "", errors.New(errMsg)
	}

	hooksDir = strings.TrimSuffix(hooksDir, "\r\n")
	hooksDir = strings.TrimSuffix(hooksDir, "\n")

	return hooksDir, nil
}

func Commit(commitCommand string) {
	resultOutput, err := RunCommand(commitCommand)
	if err != nil {
		log.Errorf("Error while commiting %s", err)
		log.Fatalf("result output: %s", resultOutput)
	}
	log.Infof("commit done: %s", resultOutput)
}

func BuildGitCommitCommandStr(isAutoAdd bool, isSignCommit bool, title string, body string) (string, error) {
	var str strings.Builder
	str.WriteString("git commit ")
	if isAutoAdd {
		str.WriteString("-a ")
	}
	if isSignCommit {
		str.WriteString("-S ")
	}
	if title == "" {
		return "", errors.New("the title must not be empty")
	}
	str.WriteString("-m ")
	str.WriteString(title)
	str.WriteString(" ")
	if body != "" {
		str.WriteString("-m ")
		str.WriteString(body)
	}
	return str.String(), nil
}
