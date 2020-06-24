package column

import (
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/go-playground/validator/v10"
)

type ColumnsService struct {
	Validate          *validator.Validate
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
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
