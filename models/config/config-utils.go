package config

import (
	"flag"
	"fmt"

	"github.com/marcantoineg/fileutil"
)

func init() {
	flag.String("config", "", "path to the config file")
	flag.String("projects", "", "path to the projects file")
}

// initConfig parses command line flags and loads the configuration accordingly. If no config file is found, creates it.
// Returns the config or panics.
func initConfig() *Config {
	flag.Parse()

	var configPath = defaultFullConfigPath()
	if f := flag.Lookup("config"); f != nil {
		if value := f.Value.String(); value != "" {
			configPath = f.Value.String()
		}
	}
	var projectsPath = defaultFullProjectsFilePath()
	if f := flag.Lookup("projects"); f != nil {
		if value := f.Value.String(); value != "" {
			projectsPath = f.Value.String()
		}
	}

	if exists := fileutil.Exists(configPath); exists {
		return readOnDiskConfig(configPath)
	} else {
		fmt.Printf("file '%s' does not exists. creating...\n", configPath)
		err := fileutil.CreateEmptyFile(configPath)
		if err != nil {
			panic(err)
		}

		newConfig := &Config{
			ProjectsPath: projectsPath,
			ConfigPath:   configPath,
		}

		newConfig.saveToDisk()

		return newConfig
	}
}

// readOnDiskConfig returns a pointer to a parsed Config from the disk.
// Returns nil if and error occurs.
func readOnDiskConfig(configPath string) *Config {
	var config Config
	err := fileutil.ReadFromFile(&config, configPath)

	if err != nil {
		fmt.Printf("error loading config file '%s'\n\n%s", configPath, err.Error())
		panic(err)
	} else if config.ConfigPath == "" || config.ProjectsPath == "" {
		panic(fmt.Errorf("config file '%s' must specify 'configPath' and 'projectsPath'", configPath))
	}

	return &config
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
