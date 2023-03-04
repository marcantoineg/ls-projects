package projectform

import "ls-projects/models/project"

type NoProjectCreatedMsg struct{}
type ProjectCreatedMsg struct {
	Project project.Project
}
type ProjectCreationErrorMsg error
type ProjectUpdatedMsg struct {
	Project project.Project
}
type ProjectUpdateErrorMsg error
