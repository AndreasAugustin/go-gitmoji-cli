package utils

import (
	"os/exec"
	"runtime"
)

var OsCommand = exec.Command
var CreateCommand = NewCommand
var runCommand = Run
var runCommandInDir = runInDir

type Command interface {
	CombinedOutput() ([]byte, error)
}

func GetRunnerInfo(operatingSystem string) (runner, runnerArg string) {
	if operatingSystem == "windows" {
		runner = "cmd.exe"
		runnerArg = "/C"
	} else {
		runner = "sh"
		runnerArg = "-c"
	}

	return runner, runnerArg
}

func NewCommand(directory, command string) Command {
	runner, runnerArg := GetRunnerInfo(runtime.GOOS)
	cmdArgs := []string{runnerArg, command}
	cmd := OsCommand(runner, cmdArgs...)

	if len(directory) > 0 {
		cmd.Dir = directory
	}

	return cmd
}

func Run(command string) (resultOutput string, err error) {
	return runInDir("", command)
}

func runInDir(directory, command string) (resultOutput string, err error) {
	cmd := CreateCommand(directory, command)

	out, err := cmd.CombinedOutput()
	resultOutput = string(out)

	return resultOutput, err
}
