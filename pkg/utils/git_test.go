package utils_test

import (
	"errors"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"

	"github.com/stretchr/testify/assert"
	"testing"
)

const expGitRootCommand = "git rev-parse --show-toplevel"
const expGitHooksCommand = "git rev-parse --git-path hooks"

func TestGetRootDirectoryPathHandlesErrorCorrectly(t *testing.T) {
	errMsgDetails := "not a git repo"
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo root directory. Error details: " + errMsgDetails
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		return "", errors.New(errMsgDetails)
	}
	defer func() { utils.RunCommand = origRunCommand }()

	_, err := utils.GetGitRepoRootDirectoryPath()
	assert.Error(t, err, expectedErrMsg)
}

func TestGetRootDirectoryPathHandlesEmptyDirectoryCorrectly(t *testing.T) {
	expectedErrMsg := "got an unexpected result for the git repo root directory."
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		return "", nil
	}
	defer func() { utils.RunCommand = origRunCommand }()

	gitDir, err := utils.GetGitRepoRootDirectoryPath()
	assert.Error(t, err, expectedErrMsg)
	assert.Equal(t, "", gitDir)
}

func TestGetGitRepoRootDirectoryPathReturnsDirectoryCorrectly(t *testing.T) {
	expectedGitDir := "/usr/repos/go-gitmoji-cli"
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		return expectedGitDir, nil
	}
	defer func() { utils.RunCommand = origRunCommand }()

	gitDir, err := utils.GetGitRepoRootDirectoryPath()
	assert.Nil(t, err)
	assert.Equal(t, expectedGitDir, gitDir)
}

func TestGetGitRepoRootDirectoryUsesCorrectCommand(t *testing.T) {
	var actualCmd string
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		actualCmd = cmd
		return "", nil
	}
	defer func() { utils.RunCommand = origRunCommand }()
	utils.GetGitRepoRootDirectoryPath()
	assert.Equal(t, expGitRootCommand, actualCmd, "Used incorrect command. Expected: %s, but got: %s", expGitRootCommand, actualCmd)
}

func TestGetHooksDirectoryReturnsErrorWhenCommandFails(t *testing.T) {
	errMsgDetails := "ouch"
	expectedErrMsg := "unexpected error encountered while trying to determine the git repo hooks directory. Error details: " + errMsgDetails
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		return "", errors.New(errMsgDetails)
	}
	defer func() { utils.RunCommand = origRunCommand }()

	_, err := utils.GetGitRepoRootDirectoryPath()
	assert.Error(t, err, expectedErrMsg)
}

func TestGetHooksDirectoryUsesCorrectCommand(t *testing.T) {
	var actualCmd string
	origRunCommand := utils.RunCommand
	utils.RunCommand = func(cmd string) (string, error) {
		actualCmd = cmd
		return "", nil
	}
	defer func() { utils.RunCommand = origRunCommand }()
	utils.GetGitRepoHooksDirectory()
	assert.Equal(t, expGitHooksCommand, actualCmd, "Used incorrect command. Expected: %s, but got: %s", expGitHooksCommand, actualCmd)
}

func TestBuildGitCommitCommandStrNoAutoAddNoSigningEqualsExpected(t *testing.T) {
	commitTitle := "feat(test): :smile: this is my title"
	commitBody := "this is just the body\njust for testing"
	res, err := utils.BuildGitCommitCommandStr(false, false, commitTitle, commitBody)
	assert.NoError(t, err)
	expected := fmt.Sprintf("git commit -m %s -m %s", commitTitle, commitBody)
	assert.Equal(t, expected, res)
}

func TestBuildGitCommitCommandStrAutoAddNoSigningEqualsExpected(t *testing.T) {
	commitTitle := "feat(test): :smile: this is my title"
	commitBody := "this is just the body\njust for testing"
	res, err := utils.BuildGitCommitCommandStr(true, false, commitTitle, commitBody)
	assert.NoError(t, err)
	expected := fmt.Sprintf("git commit -a -m %s -m %s", commitTitle, commitBody)
	assert.Equal(t, expected, res)
}

func TestBuildGitCommitCommandStrNoAutoAddSigningEqualsExpected(t *testing.T) {
	commitTitle := "feat(test): :smile: this is my title"
	commitBody := "this is just the body\njust for testing"
	res, err := utils.BuildGitCommitCommandStr(false, true, commitTitle, commitBody)
	assert.NoError(t, err)
	expected := fmt.Sprintf("git commit -S -m %s -m %s", commitTitle, commitBody)
	assert.Equal(t, expected, res)
}

func TestBuildGitCommitCommandStrAutoAddSigningEqualsExpected(t *testing.T) {
	commitTitle := "feat(test): :smile: this is my title"
	commitBody := "this is just the body\njust for testing"
	res, err := utils.BuildGitCommitCommandStr(true, true, commitTitle, commitBody)
	assert.NoError(t, err)
	expected := fmt.Sprintf("git commit -a -S -m %s -m %s", commitTitle, commitBody)
	assert.Equal(t, expected, res)
}
