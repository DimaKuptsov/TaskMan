package project

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type ProjectsFactory struct {
	Validate *validator.Validate
}

func (f ProjectsFactory) Create(createDTO CreateDTO) (project Project, err error) {
	projectName := Name{createDTO.Name}
	err = f.Validate.Struct(projectName)
	if err != nil {
		return
	}
	description := Description{createDTO.Description}
	err = f.Validate.Struct(description)
	if err != nil {
		return
	}
	project = Project{
		id:          uuid.New(),
		name:        projectName,
		description: description,
		createdAt:   time.Now(),
	}
	return project, err
}
