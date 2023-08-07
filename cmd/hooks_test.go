package cmd_test

import (
	"bytes"
	"github.com/AndreasAugustin/go-gitmoji-cli/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHooksCommand(t *testing.T) {
	actual := new(bytes.Buffer)
	cmd.RootCmd.SetOut(actual)
	cmd.RootCmd.SetErr(actual)
	cmd.RootCmd.SetArgs([]string{"hooks", "--help"})

	assert.NoError(t, cmd.RootCmd.Execute())
}
