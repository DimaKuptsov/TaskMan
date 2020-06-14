package project

import "github.com/google/uuid"

type CreateDTO struct {
	Name        string
	Description string
}

type UpdateDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
}

type DeleteDTO struct {
	ID uuid.UUID
}
