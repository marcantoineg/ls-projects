package project

import (
	"encoding/json"
	"errors"
	"fmt"
	fileutils "list-my-projects/utils"
)

// A Project stores simple information about a project on disk.
// It is a representation of the data on-disk and in-memory.
type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// implements interface list.Item for type Project
func (p Project) FilterValue() string {
	return p.Name
}

// ValidatePath returns a boolean value equal to wether or not the path exists on the host.
func (p Project) ValidatePath() bool {
	return fileutils.Exists(p.Path)
}

// GetProjects fetches the projects from the disk and returns them.
// If an error happens throughout the process, returns the error as the second return value.
func GetProjects() ([]Project, error) {
	if exists := fileutils.Exists(fileutils.GetFullProjectsFilePath()); !exists {
		err := fileutils.CreateEmptyProjectsFile()
		if err != nil {
			return nil, err
		}
	}

	bytes, err := fileutils.ReadFromFile(fileutils.GetFullProjectsFilePath())
	if err != nil {
		return nil, err
	}

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

		exists := fileutils.Exists(projects[i].Path)
		if !exists {
			return nil, errors.New(fmt.Sprintf("directory/file %s does not exists", projects[i].Path))
		}
	}
	return projects, err
}

// SaveProject fetches the projects from the disk, appends the project given as the parameter at the given index, then saves the new projects on the disk.
// If no error is encountered, it returns the newly updated projects list. Else it returns the error as the second return value.
func SaveProject(index int, project Project) ([]Project, error) {
	onDiskProjects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	if index < 0 || (index >= len(onDiskProjects) && len(onDiskProjects) != 0) {
		return nil, errors.New("index out of bound")
	}

	var projects []Project
	if len(onDiskProjects) == 0 {
		projects = []Project{project}
	} else {
		projects = append([]Project{}, onDiskProjects[:index+1]...)
		projects = append(projects, project)
		projects = append(projects, onDiskProjects[index+1:]...)
	}

	err = fileutils.SaveToFile(projects, fileutils.GetFullProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// EditProject edit the project list on-disk.
// If the index is not found, an error is returned as the second parameter
func UpdateProject(index int, project Project) ([]Project, error) {
	projects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	if index < 0 || index >= len(projects) {
		return nil, errors.New("index out of bound")
	}

	projects[index] = project

	err = fileutils.SaveToFile(projects, fileutils.GetFullProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// DeleteProject fetches the projects from the disk by index, checks it's the same as the in-memory project, then deletes it from the disk.
// If no error is encountered, it returns the newly updated projects list. Else it returns the error as the second return value.
func DeleteProject(index int, project Project) ([]Project, error) {
	projects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	if index < 0 || index >= len(projects) {
		return nil, errors.New("project not found")
	}

	onDiskProject := projects[index]
	if onDiskProject.Name != project.Name || onDiskProject.Path != project.Path {
		return nil, errors.New("project on disk did not match project in memory")
	}

	projects = append(projects[:index], projects[index+1:]...)

	err = fileutils.SaveToFile(projects, fileutils.GetFullProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}
