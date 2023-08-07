package cmd_test

import (
	"bytes"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/cmd"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootVersionCommand(t *testing.T) {
	var resVersion = cmd.RootCmd.Version

	assert.Equal(t, pkg.Version, resVersion)
}

func TestRootHelpCommandAll(t *testing.T) {
	var casesArgs = []string{"--help", "help"}
	var testFkt = func(arg string) func(t *testing.T) {
		return func(t *testing.T) {
			actual := new(bytes.Buffer)
			cmd.RootCmd.SetOut(actual)
			cmd.RootCmd.SetErr(actual)
			cmd.RootCmd.SetArgs([]string{arg})
			assert.NoError(t, cmd.RootCmd.Execute())

			assert.Containsf(t, actual.String(), fmt.Sprintf("See %s for more information about Gitmoji", pkg.DefaultGitmojiUrl), "error %s", "formatted")
		}
	}
	for _, arg := range casesArgs {
		t.Run(arg, testFkt(arg))
	}
}
