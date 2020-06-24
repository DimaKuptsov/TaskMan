package task

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type TasksFactory struct {
	Validate        *validator.Validate
	TasksRepository TasksRepository
}

func (f TasksFactory) Create(createDTO CreateTaskDTO) (task Task, err error) {
	columnId := createDTO.ColumnID
	if columnId.String() == "" {
		err = errors.New("invalid column id")
	}
	taskName := Name{createDTO.Name}
	err = f.Validate.Struct(taskName)
	if err != nil {
		return
	}
	taskDescription := Description{createDTO.Description}
	err = f.Validate.Struct(taskDescription)
	if err != nil {
		return
	}
	columnTasks, err := f.TasksRepository.FindForColumn(createDTO.ColumnID, WithoutDeletedTasks)
	if err != nil {
		return
	}
	task = Task{
		id:          uuid.New(),
		columnID:    columnId,
		name:        taskName,
		description: taskDescription,
		priority:    columnTasks.Len(),
		createdAt:   time.Now(),
	}
	return task, err
}
