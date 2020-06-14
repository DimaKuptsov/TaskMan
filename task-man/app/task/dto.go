package task

import "github.com/google/uuid"

type CreateTaskDTO struct {
	ColumnID    uuid.UUID
	Name        string
	Description string
}

type UpdateDTO struct {
	ID          uuid.UUID
	Name        string
	Description string
	Priority    int
}

type ChangeTasksColumnDTO struct {
	TasksIDs []uuid.UUID
	ColumnID uuid.UUID
}

func (dto *ChangeTasksColumnDTO) AddTaskId(taskID uuid.UUID) {
	dto.TasksIDs = append(dto.TasksIDs, taskID)
}

type DeleteTasksDTO struct {
	TasksIDs []uuid.UUID
}

func (dto *DeleteTasksDTO) AddTaskId(taskID uuid.UUID) {
	dto.TasksIDs = append(dto.TasksIDs, taskID)
}

type DeleteColumnTasksDTO struct {
	ColumnID uuid.UUID
}
