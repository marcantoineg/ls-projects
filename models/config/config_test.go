package config

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// setup
	os.Remove(testConfigFilePath)

	code := m.Run()

	// teardown
	os.Remove(testConfigFilePath)

	os.Exit(code)
}

func TestGetInstance(t *testing.T) {
	testRuns := []struct {
		testName         string
		initialConfigPtr *Config

		expectedConfig Config
	}{
		{
			testName:         "GetInstance with nil internal config",
			initialConfigPtr: nil,

			expectedConfig: Config{
				ProjectsPath:            defaultFullProjectsFilePath(),
				ConfigPath:              defaultFullConfigPath(),
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},
		},
		{
			testName: "GetInstance with initial internal config",
			initialConfigPtr: &Config{
				ProjectsPath:            "./",
				ConfigPath:              "./",
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},

			expectedConfig: Config{
				ProjectsPath:            "./",
				ConfigPath:              "./",
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			config = testRun.initialConfigPtr

			actual := GetInstance()

			assert.Equal(t, testRun.expectedConfig, actual)
		})
	}
}

func Test_initConfig(t *testing.T) {
	testRuns := []struct {
		testName              string
		configFlag            string
		projectsFlag          string
		initialConfigFileData string

		expectedConfig Config
		expectPanic    bool
	}{
		{
			testName:              "no flag without config file expects default config",
			initialConfigFileData: "",

			expectedConfig: Config{
				ConfigPath:              defaultFullConfigPath(),
				ProjectsPath:            defaultFullProjectsFilePath(),
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},
			expectPanic: false,
		},
		{
			testName:              "no flag with invalid config file expects panic",
			initialConfigFileData: "{}",

			expectedConfig: Config{},
			expectPanic:    true,
		},
		{
			testName: "no flag with valid custom config file expects custom config",
			initialConfigFileData: `
			{
				"configPath": "some-custom-path",
				"projectsPath": "some-custom-path-2",
				"projectSelectionCommand": "cd"
			}
			`,

			expectedConfig: Config{
				ConfigPath:              "some-custom-path",
				ProjectsPath:            "some-custom-path-2",
				ProjectSelectionCommand: "cd",
				ProjectSelectionArgs:    []string(nil),
			},
			expectPanic: false,
		},
		{
			testName: "valid config path flag with valid custom config file expects custom config",
			initialConfigFileData: `
			{
				"configPath": "some-custom-path",
				"projectsPath": "some-custom-path-2",
				"projectSelectionCommand": "cd"
			}
			`,
			configFlag: "./tests/config.json",

			expectedConfig: Config{
				ConfigPath:              "some-custom-path",
				ProjectsPath:            "some-custom-path-2",
				ProjectSelectionCommand: "cd",
				ProjectSelectionArgs:    []string(nil),
			},
			expectPanic: false,
		},
		{
			testName:              "valid projects path flag without file returns custom config",
			initialConfigFileData: "",
			projectsFlag:          "./tests/projects.json",

			expectedConfig: Config{
				ConfigPath:              defaultFullConfigPath(),
				ProjectsPath:            "./tests/projects.json",
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},
			expectPanic: false,
		},
		{
			testName:              "valid projects path and projects flags without file returns custom config",
			initialConfigFileData: "",
			configFlag:            "./tests/config.json",
			projectsFlag:          "./tests/projects.json",

			expectedConfig: Config{
				ConfigPath:              "./tests/config.json",
				ProjectsPath:            "./tests/projects.json",
				ProjectSelectionCommand: defaultProjectSelectionCommand(),
				ProjectSelectionArgs:    defaultProjectSelectionArgsSlice(),
			},
			expectPanic: false,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			var configPath = testRun.configFlag

			flag.Set("config", testRun.configFlag)
			flag.Set("projects", testRun.projectsFlag)

			if configPath == "" {
				configPath = testConfigFilePath
			}

			if testRun.initialConfigFileData != "" {
				saveStringToFile(configPath, testRun.initialConfigFileData)
			}

			if testRun.expectPanic {
				assert.Panics(t, func() { initConfig() })
			} else {
				assert.Equal(t, testRun.expectedConfig, *initConfig())
			}

			assert.FileExists(t, configPath)

			os.Remove(configPath)
		})
	}
}

func saveStringToFile(filePath, data string) error {
	return ioutil.WriteFile(filePath, []byte(data), os.ModePerm)
}
