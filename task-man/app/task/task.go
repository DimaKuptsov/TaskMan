package task

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
	ColumnIDField    = "column id"
	NameField        = "name"
	DescriptionField = "description"
)

type Task struct {
	ID          uuid.UUID   `pg:"type:uuid"`
	ColumnID    uuid.UUID   `pg:"type:uuid"`
	Name        Name        `pg:"type:varchar"`
	Description Description `pg:"type:varchar"`
	Priority    int         `pg:"type:int"`
	CreatedAt   time.Time   `pg:"type:timestamp"`
	UpdatedAt   pg.NullTime `pg:"type:timestamp"`
	DeletedAt   pg.NullTime `pg:"type:timestamp"`
}

func (t *Task) GetID() uuid.UUID {
	return t.ID
}

func (t *Task) GetColumnID() uuid.UUID {
	return t.ColumnID
}

func (t *Task) ChangeColumnID(columnID uuid.UUID) error {
	if columnID.String() == "" {
		return appErrors.ValidationError{Field: ColumnIDField, Message: "column id should be in the uuid format"}
	}
	t.ColumnID = columnID
	t.markUpdated()
	return nil
}

func (t *Task) GetName() Name {
	return t.Name
}

func (t *Task) ChangeName(name Name) error {
	err := validator.New().Struct(name)
	if err != nil {
		return appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	t.Name = name
	t.markUpdated()
	return nil
}

func (t *Task) GetDescription() Description {
	return t.Description
}

func (t *Task) ChangeDescription(description Description) error {
	err := validator.New().Struct(description)
	if err != nil {
		return appErrors.ValidationError{Field: DescriptionField, Message: err.Error()}
	}
	t.Description = description
	t.markUpdated()
	return nil
}

func (t *Task) GetPriority() int {
	return t.Priority
}

func (t *Task) ChangePriority(priority int) {
	t.Priority = priority
	t.markUpdated()
}

func (t *Task) IsDeleted() bool {
	return !t.DeletedAt.IsZero()
}

func (t *Task) MarkDeleted() {
	t.DeletedAt = pg.NullTime{Time: time.Now()}
}

func (t *Task) markUpdated() {
	t.UpdatedAt = pg.NullTime{Time: time.Now()}
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
		ID:          t.ID.String(),
		ColumnID:    t.ColumnID.String(),
		Name:        t.Name.String(),
		Description: t.Description.String(),
		Priority:    t.Priority,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
	}
	if !t.UpdatedAt.IsZero() {
		jsonTask.UpdatedAt = t.UpdatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonTask)
}
