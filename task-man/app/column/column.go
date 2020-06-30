package column

import (
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-pg/pg/v10"
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
	ID        uuid.UUID   `pg:"type:uuid"`
	ProjectID uuid.UUID   `pg:"type:uuid"`
	Name      Name        `pg:"type:varchar"`
	Priority  int         `pg:"type:int"`
	CreatedAt time.Time   `pg:"type:timestamp"`
	UpdatedAt pg.NullTime `pg:"type:timestamp"`
	DeletedAt pg.NullTime `pg:"type:timestamp"`
}

func (c *Column) GetID() uuid.UUID {
	return c.ID
}

func (c *Column) GetProjectID() uuid.UUID {
	return c.ProjectID
}

func (c *Column) GetName() Name {
	return c.Name
}

func (c *Column) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	c.Name = name
	c.markUpdated()
	return nil
}

func (c *Column) GetPriority() int {
	return c.Priority
}

func (c *Column) ChangePriority(priority int) {
	c.Priority = priority
	c.markUpdated()
}

func (c *Column) IsDeleted() bool {
	return !c.DeletedAt.IsZero()
}

func (c *Column) MarkDeleted() {
	c.DeletedAt = pg.NullTime{Time: time.Now()}
}

func (c *Column) markUpdated() {
	c.UpdatedAt = pg.NullTime{Time: time.Now()}
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
		ID:        c.ID.String(),
		ProjectID: c.ProjectID.String(),
		Name:      c.Name.String(),
		Priority:  c.Priority,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	}
	if !c.UpdatedAt.IsZero() {
		jsonColumn.UpdatedAt = c.UpdatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonColumn)
}
