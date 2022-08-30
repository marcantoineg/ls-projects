// todo
package config

import "list-my-projects/fileutil"

const (
	appDataPath          = "~/.config/list-my-projects"
	projectsFileName     = ".projects.json"
	configFileName       = ".config.json"
	testProjectsFilePath = "../../tests/test.projects.json"
	testConfigFilePath   = "../../tests/test.config.json"
)

// Config represents the app's configuration.
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
