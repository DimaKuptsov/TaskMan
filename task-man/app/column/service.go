package column

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ColumnsService struct {
	Validate          *validator.Validate
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
}

func (cs ColumnsService) GetById(columnID uuid.UUID) (column Column, err error) {
	if columnID.String() == "" {
		return column, appErrors.ValidationError{Field: IDField, Message: "column id should be in the uuid format"}
	}
	return cs.ColumnsRepository.FindById(columnID)
}

func (cs ColumnsService) GetForProject(projectID uuid.UUID) (columns ColumnsCollection, err error) {
	if projectID.String() == "" {
		return columns, appErrors.ValidationError{Field: ProjectIDField, Message: "project id should be in the uuid format"}
	}
	columns, err = cs.ColumnsRepository.FindForProject(projectID, WithoutDeletedColumns)
	if err != nil {
		return columns, err
	}
	columns.SortByPriority()
	return columns, err
}

func (cs ColumnsService) CreateColumn(createDTO CreateDTO) (column Column, err error) {
	columnsFactory := ColumnsFactory{Validate: cs.Validate, ColumnsRepository: cs.ColumnsRepository}
	createAction := CreateColumnAction{
		DTO:        createDTO,
		Repository: cs.ColumnsRepository,
		Factory:    columnsFactory,
	}
	return createAction.Execute()
}

func (cs ColumnsService) UpdateColumn(updateDTO UpdateDTO) (column Column, err error) {
	updateAction := UpdateColumnAction{
		DTO:        updateDTO,
		Repository: cs.ColumnsRepository,
	}
	return updateAction.Execute()
}

func (cs ColumnsService) DeleteColumn(deleteDTO DeleteDTO) error {
	deleteAction := DeleteColumnAction{
		DTO:               deleteDTO,
		ColumnsRepository: cs.ColumnsRepository,
		TasksService:      cs.TasksService,
	}
	return deleteAction.Execute()
}

func (cs ColumnsService) DeleteProjectColumns(deleteDTO DeleteProjectColumnsDTO) error {
	deleteAction := DeleteProjectColumnsAction{
		DTO:               deleteDTO,
		ColumnsRepository: cs.ColumnsRepository,
		TasksService:      cs.TasksService,
	}
	return deleteAction.Execute()
}
