package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

const projectsFilePath = "~/.config/go-apps/.projects.json"

// GetProjects fetch the projects from the disk and returns them.
// If an error happens throughout the process, returns the error as the second return value.
func GetProjects() ([]Project, error) {
	file, err := os.Open(replaceTilde(projectsFilePath))
	if err != nil {
		return []Project{}, err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return []Project{}, err
	}
	file.Close()

	var projects []Project
	err = json.Unmarshal(bytes, &projects)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range projects {
		projects[i].Path = replaceTilde(projects[i].Path)
		exists := exists(projects[i].Path)
		if !exists {
			fmt.Println(fmt.Sprintf("directory/file %s does not exists", projects[i].Path))
			os.Exit(1)
		}
	}
	return projects, err
}

//  replaceTilde returns a string with the tilde character replaced by the user's home directory
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
	_, err := os.Stat(path)
	return err == nil
}
