package column

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/google/uuid"
)

type DeleteColumnAction struct {
	DTO               DeleteDTO
	ColumnsRepository ColumnsRepository
	TasksService      task.TasksService
}

func (action DeleteColumnAction) Execute() error {
	columnForDelete, err := action.ColumnsRepository.FindById(action.DTO.ID, WithoutDeletedColumns)
	if err != nil {
		return err
	}
	notDeletedProjectsColumns, err := action.getNotDeletedProjectsColumns(columnForDelete)
	if err != nil {
		return err
	}
	const minProjectColumnsCount = 1
	if notDeletedProjectsColumns.Len() == minProjectColumnsCount {
		return appErrors.ValidationError{Field: IDField, Message: "can't delete last column"}
	}
	err = action.deleteColumn(columnForDelete)
	if err != nil {
		return err
	}
	columnIDForTaskTransfer := action.getColumnIDForTaskTransfer(columnForDelete, notDeletedProjectsColumns)
	return action.manageDeleteColumnTasks(columnForDelete, columnIDForTaskTransfer)
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
	columnForDelete.MarkDeleted()
	return action.ColumnsRepository.Update(columnForDelete)
}

func (action DeleteColumnAction) getColumnIDForTaskTransfer(columnForCompare Column, columns ColumnsCollection) uuid.UUID {
	var columnForTransferID uuid.UUID
	columns.SortByPriority()
	for i, projectColumn := range columns.Columns {
		columnHasHighestPriority := i == 0
		currentColumnWillBeDeleted := projectColumn.GetID() == columnForCompare.GetID()
		lessPriorityColumnIndex := i + 1
		if columnHasHighestPriority && currentColumnWillBeDeleted {
			if lessPriorityColumnIndex < columns.Len() {
				columnForTransfer := columns.Columns[lessPriorityColumnIndex]
				columnForTransferID = columnForTransfer.GetID()
			}
			break
		}
		highestPriorityColumnIndex := i - 1
		if !columnHasHighestPriority && currentColumnWillBeDeleted {
			columnForTransfer := columns.Columns[highestPriorityColumnIndex]
			columnForTransferID = columnForTransfer.GetID()
		}
	}
	return columnForTransferID
}

func (action DeleteColumnAction) manageDeleteColumnTasks(columnForDelete Column, columnIDForTaskTransfer uuid.UUID) error {
	columnTasks, err := action.TasksService.TasksRepository.FindForColumn(columnForDelete.GetID(), task.WithDeletedTasks)
	if err != nil {
		return err
	}
	changeTasksColumnDTO := task.ChangeTasksColumnDTO{ColumnID: columnIDForTaskTransfer}
	deleteTasksDTO := task.DeleteTasksDTO{}
	for _, columnTask := range columnTasks.Tasks {
		if columnIDForTaskTransfer.String() != "" {
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
