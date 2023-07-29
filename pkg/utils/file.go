package utils

import (
	"os"
)

var WriteFile = write
var ReadFile = read
var OsWrite = os.WriteFile
var OsRead = os.ReadFile
var FileExists = exists
var OsStat = os.Stat
var OsIsNotExist = os.IsNotExist
var OsRemove = os.Remove
var RemoveFile = remove

func write(filePath string, contents []byte) error {
	return OsWrite(filePath, contents, os.ModePerm)
}

func remove(filePath string) error {
	return OsRemove(filePath)
}

func read(filePath string) ([]byte, error) {
	return OsRead(filePath)
}

func exists(filePath string) bool {
	if _, err := OsStat(filePath); OsIsNotExist(err) {
		return false
	}

	return true
}
