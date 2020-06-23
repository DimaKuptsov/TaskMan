package column

import (
	"database/sql"
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const (
	DefaultColumnName = "Default column"
	IDField           = "id"
	NameField         = "name"
	ProjectIDField    = "project id"
)

type Column struct {
	id        uuid.UUID
	projectID uuid.UUID
	name      Name
	priority  int
	createdAt time.Time
	updatedAt sql.NullTime
	deletedAt sql.NullTime
}

func (c *Column) GetID() uuid.UUID {
	return c.id
}

func (c *Column) GetProjectID() uuid.UUID {
	return c.projectID
}

func (c *Column) GetName() Name {
	return c.name
}

func (c *Column) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	c.name = name
	return c.markUpdated()
}

func (c *Column) GetPriority() int {
	return c.priority
}

func (c *Column) ChangePriority(priority int) error {
	c.priority = priority
	return c.markUpdated()
}

func (c *Column) IsDeleted() bool {
	return c.deletedAt.Valid
}

func (c *Column) MarkDeleted() error {
	return c.deletedAt.Scan(time.Now())
}

func (c *Column) markUpdated() error {
	return c.updatedAt.Scan(time.Now())
}

func (c Column) MarshalJSON() ([]byte, error) {
	jsonColumn := struct {
		ID        string `json:"id"`
		ProjectID string `json:"project_id"`
		Name      string `json:"name"`
		Priority  int    `json:"priority"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}{
		ID:        c.id.String(),
		ProjectID: c.projectID.String(),
		Name:      c.name.String(),
		Priority:  c.priority,
		CreatedAt: c.createdAt.Format(time.RFC3339),
	}
	if c.updatedAt.Valid {
		jsonColumn.UpdatedAt = c.updatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonColumn)
}
