package models

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	projectsFilePath     = "~/.config/list-my-projects"
	projectsFileName     = ".projects.json"
	testProjectsFilePath = "./test.projects.json"
)

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// ValidatePath returns a boolean value equal to wether or not the path exists on the host.
func (p Project) ValidatePath() bool {
	return exists(replaceTilde(p.Path))
}

// implements interface list.Item for type Project
func (p Project) FilterValue() string {
	return p.Name
}

// getProjectsFilePath returns an absolute path to the projects file directory.
func getProjectsFilePath() string {
	return replaceTilde(projectsFilePath)
}

// getFullProjectsFilePath returns an absolute path to the projects file.
func getFullProjectsFilePath() string {
	if flag.Lookup("test.v") == nil {
		return getProjectsFilePath() + "/" + projectsFileName
	} else {
		return testProjectsFilePath
	}
}

// GetProjects fetches the projects from the disk and returns them.
// If an error happens throughout the process, returns the error as the second return value.
func GetProjects() ([]Project, error) {
	if exists := exists(getFullProjectsFilePath()); !exists {
		err := createEmptyProjectsFile()
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Open(getFullProjectsFilePath())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	file.Close()

	var projects []Project
	err = json.Unmarshal(bytes, &projects)
	if err != nil {
		return nil, err
	}

	for i := range projects {
		var project = projects[i]
		if project.Name == "" || project.Path == "" {
			return nil, errors.New("both Name and Path fields are required.")
		}

		projects[i].Path = replaceTilde(projects[i].Path)
		exists := exists(projects[i].Path)
		if !exists {
			return nil, errors.New(fmt.Sprintf("directory/file %s does not exists", projects[i].Path))
		}
	}
	return projects, err
}

// SaveProject fetches the projects from the disk, appends the project given as the parameter then save the new projects on the disk.
// If no error is encountered, it returns the newly updated projects list. Else it returns the error as the second return value.
func SaveProject(project Project) ([]Project, error) {
	projects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	projects = append(projects, project)

	err = saveToFile(projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// createEmptyProjectsFile creates the required file to load and add new projects.
func createEmptyProjectsFile() error {
	err := os.MkdirAll(getProjectsFilePath(), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(getFullProjectsFilePath())
	if err != nil {
		return err
	}

	_, err = file.WriteString("[]")
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

func saveToFile(projects []Project) error {
	v, err := json.MarshalIndent(projects, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getFullProjectsFilePath(), v, os.ModePerm)
	return err
}

// replaceTilde returns a string with the tilde character replaced by the user's home directory
func replaceTilde(filePath string) string {
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

// exists returns true if the directory/file exists. Returns false if any error araise when fetching the directory/file info.
func exists(path string) bool {
	path = replaceTilde(path)
	_, err := os.Stat(path)
	return err == nil
}
