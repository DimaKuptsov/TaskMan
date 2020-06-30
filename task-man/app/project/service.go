package project

import (
	"github.com/DimaKuptsov/task-man/app/column"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ProjectsService struct {
	Validate       *validator.Validate
	Repository     ProjectsRepository
	ColumnsService column.ColumnsService
	Logger         *zap.Logger
}

func (ps ProjectsService) GetAll() (projects ProjectsCollection, err error) {
	projects, err = ps.Repository.FindNotDeleted()
	if err != nil {
		return projects, err
	}
	projects.SortByCreateTime()
	return projects, err
}

func (ps ProjectsService) GetById(uuid uuid.UUID) (project Project, err error) {
	if uuid.String() == "" {
		return project, appErrors.ValidationError{Field: IDField, Message: "project id should be in the uuid format"}
	}
	return ps.Repository.FindById(uuid, WithoutDeletedProjects)
}

func (ps ProjectsService) CreateProject(createDTO CreateDTO) (project Project, err error) {
	projectsFactory := ProjectsFactory{Validate: ps.Validate}
	createAction := CreateProjectAction{
		DTO:           createDTO,
		Repository:    ps.Repository,
		Factory:       projectsFactory,
		ColumnService: ps.ColumnsService,
		Logger:        ps.Logger,
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
