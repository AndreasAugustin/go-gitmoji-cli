package utils

import (
	"os"
	"path"
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
var CreateDirIfNotExists = createDirIfNotExists
var GetCacheDir = getCacheDir
var GetCacheDirCreateIfNotExists = getCacheDirCreateIfNotExists
var GetGlobalConfigDir = getGlobalConfigDir
var GetUserConfigDirCreateIfNotExists = getUserConfigDirCreateIfNotExists

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

func createDirIfNotExists(dir string) error {
	if !exists(dir) {
		return os.MkdirAll(dir, os.ModeDir|0755)
	}
	return nil
}

func getUserConfigDirCreateIfNotExists(programName string) (string, error) {
	cacheDir, err := GetGlobalConfigDir(programName)

	if err != nil {
		return "", err
	}

	if err := createDirIfNotExists(cacheDir); err != nil {
		return "", err
	}

	return cacheDir, nil
}

func getGlobalConfigDir(programName string) (string, error) {
	userConfigDir, err := os.UserConfigDir()

	if err != nil {
		return "", err
	}
	return path.Join(userConfigDir, programName), nil
}

func getCacheDir(programName string) (string, error) {
	userCacheDir, err := os.UserCacheDir()

	if err != nil {
		return "", err
	}
	return path.Join(userCacheDir, programName), nil
}

func getCacheDirCreateIfNotExists(programName string) (string, error) {
	cacheDir, err := GetCacheDir(programName)

	if err != nil {
		return "", err
	}

	if err := createDirIfNotExists(cacheDir); err != nil {
		return "", err
	}

	return cacheDir, nil
}
