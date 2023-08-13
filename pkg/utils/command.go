package utils

import (
	"os/exec"
	"runtime"
)

var GetRunnerInfo = getRunnerInfo
var OsCommand = exec.Command
var CreateCommand = newCommand
var RunCommand = run
var RunCommandInDir = runInDir

type Command interface {
	CombinedOutput() ([]byte, error)
}

func getRunnerInfo(operatingSystem string) (runner, runnerArg string) {
	if operatingSystem == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}

	return runner, runnerArg
}

func newCommand(directory, command string) Command {
	runner, runnerArg := getRunnerInfo(runtime.GOOS)
	cmdArgs := []string{runnerArg, command}
	cmd := OsCommand(runner, cmdArgs...)

	if len(directory) > 0 {
		cmd.Dir = directory
	}

	return cmd
}

func run(command string) (resultOutput string, err error) {
	return runInDir("", command)
}

func runInDir(directory, command string) (resultOutput string, err error) {
	cmd := CreateCommand(directory, command)

	out, err := cmd.CombinedOutput()
	resultOutput = string(out)

	return resultOutput, err
}
