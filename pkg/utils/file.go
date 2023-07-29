package utils

import (
	"os"
)

var WriteFile = Write
var readFile = Read
var IoWrite = os.WriteFile
var IoRead = os.ReadFile
var fileExists = Exists
var OsStat = os.Stat
var OsIsNotExist = os.IsNotExist

func Write(filePath string, contents []byte) error {
	return IoWrite(filePath, contents, os.ModePerm)
}

func Read(filePath string) ([]byte, error) {
	return IoRead(filePath)
}

func Exists(filePath string) bool {
	if _, err := OsStat(filePath); OsIsNotExist(err) {
		return false
	}

	return true
}
