package column

import (
	"github.com/DimaKuptsov/task-man/app/task"
)

type DeleteProjectColumnsAction struct {
	DTO               DeleteProjectColumnsDTO
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
}

func (action DeleteProjectColumnsAction) Execute() error {
	columnsForDelete, err := action.ColumnsRepository.FindForProject(action.DTO.ProjectID, WithoutDeletedColumns)
	if err != nil {
		return err
	}
	return action.deleteColumns(columnsForDelete)
}

func (action DeleteProjectColumnsAction) deleteColumns(columnsForDelete ColumnsCollection) error {
	deletedColumns := ColumnsCollection{}
	for _, column := range columnsForDelete.Columns {
		err := column.MarkDeleted()
		if err != nil {
			//TODO log
			continue
		}
		deletedColumns.Add(column)
		err = action.deleteColumnTasks(column)
		if err != nil {
			//TODO log
			continue
		}
	}
	return action.ColumnsRepository.BatchUpdate(deletedColumns)
}

func (action DeleteProjectColumnsAction) deleteColumnTasks(deletedColumn Column) error {
	deleteColumnTasksDTO := task.DeleteColumnTasksDTO{ColumnID: deletedColumn.GetID()}
	return action.TasksService.DeleteColumnTasks(deleteColumnTasksDTO)
}
