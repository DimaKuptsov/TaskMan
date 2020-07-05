package task

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
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
		return task, appErrors.ValidationError{Field: IDField, Message: "task id should be in the uuid format"}
	}
	taskName := Name{createDTO.Name}
	err = f.Validate.Struct(taskName)
	if err != nil {
		return task, appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	taskDescription := Description{createDTO.Description}
	err = f.Validate.Struct(taskDescription)
	if err != nil {
		return task, appErrors.ValidationError{Field: DescriptionField, Message: err.Error()}
	}
	columnTasks, err := f.TasksRepository.FindForColumn(createDTO.ColumnID, WithoutDeletedTasks)
	if err != nil {
		return
	}
	priority := columnTasks.Len() + 1
	task = Task{
		ID:          uuid.New(),
		ColumnID:    columnId,
		Name:        taskName,
		Description: taskDescription,
		Priority:    priority,
		CreatedAt:   time.Now(),
	}
	return task, err
}
