package column

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

const DefaultColumnName = "Default column"

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
		return err
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
