package column

import "github.com/google/uuid"

type CreateDTO struct {
	ProjectID uuid.UUID
	Name      string
}

type UpdateDTO struct {
	ID       uuid.UUID
	Name     string
	Priority int
}

type DeleteDTO struct {
	ID uuid.UUID
}

type DeleteProjectColumnsDTO struct {
	ProjectID uuid.UUID
}
