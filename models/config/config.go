// Package config implements functions required to create, load and edit a config.
package config

import "ls-projects/fileutil"

const (
	appDataPath          = "~/.config/ls-projects"
	projectsFileName     = ".projects.json"
	configFileName       = ".config.json"
	testProjectsFilePath = "./tests/default.test.projects.json"
	testConfigFilePath   = "./tests/default.test.config.json"
)

// Config represents the app's configuration on-disk as well as in memory.
type Config struct {
	// absolute path to the projects list file
	ProjectsPath string `json:"projectsPath"`

	// absolute path to the config file
	ConfigPath string `json:"configPath"`
}

// saveToDisk saves the config to the file
func (c Config) saveToDisk() error {
	return fileutil.SaveToFile(c, c.ConfigPath)
}

// config is the application's configuration
var config *Config

// GetInstace returns the app's configuration singleton.
func GetInstance() Config {
	if config == nil {
		config = initConfig()
	}
	return *config
}
