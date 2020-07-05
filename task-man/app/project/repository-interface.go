package project

import "github.com/google/uuid"

const (
	WithDeletedProjects    = true
	WithoutDeletedProjects = false
)

type ProjectsRepository interface {
	FindNotDeleted() (ProjectsCollection, error)
	FindById(id uuid.UUID, withDeleted bool) (Project, error)
	Save(project Project) error
	Update(project Project) error
}
