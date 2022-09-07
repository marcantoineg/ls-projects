package config

import (
	"flag"
	"fmt"
	"list-my-projects/fileutil"
	"strings"
)

func init() {
	flag.String("config", "", "path to the config file")
	flag.String("projects", "", "path to the projects file")
	flag.String("command", "", "command to exectute when a project gets selected")
	flag.String("commandArgs", "", "Arguments for the command to exectute when a project gets selected")
}

// initConfig parses command line flags and loads the configuration accordingly. If no config file is found, creates it.
// Returns the config or panics.
func initConfig() *Config {
	flag.Parse()

	configPath := getFlagValue(defaultFullConfigPath(), "config")
	projectsPath := getFlagValue(defaultFullProjectsFilePath(), "projects")
	commandName := getFlagValue(defaultProjectSelectionCommand(), "command")
	commandAgs := getFlagValue(defaultProjectSelectionArgs(), "commandArgs")

	if exists := fileutil.Exists(configPath); exists {
		return readOnDiskConfig(configPath)
	} else {
		fmt.Printf("file '%s' does not exists. creating...\n", configPath)
		err := fileutil.CreateEmptyFile(configPath)
		if err != nil {
			panic(err)
		}

		newConfig := &Config{
			ProjectsPath:            projectsPath,
			ConfigPath:              configPath,
			ProjectSelectionCommand: commandName,
			ProjectSelectionArgs:    strings.Split(commandAgs, " "),
		}

		newConfig.saveToDisk()

		return newConfig
	}
}

func getFlagValue(defaultValue string, flagName string) string {
	if f := flag.Lookup(flagName); f != nil {
		if value := f.Value.String(); value != "" {
			return f.Value.String()
		}
	}
	return defaultValue
}

// readOnDiskConfig returns a pointer to a parsed Config from the disk.
// Returns nil if and error occurs.
func readOnDiskConfig(configPath string) *Config {
	var config Config
	err := fileutil.ReadFromFile(&config, configPath)

	if err != nil {
		fmt.Printf("error loading config file '%s'\n\n%s", configPath, err.Error())
		panic(err)
	} else if config.ConfigPath == "" || config.ProjectsPath == "" || config.ProjectSelectionCommand == "" {
		fmt.Println(fmt.Sprintf("config file '%s' must specify 'configPath', 'projectsPath' and 'projectSelectionCommand'", configPath))
		panic(err)
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

func defaultProjectSelectionCommand() string {
	return "code"
}

func defaultProjectSelectionArgs() string {
	return "-n ."
}

func defaultProjectSelectionArgsSlice() []string {
	return strings.Split(defaultProjectSelectionArgs(), " ")
}
