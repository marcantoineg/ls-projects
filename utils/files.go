// Package files utils implements helper function to work with the file system
package utils

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	appDataPath          = "~/.config/list-my-projects"
	projectsFileName     = ".projects.json"
	configFileName       = ".config.json"
	testProjectsFilePath = "./test.projects.json"
	testConfigFilePath   = "./test.config.json"
)

// GetFullProjectsFilePath returns the absolute path to the projects file.
func GetFullProjectsFilePath() string {
	if flag.Lookup("test.v") == nil {
		return dataAbsolutePath() + "/" + projectsFileName
	} else {
		return testProjectsFilePath
	}
}

// GetFullConfigPath returns the absolute path to the config file.
func GetFullConfigPath() string {
	if flag.Lookup("test.v") == nil {
		return dataAbsolutePath() + "/" + configFileName
	} else {
		return testConfigFilePath
	}
}

// dataAbsoluatePath returns the absolute path to app's data directory.
func dataAbsolutePath() string {
	return ReplaceTilde(appDataPath)
}

// SaveToFile encodes to JSON a list of items then saves it to a specified file.
func SaveToFile[T any](items []T, filePath string) error {
	v, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filePath, v, os.ModePerm)
	return err
}

// ReadFromFile tries to open the file at the given file path and returns all its content.
func ReadFromFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// CreateEmptyProjectsFile creates the required file to load and add new projects.
func CreateEmptyProjectsFile() error {
	err := os.MkdirAll(dataAbsolutePath(), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(GetFullProjectsFilePath())
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("[]")
	if err != nil {
		return err
	}

	return nil
}

// ReplaceTilde returns a string with the tilde character replaced by the user's home directory
func ReplaceTilde(filePath string) string {
	var newString = filePath
	if filePath[0] == '~' {
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
