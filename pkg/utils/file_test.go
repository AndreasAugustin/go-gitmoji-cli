package utils_test

import (
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"path"
	"testing"
)

const tmpTestDir = "./tmp/test_assets/"

func TestWriteFileUsesCorrectValues(t *testing.T) {
	var actualFileName string
	var actualData []byte
	var actualPerm os.FileMode
	expectedPerm := os.ModePerm
	expectedFileName := path.Join(tmpTestDir, "foo/bar.txt")
	data := []byte("fooBarRoo")
	expectedData := []byte(data)

	utils.OsWrite = func(filename string, data []byte, perm os.FileMode) error {
		actualFileName = filename
		actualData = data
		actualPerm = perm
		return nil
	}
	defer func() { utils.OsWrite = os.WriteFile }()
	actualErr := utils.WriteFile(expectedFileName, data)
	assert.Nil(t, actualErr)
	assert.Equal(t, expectedFileName, actualFileName)
	assert.Equal(t, expectedData, actualData, "Attempted to write incorrect data file. Expected: %s, but got: %s.", expectedData, actualData)
	assert.Equal(t, expectedPerm, actualPerm, "Attempted to use wrong perms on file. Expected: %d, but got: %d.", expectedPerm, actualPerm)
}

func TestReadFileUsesCorrectValues(t *testing.T) {
	var actualFilePath string
	expectedFilePath := path.Join(tmpTestDir, "foo/bar.txt")
	expectedData := []byte("fooBarRoo")

	utils.OsRead = func(filepath string) ([]byte, error) {
		actualFilePath = filepath
		return expectedData, nil
	}
	defer func() { utils.OsRead = os.ReadFile }()
	data, err := utils.ReadFile(expectedFilePath)
	assert.Nil(t, err)
	assert.Equal(t, expectedFilePath, actualFilePath, "Attempted to write to wrong file. Expected: %s, but got: %s.", expectedFilePath, actualFilePath)
	assert.Equal(t, expectedData, data, "Did not get correct file data. Expected: %s, but got: %s.", expectedData, data)
}

func TestFileExistsReturnsFalseWhenErrorIsOsNotExist(t *testing.T) {
	utils.OsStat = func(file string) (os.FileInfo, error) { return nil, nil }
	defer func() { utils.OsStat = os.Stat }()
	utils.OsIsNotExist = func(err error) bool { return true }
	defer func() { utils.OsIsNotExist = os.IsNotExist }()
	assert.False(t, utils.FileExists(path.Join(tmpTestDir, "/foo/repos/my-repo/go-gitmoji-cli.json")))
}

func TestFileExistsReturnsTrueWhenErrorIsNotOsNotExist(t *testing.T) {
	utils.OsStat = func(file string) (os.FileInfo, error) { return nil, nil }
	defer func() { utils.OsStat = os.Stat }()
	utils.OsIsNotExist = func(err error) bool { return false }
	defer func() { utils.OsIsNotExist = os.IsNotExist }()
	assert.True(t, utils.FileExists(path.Join(tmpTestDir, "foo/repos/my-repo/go-gitmoji-cli.json")))
}
