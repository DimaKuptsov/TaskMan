package comment

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
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
		return err
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
