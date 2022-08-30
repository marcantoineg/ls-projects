// Package fileutil implements helper function to work with the file system.
package fileutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// SaveToFile encodes to JSON an object then saves it to a specified file.
func SaveToFile[T any](data T, filePath string) error {
	v, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, v, os.ModePerm)
	return err
}

// ReadFromFile tries to open the file at the given file path and returns all its content unmarshalled into the data parameter.
// If an error occurs, it is forwarded to the return value.
func ReadFromFile[T any](data *T, filePath string) error {
	file, err := os.Open(ReplaceTilde(filePath))
	if err != nil {
		return errors.New(fmt.Sprintf("error opening file '%s'\n\n%s", filePath, err))
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return errors.New(fmt.Sprintf("error reading file '%s'\n\n%s", filePath, err))
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return errors.New(fmt.Sprintf("error decoding objects from file '%s'\n\n%s", filePath, err))
	}

	return nil
}

// CreateEmptyFile creates the required directories and file. An empty string overwrites the content of the file.
// If an error occurs, it is forwarded to the return value.
func CreateEmptyFile(filePath string) error {
	return overwriteFileWithString(filePath, "")
}

// CreateEmptyListFile creates the required directories and file containing an empty list.
// If an error occurs, it is forwarded to the return value.
func CreateEmptyListFile(filePath string) error {
	return overwriteFileWithString(filePath, "[]")
}

// overwriteFileWithString creates or overwrites an existing file with the data provided.
// It also creates all required directory to the file if necessary.
// If an error occurs, it is forwarded to the return value.
func overwriteFileWithString(filePath string, data string) error {
	dataDir := filepath.Dir(filePath)
	err := os.MkdirAll(dataDir, os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

// ReplaceTilde returns a string with the tilde character replaced by the user's home directory.
func ReplaceTilde(filePath string) string {
	var newString = filePath
	if len(filePath) > 0 && filePath[0] == '~' {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newString = homeDir + filePath[1:]
	}
	return newString
}

// Exists returns true if the directory/file exists. Returns false if any error araise when fetching the directory/file info.
func Exists(path string) bool {
	path = ReplaceTilde(path)
	_, err := os.Stat(path)
	return err == nil
}
