package comment

import (
	"encoding/json"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-pg/pg/v10"
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
	ID        uuid.UUID   `pg:"type:uuid"`
	TaskID    uuid.UUID   `pg:"type:uuid"`
	Text      Text        `pg:"type:varchar"`
	CreatedAt time.Time   `pg:"type:timestamp"`
	UpdatedAt pg.NullTime `pg:"type:timestamp"`
	DeletedAt pg.NullTime `pg:"type:timestamp"`
}

func (c *Comment) GetID() uuid.UUID {
	return c.ID
}

func (c *Comment) GetTaskID() uuid.UUID {
	return c.TaskID
}

func (c *Comment) GetText() Text {
	return c.Text
}

func (c *Comment) ChangeText(text Text) error {
	err := validator.New().Struct(text)
	if err != nil {
		return appErrors.ValidationError{Field: TextField, Message: err.Error()}
	}
	c.Text = text
	c.markUpdated()
	return nil
}

func (c *Comment) GetCreateTime() time.Time {
	return c.CreatedAt
}

func (c *Comment) IsDeleted() bool {
	return !c.DeletedAt.IsZero()
}

func (c *Comment) MarkDeleted() {
	c.DeletedAt = pg.NullTime{Time: time.Now()}
}

func (c *Comment) markUpdated() {
	c.UpdatedAt = pg.NullTime{Time: time.Now()}
}

func (c Comment) MarshalJSON() ([]byte, error) {
	jsonColumn := struct {
		ID        string `json:"id"`
		TaskID    string `json:"task_id"`
		Text      string `json:"text"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at,omitempty"`
	}{
		ID:        c.ID.String(),
		TaskID:    c.TaskID.String(),
		Text:      c.Text.String(),
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	}
	if !c.UpdatedAt.IsZero() {
		jsonColumn.UpdatedAt = c.UpdatedAt.Time.Format(time.RFC3339)
	}
	return json.Marshal(&jsonColumn)
}
