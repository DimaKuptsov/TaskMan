package project

import (
	"database/sql"
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const (
	IDField          = "id"
	NameField        = "name"
	DescriptionField = "description"
)

type Project struct {
	id          uuid.UUID
	name        Name
	description Description
	createdAt   time.Time
	updatedAt   sql.NullTime
	deletedAt   sql.NullTime
}

func (p *Project) GetID() uuid.UUID {
	return p.id
}

func (p *Project) GetName() Name {
	return p.name
}

func (p *Project) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	p.name = name
	return p.markUpdated()
}

func (p *Project) GetDescription() Description {
	return p.description
}

func (p *Project) ChangeDescription(description Description) error {
	err := validator.New().Struct(description)
	if err != nil {
		return appErrors.ValidationError{Field: DescriptionField, Message: err.Error()}
	}
	p.description = description
	return p.markUpdated()
}

func (p *Project) GetCreateTime() time.Time {
	return p.createdAt
}

func (p *Project) IsDeleted() bool {
	return p.deletedAt.Valid
}

func (p *Project) MarkDeleted() error {
	return p.deletedAt.Scan(time.Now())
}

func (p *Project) markUpdated() error {
	return p.updatedAt.Scan(time.Now())
}

func (p Project) MarshalJSON() ([]byte, error) {
	jsonProject := struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}{
		ID:          p.id.String(),
		Name:        p.name.String(),
		Description: p.description.String(),
		CreatedAt:   p.createdAt.Format(time.RFC3339),
	}
	if p.updatedAt.Valid {
		jsonProject.UpdatedAt = p.updatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonProject)
}
