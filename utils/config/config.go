package config

import (
	"flag"
	"fmt"
	"list-my-projects/utils"
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
	return utils.SaveToFile(c, c.ConfigPath)
}

const (
	appDataPath          = "~/.config/list-my-projects"
	projectsFileName     = ".projects.json"
	configFileName       = ".config.json"
	testProjectsFilePath = "./test.projects.json"
	testConfigFilePath   = "./test.config.json"
)

// config is the application's configuration
var config *Config

// GetInstace returns the app's configuration singleton.
func GetInstance() Config {
	if config == nil {
		config = InitConfig()
	}
	return *config
}

// initConfig creates and returns the initial configuration.
func InitConfig() *Config {
	projectsPath := flag.String("projects", defaultFullProjectsFilePath(), "path to the projects file")
	configPath := flag.String("config", defaultFullConfigPath(), "path to the config file")

	flag.Parse()

	if exists := utils.Exists(*configPath); exists {
		c := readOnDiskConfig(*configPath)
		if c.ConfigPath == "" || c.ProjectsPath == "" {
			fmt.Println("config file must specify 'configPath' and 'projectPath'", *configPath)
		}
		return c
	} else {
		fmt.Printf("file '%s' does not exists.", *configPath)
		utils.CreateEmptyFile(*configPath)

		newConfig := &Config{
			ProjectsPath: *projectsPath,
			ConfigPath:   *configPath,
		}

		newConfig.saveToDisk()

		return newConfig
	}
}

// readOnDiskConfig returns a pointer to a parsed Config from the disk.
// Returns nil if and error occurs.
func readOnDiskConfig(configPath string) *Config {
	var config Config
	err := utils.ReadFromFile(&config, configPath)
	if err != nil {
		fmt.Printf("error loading config file '%s'\n\n%s", configPath, err.Error())
		panic(err)
	} else {
		return &config
	}
}

// defaultFullProjectsFilePath returns the ddefault absolute path to the projects file.
func defaultFullProjectsFilePath() string {
	if flag.Lookup("test.v") == nil {
		return appDataPath + "/" + projectsFileName
	} else {
		return testProjectsFilePath
	}
}

// defaultFullConfigPath returns the default absolute path to the config file.
func defaultFullConfigPath() string {
	if flag.Lookup("test.v") == nil {
		return appDataPath + "/" + configFileName
	} else {
		return testConfigFilePath
	}
}
