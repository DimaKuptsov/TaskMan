package project

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
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
		return err
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
		return err
	}
	p.description = description
	return p.markUpdated()
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
