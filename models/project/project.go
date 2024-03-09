// Package project implements functions required to create, load, edit and delete projects.
package project

import (
	"errors"
	"fmt"

	"ls-projects/models/config"

	"github.com/marcantoineg/fileutil"
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
	return fileutil.Exists(p.Path)
}

// GetAll fetches the projects from the disk and returns them.
// If an error happens throughout the process, it returns the error as the second return value.
func GetAll() ([]Project, error) {
	if exists := fileutil.Exists(getProjectsFilePath()); !exists {
		err := fileutil.CreateEmptyListFile(getProjectsFilePath())
		if err != nil {
			return nil, err
		}
	}

	var projects []Project
	err := fileutil.ReadFromFile(&projects, getProjectsFilePath())
	if err != nil {
		return nil, err
	}

	for i := range projects {
		var project = projects[i]
		if project.Name == "" || project.Path == "" {
			return nil, errors.New("both Name and Path fields are required")
		}

		exists := fileutil.Exists(projects[i].Path)
		if !exists {
			return nil, fmt.Errorf("directory/file %s does not exists", projects[i].Path)
		}
	}
	return projects, err
}

// Save fetches the projects from the disk, appends the project given as the parameter at the given index, then saves the new projects on the disk.
// If no error is encountered, it returns the newly updated projects list. Else it returns the error as the second return value.
func Save(index int, project Project) ([]Project, error) {
	onDiskProjects, err := GetAll()
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

	err = fileutil.SaveToFile(projects, getProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// Update edit the project list on-disk.
// If the index is not found, an error is returned as the second parameter
func Update(index int, project Project) ([]Project, error) {
	projects, err := GetAll()
	if err != nil {
		return nil, err
	}

	if index < 0 || index >= len(projects) {
		return nil, errors.New("index out of bound")
	}

	projects[index] = project

	err = fileutil.SaveToFile(projects, getProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// Delete fetches the projects from the disk by index, checks it's the same as the in-memory project, then deletes it from the disk.
// If no error is encountered, it returns the newly updated projects list. Else it returns the error as the second return value.
func Delete(index int, project Project) ([]Project, error) {
	projects, err := GetAll()
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

	err = fileutil.SaveToFile(projects, getProjectsFilePath())
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// SwapIndex fetches the projects from the disk, swap both projects by index then saves the updated list.
// Returns the updated list if no error occurs. Forwards the error otherwise.
func SwapIndex(initialIndex int, targetIndex int) ([]Project, error) {
	projects, err := GetAll()
	if err != nil {
		return nil, err
	}

	if initialIndex < 0 || initialIndex >= len(projects) {
		return nil, errors.New("initial index out of bound")
	} else if targetIndex < 0 || targetIndex >= len(projects) {
		return nil, errors.New("target index out of bound")
	}

	if initialIndex == targetIndex {
		return projects, nil
	}

	p := projects[initialIndex]
	projects[initialIndex] = projects[targetIndex]
	projects[targetIndex] = p

	err = fileutil.SaveToFile(projects, getProjectsFilePath())

	return projects, err
}

// getProjectsFilePath fetches the projects file path from the app's config.
func getProjectsFilePath() string {
	return config.GetInstance().ProjectsPath
}
