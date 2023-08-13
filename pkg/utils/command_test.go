package utils_test

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

type MockCommand struct {
	CombinedOutputFunc func() ([]byte, error)
}

func (m MockCommand) CombinedOutput() ([]byte, error) {
	if m.CombinedOutputFunc != nil {
		return m.CombinedOutputFunc()
	}
	return nil, nil
}

const echoScript = "echo foobar"

var echoCmd = [...]string{"sh", "-c", echoScript}
var expNumArgs = len(echoCmd)

func TestGetRunnerInfoReturnsCorrectValueOnWindows(t *testing.T) {
	const expectedRunner = "cmd.exe"
	const expectedRunnerArg = "/C"
	runner, runnerArg := utils.GetRunnerInfo("windows")
	assert.Equal(t, expectedRunner, runner)
	assert.Equal(t, expectedRunnerArg, runnerArg)
}

func TestGetRunnerInfoReturnsCorrectValueOnNonWindows(t *testing.T) {
	const expectedRunner = "sh"
	const expectedRunnerArg = "-c"
	nonWindowsOperatingSystems := []string{
		"linux",
		"darwin",
		"freebsd",
	}

	for _, os := range nonWindowsOperatingSystems {
		runner, runnerArg := utils.GetRunnerInfo(os)
		assert.Equal(t, expectedRunner, runner, "Runner was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunner, runner)
		assert.Equal(t, expectedRunnerArg, runnerArg, "Runner Arg was incorrect for OS: %s. Expected: %s, but got: %s.", os, expectedRunnerArg, runnerArg)
	}
}

func TestNewCommandUsesDirectoryWhenSpecified(t *testing.T) {
	dir := "/usr/some/other/directory"
	mockCmd := &exec.Cmd{}
	utils.OsCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	defer func() { utils.OsCommand = exec.Command }()
	assert.NotNil(t, utils.CreateCommand(dir, echoScript))
	assert.Equal(t, dir, mockCmd.Dir, "Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	numArgs := len(mockCmd.Args)
	assert.Equal(t, expNumArgs, numArgs, "Did not get correct number of command args. Expected: %d, but got: %d", expNumArgs, numArgs)
}

func TestNewCommandUsesCallingProcDirectoryWhenNotSpecified(t *testing.T) {
	dir := ""
	mockCmd := &exec.Cmd{}
	utils.OsCommand = func(name string, arg ...string) *exec.Cmd {
		mockCmd.Path = name
		mockCmd.Args = append([]string{name}, arg...)
		return mockCmd
	}
	defer func() { utils.OsCommand = exec.Command }()
	assert.NotNil(t, utils.CreateCommand(dir, echoScript))
	assert.Equal(t, dir, mockCmd.Dir, "Target directory for command was incorrect. Expected: %s, but got: %s.", dir, mockCmd.Dir)
	numArgs := len(mockCmd.Args)
	assert.Equal(t, expNumArgs, numArgs, "Did not get correct number of command args. Expected: %d, but got: %d", expNumArgs, numArgs)
}

func TestRunReturnsCorrectResults(t *testing.T) {
	mockBytes := []byte("foobar")
	createCommand := utils.CreateCommand
	utils.CreateCommand = func(directory, script string) utils.Command {
		return &MockCommand{CombinedOutputFunc: func() ([]byte, error) {
			return mockBytes, nil
		}}
	}
	defer func() { utils.CreateCommand = createCommand }()

	result, err := utils.RunCommand("")
	assert.Nil(t, err)
	assert.Equal(t, string(mockBytes), result, "Result from run was incorrect. Expected: %s, but got: %s.", result, string(mockBytes))
}
