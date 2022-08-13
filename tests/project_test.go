package tests

import (
	"io/ioutil"
	"list-my-projects/models"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testProjectsFilePath = "./test.projects.json"
)

func TestMain(m *testing.M) {
	// setup
	os.Remove(testProjectsFilePath)

	code := m.Run()

	// teardown
	os.Remove(testProjectsFilePath)

	os.Exit(code)
}
func TestGetProjects(t *testing.T) {
	testRuns := []struct {
		testName     string
		onDiskData   string
		expectedData []models.Project
		expectErr    bool
	}{
		{
			testName:     "empty list",
			onDiskData:   "[]",
			expectedData: []models.Project{},
			expectErr:    false,
		},
		{
			testName: "single valid project",
			onDiskData: `
			[
				{
					"name": "example-project",
					"path": "./"
				}
			]
			`,
			expectedData: []models.Project{
				{Name: "example-project", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "multiple valid projects",
			onDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				},
				{
					"name": "example-project-2",
					"path": "./"
				}
			]
			`,
			expectedData: []models.Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "single project with invalid path",
			onDiskData: `
			[
				{
					"name": "example-project",
					"path": "not-a-valid-path"
				}
			]
			`,
			expectedData: nil,
			expectErr:    true,
		},
		{
			testName: "invalid object",
			onDiskData: `
			[
				{
					"not-name": true,
					"not-path": 1
				}
			]
			`,
			expectedData: nil,
			expectErr:    true,
		},
		{
			testName: "empty object",
			onDiskData: `
			[
				{
				}
			]
			`,
			expectedData: nil,
			expectErr:    true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.onDiskData)

			p, err := models.GetProjects()

			assert.Equal(t, testRun.expectedData, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestSaveProject(t *testing.T) {
	testRuns := []struct {
		testName         string
		initialDiskData  string
		project          models.Project
		expectedProjects []models.Project
		expectErr        bool
	}{
		{
			testName:        "save project into empty list",
			initialDiskData: "[]",
			project:         models.Project{Name: "example-project", Path: "./"},
			expectedProjects: []models.Project{
				{Name: "example-project", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "save project into single element list",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				}
			]
			`,
			project: models.Project{Name: "example-project-2", Path: "./"},
			expectedProjects: []models.Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "save project into multi elements list",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				},
				{
					"name": "example-project-2",
					"path": "./"
				}
			]
			`,
			project: models.Project{Name: "example-project-3", Path: "./"},
			expectedProjects: []models.Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
				{Name: "example-project-3", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName:         "save into invalid list",
			initialDiskData:  "[{}]",
			project:          models.Project{Name: "example-project-2", Path: "./"},
			expectedProjects: nil,
			expectErr:        true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.initialDiskData)

			p, err := models.SaveProject(testRun.project)

			assert.Equal(t, testRun.expectedProjects, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func saveStringToFile(data string) error {
	return ioutil.WriteFile(testProjectsFilePath, []byte(data), os.ModePerm)
}
