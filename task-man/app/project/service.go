package project

import (
	"github.com/DimaKuptsov/task-man/app/column"
	"github.com/go-playground/validator/v10"
)

type ProjectsService struct {
	Validate       *validator.Validate
	Repository     ProjectsRepository
	ColumnsService column.ColumnsService
}

func (ps ProjectsService) CreateProject(createDTO CreateDTO) (project Project, err error) {
	projectsFactory := ProjectsFactory{Validate: ps.Validate}
	createAction := CreateProjectAction{
		DTO:           createDTO,
		Repository:    ps.Repository,
		Factory:       projectsFactory,
		ColumnService: ps.ColumnsService,
	}
	return createAction.Execute()
}

func (ps ProjectsService) UpdateProject(updateDTO UpdateDTO) (project Project, err error) {
	updateAction := UpdateProjectAction{
		DTO:        updateDTO,
		Repository: ps.Repository,
	}
	return updateAction.Execute()
}

func (ps ProjectsService) DeleteProject(deleteDTO DeleteDTO) error {
	deleteAction := DeleteProjectAction{
		DTO:            deleteDTO,
		Repository:     ps.Repository,
		ColumnsService: ps.ColumnsService,
	}
	return deleteAction.Execute()
}
