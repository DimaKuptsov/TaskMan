package task

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
	ColumnIDField    = "column id"
	NameField        = "name"
	DescriptionField = "description"
)

type Task struct {
	id          uuid.UUID
	columnID    uuid.UUID
	name        Name
	description Description
	priority    int
	createdAt   time.Time
	updatedAt   sql.NullTime
	deletedAt   sql.NullTime
}

func (t *Task) GetID() uuid.UUID {
	return t.id
}

func (t *Task) GetColumnID() uuid.UUID {
	return t.columnID
}

func (t *Task) ChangeColumnID(columnID uuid.UUID) error {
	t.columnID = columnID
	return t.markUpdated()
}

func (t *Task) GetName() Name {
	return t.name
}

func (t *Task) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	t.name = name
	return t.markUpdated()
}

func (t *Task) GetDescription() Description {
	return t.description
}

func (t *Task) ChangeDescription(description Description) error {
	err := validator.New().Struct(description)
	if err != nil {
		return appErrors.ValidationError{Field: DescriptionField, Message: err.Error()}
	}
	t.description = description
	return t.markUpdated()
}

func (t *Task) GetPriority() int {
	return t.priority
}

func (t *Task) ChangePriority(priority int) error {
	t.priority = priority
	return t.markUpdated()
}

func (t *Task) IsDeleted() bool {
	return t.deletedAt.Valid
}

func (t *Task) MarkDeleted() error {
	return t.deletedAt.Scan(time.Now())
}

func (t *Task) markUpdated() error {
	return t.updatedAt.Scan(time.Now())
}

func (t Task) MarshalJSON() ([]byte, error) {
	jsonTask := struct {
		ID          string `json:"id"`
		ColumnID    string `json:"column_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Priority    int    `json:"priority"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}{
		ID:          t.id.String(),
		ColumnID:    t.columnID.String(),
		Name:        t.name.String(),
		Description: t.description.String(),
		Priority:    t.priority,
		CreatedAt:   t.createdAt.Format(time.RFC3339),
	}
	if t.updatedAt.Valid {
		jsonTask.UpdatedAt = t.updatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonTask)
}
