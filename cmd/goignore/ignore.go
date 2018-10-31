package goignore

import (
	"io/ioutil"
	"os"
)

// IsExistGitignore is exist
func IsExistGitignore() bool {
	_, err := os.Stat("./.gitignore")
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true

}

// CreateGitignore is create .gitignore
func CreateGitignore() error {
	_, err := os.Create(".gitignore")
	return err
}

// ReadTheExistGitignore is read the exist .gitignore
func ReadTheExistGitignore() ([]byte, error) {
	return ioutil.ReadFile("./.gitignore")
}

// OpenGitignore is open the exist .gitignore
func OpenGitignore() (*os.File, error) {
	return os.OpenFile("./.gitignore", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
}

func TruncateGitignore() error {
	return os.Truncate("./.gitignore", 0)
}
