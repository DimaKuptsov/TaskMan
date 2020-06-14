package column

import (
	"errors"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/google/uuid"
)

type DeleteColumnAction struct {
	DTO               DeleteDTO
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
}

func (action DeleteColumnAction) Execute() error {
	columnForDelete, err := action.ColumnsRepository.FindById(action.DTO.ID)
	if err != nil {
		return err
	}
	notDeletedProjectsColumns, err := action.getNotDeletedProjectsColumns(columnForDelete)
	if err != nil {
		return err
	}
	const minProjectColumnsCount = 1
	if notDeletedProjectsColumns.Len() == minProjectColumnsCount {
		return errors.New("can't delete last column")
	}
	err = action.deleteColumn(columnForDelete)
	if err != nil {
		return err
	}
	previousColumnID := action.getPreviousColumnID(columnForDelete, notDeletedProjectsColumns)
	return action.manageDeleteColumnTasks(columnForDelete, previousColumnID)
}

func (action DeleteColumnAction) getNotDeletedProjectsColumns(columnForDelete Column) (columns ColumnsCollection, err error) {
	notDeletedProjectsColumns, err := action.ColumnsRepository.FindForProject(columnForDelete.GetProjectID(), WithoutDeletedColumns)
	if err != nil {
		return
	}
	notDeletedProjectsColumns.SortByPriority()
	return notDeletedProjectsColumns, err
}

func (action DeleteColumnAction) deleteColumn(columnForDelete Column) error {
	err := columnForDelete.MarkDeleted()
	if err != nil {
		return err
	}
	return action.ColumnsRepository.Update(columnForDelete)
}

func (action DeleteColumnAction) getPreviousColumnID(columnForCompare Column, columns ColumnsCollection) uuid.UUID {
	var previousColumnID uuid.UUID
	for i, projectColumn := range columns.Columns {
		columnHasHighestPriority := i == 0
		currentColumnWillBeDeleted := projectColumn.GetID() == columnForCompare.GetID()
		if columnHasHighestPriority && currentColumnWillBeDeleted {
			break
		}
		if !columnHasHighestPriority && currentColumnWillBeDeleted {
			previousColumn := columns.Columns[i-1]
			previousColumnID = previousColumn.GetID()
		}
	}
	return previousColumnID
}

func (action DeleteColumnAction) manageDeleteColumnTasks(columnForDelete Column, previousColumnID uuid.UUID) error {
	columnTasks, err := action.TasksService.TasksRepository.FindForColumn(columnForDelete.GetID(), task.WithDeletedTasks)
	if err != nil {
		return err
	}
	changeTasksColumnDTO := task.ChangeTasksColumnDTO{ColumnID: previousColumnID}
	deleteTasksDTO := task.DeleteTasksDTO{}
	for _, columnTask := range columnTasks.Tasks {
		if previousColumnID.String() != "" {
			changeTasksColumnDTO.AddTaskId(columnTask.GetID())
		} else {
			deleteTasksDTO.AddTaskId(columnTask.GetID())
		}
	}
	err = action.TasksService.ChangeTasksColumn(changeTasksColumnDTO)
	if err != nil {
		return err
	}
	return action.TasksService.DeleteTasks(deleteTasksDTO)
}
