package comment

import (
	"database/sql"
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const (
	IDField     = "id"
	TaskIDField = "task id"
	TextField   = "text"
)

type Comment struct {
	id        uuid.UUID
	taskID    uuid.UUID
	text      Text
	createdAt time.Time
	updatedAt sql.NullTime
	deletedAt sql.NullTime
}

func (c *Comment) GetID() uuid.UUID {
	return c.id
}

func (c *Comment) GetTaskID() uuid.UUID {
	return c.taskID
}

func (c *Comment) GetText() Text {
	return c.text
}

func (c *Comment) ChangeText(text Text) error {
	err := validator.New().Struct(text)
	if err != nil {
		return appErrors.ValidationError{Field: TextField, Message: err.Error()}
	}
	c.text = text
	return c.markUpdated()
}

func (c *Comment) GetCreateTime() time.Time {
	return c.createdAt
}

func (c *Comment) IsDeleted() bool {
	return c.deletedAt.Valid
}

func (c *Comment) MarkDeleted() error {
	return c.deletedAt.Scan(time.Now())
}

func (c *Comment) markUpdated() error {
	return c.updatedAt.Scan(time.Now())
}

func (c Comment) MarshalJSON() ([]byte, error) {
	jsonColumn := struct {
		ID        string `json:"id"`
		TaskID    string `json:"task_id"`
		Text      string `json:"text"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}{
		ID:        c.id.String(),
		TaskID:    c.taskID.String(),
		Text:      c.text.String(),
		CreatedAt: c.createdAt.Format(time.RFC3339),
	}
	if c.updatedAt.Valid {
		jsonColumn.UpdatedAt = c.updatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonColumn)
}
