package project

import (
	"ls-projects/models/config"

	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// setup
	os.Remove(getProjectsFilePath())

	code := m.Run()

	// teardown
	os.Remove(getProjectsFilePath())
	os.Remove(config.GetInstance().ConfigPath)

	os.Exit(code)
}

func Test_GetProjects(t *testing.T) {
	testRuns := []struct {
		testName        string
		initialDiskData string

		expectedData []Project
		expectErr    bool
	}{
		{
			testName:        "empty list",
			initialDiskData: "[]",

			expectedData: []Project{},
			expectErr:    false,
		},
		{
			testName: "single valid project",
			initialDiskData: `
			[
				{
					"name": "example-project",
					"path": "./"
				}
			]
			`,

			expectedData: []Project{
				{Name: "example-project", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "multiple valid projects",
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

			expectedData: []Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "single project with invalid path",
			initialDiskData: `
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
			testName: "single project with valid path including '~'",
			initialDiskData: `
			[
				{
					"name": "example-project",
					"path": "~"
				}
			]
			`,

			expectedData: []Project{
				{Name: "example-project", Path: "~"},
			},
			expectErr: false,
		},
		{
			testName: "invalid object",
			initialDiskData: `
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
			initialDiskData: `
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
			saveStringToFile(testRun.initialDiskData)

			p, err := GetAll()

			assert.Equal(t, testRun.expectedData, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_SaveProject(t *testing.T) {
	testRuns := []struct {
		testName        string
		initialDiskData string
		index           int
		project         Project

		expectedProjects []Project
		expectErr        bool
	}{
		{
			testName:        "save project into empty list",
			initialDiskData: "[]",
			project:         Project{Name: "example-project", Path: "./"},
			index:           0,

			expectedProjects: []Project{
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
			project: Project{Name: "example-project-2", Path: "./"},
			index:   0,

			expectedProjects: []Project{
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
			project: Project{Name: "example-project-3", Path: "./"},
			index:   1,

			expectedProjects: []Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
				{Name: "example-project-3", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "index out of bound with multiple elements list",
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
			project: Project{Name: "example-project-3", Path: "./"},
			index:   2,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName:        "negative index with empty list",
			initialDiskData: "[]",
			project:         Project{Name: "example-project-1", Path: "./"},
			index:           -1,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "negative index with multiple elements list",
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
			project: Project{Name: "example-project-1", Path: "./"},
			index:   -1,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName:        "save into invalid list",
			initialDiskData: "[{}]",
			project:         Project{Name: "example-project-2", Path: "./"},
			index:           0,

			expectedProjects: nil,
			expectErr:        true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.initialDiskData)

			p, err := Save(testRun.index, testRun.project)

			assert.Equal(t, testRun.expectedProjects, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_UpdateProject(t *testing.T) {
	testRuns := []struct {
		testName        string
		initialDiskData string
		index           int
		project         Project

		expectedProjects []Project
		expectErr        bool
	}{
		{
			testName: "update from single on-disk project",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				}
			]
			`,
			index:   0,
			project: Project{Name: "example-project-2", Path: "./"},

			expectedProjects: []Project{
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "update from multiple on-disk projects",
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
			index:   1,
			project: Project{Name: "example-project-3", Path: "./"},

			expectedProjects: []Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-3", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName:        "out of bound from empty list on-disk",
			initialDiskData: "[]",
			index:           0,
			project:         Project{Name: "example-project-1", Path: "./"},

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "out of bound from single on-disk project",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				}
			]
			`,
			index:   1,
			project: Project{Name: "example-project-1", Path: "./"},

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "out of bound from multiple on-disk projects",
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
			index:   3,
			project: Project{Name: "example-project-1", Path: "./"},

			expectedProjects: nil,
			expectErr:        true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.initialDiskData)

			p, err := Update(testRun.index, testRun.project)

			assert.Equal(t, testRun.expectedProjects, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_DeleteProject(t *testing.T) {
	testRuns := []struct {
		testName        string
		initialDiskData string
		index           int
		project         Project

		expectedProjects []Project
		expectErr        bool
	}{
		{
			testName: "delete valid project with single project on disk",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				}
			]
			`,
			index:   0,
			project: Project{Name: "example-project-1", Path: "./"},

			expectedProjects: []Project{},
			expectErr:        false,
		},
		{
			testName: "delete first project with mutliple project on disk",
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
			index:   0,
			project: Project{Name: "example-project-1", Path: "./"},

			expectedProjects: []Project{
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "delete last project with mutliple project on disk",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				},
				{
					"name": "example-project-2",
					"path": "./"
				},
				{
					"name": "example-project-3",
					"path": "./"
				}
			]
			`,
			index:   2,
			project: Project{Name: "example-project-3", Path: "./"},

			expectedProjects: []Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "delete out of bound index",
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
			index:   2,
			project: Project{Name: "example-project-2", Path: "./"},

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "delete invalid project",
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
			index:   0,
			project: Project{Name: "example-project-2", Path: "./"},

			expectedProjects: nil,
			expectErr:        true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.initialDiskData)

			p, err := Delete(testRun.index, testRun.project)

			assert.Equal(t, testRun.expectedProjects, p)
			if testRun.expectErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func Test_SwapIndex(t *testing.T) {
	testRuns := []struct {
		testName        string
		initialDiskData string
		initialIndex    int
		targetIndex     int

		expectedProjects []Project
		expectErr        bool
	}{
		{
			testName: "swap project from single element list",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				}
			]
			`,
			initialIndex: 0,
			targetIndex:  0,

			expectedProjects: []Project{
				{Name: "example-project-1", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "swap project from two elements list",
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
			initialIndex: 0,
			targetIndex:  1,

			expectedProjects: []Project{
				{Name: "example-project-2", Path: "./"},
				{Name: "example-project-1", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName: "swap project from mutliple (3+) elements list",
			initialDiskData: `
			[
				{
					"name": "example-project-1",
					"path": "./"
				},
				{
					"name": "example-project-2",
					"path": "./"
				},
				{
					"name": "example-project-3",
					"path": "./"
				}
			]
			`,
			initialIndex: 1,
			targetIndex:  2,

			expectedProjects: []Project{
				{Name: "example-project-1", Path: "./"},
				{Name: "example-project-3", Path: "./"},
				{Name: "example-project-2", Path: "./"},
			},
			expectErr: false,
		},
		{
			testName:        "swap project out of bound initial index - empty list",
			initialDiskData: "[]",
			initialIndex:    1,
			targetIndex:     0,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName:        "swap project out of bound target index - empty list",
			initialDiskData: "[]",
			initialIndex:    0,
			targetIndex:     1,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName:        "swap project out of bound initial index - negative",
			initialDiskData: "[]",
			initialIndex:    -1,
			targetIndex:     0,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName:        "swap project out of bound target index - negative",
			initialDiskData: "[]",
			initialIndex:    0,
			targetIndex:     -1,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "swap project out of bound initial index - not empty list",
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
			initialIndex: 2,
			targetIndex:  0,

			expectedProjects: nil,
			expectErr:        true,
		},
		{
			testName: "swap project out of bound target index - not empty list",
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
			initialIndex: 0,
			targetIndex:  2,

			expectedProjects: nil,
			expectErr:        true,
		},
	}

	for _, testRun := range testRuns {
		t.Run(testRun.testName, func(t *testing.T) {
			saveStringToFile(testRun.initialDiskData)

			p, err := SwapIndex(testRun.initialIndex, testRun.targetIndex)
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
	return os.WriteFile(getProjectsFilePath(), []byte(data), os.ModePerm)
}
