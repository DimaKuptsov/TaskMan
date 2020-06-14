package project

import "github.com/google/uuid"

type ProjectsRepository interface {
	FindNotDeleted() ([]Project, error)
	FindById(id uuid.UUID) (Project, error)
	Save(project Project) error
	Update(project Project) error
}
