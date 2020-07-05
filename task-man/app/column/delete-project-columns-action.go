package column

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/task"
	"go.uber.org/zap"
)

type DeleteProjectColumnsAction struct {
	DTO               DeleteProjectColumnsDTO
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
	Logger            *zap.Logger
}

func (action DeleteProjectColumnsAction) Execute() error {
	columnsForDelete, err := action.ColumnsRepository.FindForProject(action.DTO.ProjectID, WithoutDeletedColumns)
	if err != nil {
		return err
	}
	return action.deleteColumns(columnsForDelete)
}

func (action DeleteProjectColumnsAction) deleteColumns(columnsForDelete ColumnsCollection) error {
	defer action.Logger.Sync()
	deletedColumns := ColumnsCollection{}
	for _, column := range columnsForDelete.Columns {
		column.MarkDeleted()
		deletedColumns.Add(column)
		err := action.deleteColumnTasks(column)
		if err != nil {
			errMsg := fmt.Sprintf("column tasks delete was fallen: %s", err.Error())
			action.Logger.Warn(errMsg, zap.String("id", column.GetID().String()))
			continue
		}
	}
	return action.ColumnsRepository.BatchUpdate(deletedColumns)
}

func (action DeleteProjectColumnsAction) deleteColumnTasks(deletedColumn Column) error {
	deleteColumnTasksDTO := task.DeleteColumnTasksDTO{ColumnID: deletedColumn.GetID()}
	return action.TasksService.DeleteColumnTasks(deleteColumnTasksDTO)
}
