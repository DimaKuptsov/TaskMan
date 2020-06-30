package project

import (
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-pg/pg/v10"
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
	ID          uuid.UUID   `pg:"type:uuid"`
	Name        Name        `pg:"type:varchar"`
	Description Description `pg:"type:varchar"`
	CreatedAt   time.Time   `pg:"type:timestamp"`
	UpdatedAt   pg.NullTime `pg:"type:timestamp"`
	DeletedAt   pg.NullTime `pg:"type:timestamp"`
}

func (p *Project) GetID() uuid.UUID {
	return p.ID
}

func (p *Project) GetName() Name {
	return p.Name
}

func (p *Project) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	p.Name = name
	p.markUpdated()
	return nil
}

func (p *Project) GetDescription() Description {
	return p.Description
}

func (p *Project) ChangeDescription(description Description) error {
	err := validator.New().Struct(description)
	if err != nil {
		return appErrors.ValidationError{Field: DescriptionField, Message: err.Error()}
	}
	p.Description = description
	p.markUpdated()
	return nil
}

func (p *Project) GetCreateTime() time.Time {
	return p.CreatedAt
}

func (p *Project) IsDeleted() bool {
	return !p.DeletedAt.IsZero()
}

func (p *Project) MarkDeleted() {
	p.DeletedAt = pg.NullTime{Time: time.Now()}
}

func (p *Project) markUpdated() {
	p.UpdatedAt = pg.NullTime{Time: time.Now()}
}

func (p Project) MarshalJSON() ([]byte, error) {
	jsonProject := struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}{
		ID:          p.ID.String(),
		Name:        p.Name.String(),
		Description: p.Description.String(),
		CreatedAt:   p.CreatedAt.Format(time.RFC3339),
	}
	if !p.UpdatedAt.IsZero() {
		jsonProject.UpdatedAt = p.UpdatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonProject)
}
