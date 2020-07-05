package task

import (
	"github.com/DimaKuptsov/task-man/app/comment"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type TasksService struct {
	Validate        *validator.Validate
	TasksRepository TasksRepository
	CommentsService comment.CommentsService
	Logger          *zap.Logger
}

func (ts TasksService) GetById(taskID uuid.UUID) (task Task, err error) {
	if taskID.String() == "" {
		return task, appErrors.ValidationError{Field: IDField, Message: "task id should be in the uuid format"}
	}
	return ts.TasksRepository.FindById(taskID, WithoutDeletedTasks)
}

func (ts TasksService) GetForColumn(columnID uuid.UUID) (tasks TasksCollection, err error) {
	if columnID.String() == "" {
		return tasks, appErrors.ValidationError{Field: ColumnIDField, Message: "column id should be in the uuid format"}
	}
	tasks, err = ts.TasksRepository.FindForColumn(columnID, WithoutDeletedTasks)
	if err != nil {
		return tasks, err
	}
	tasks.SortByPriority()
	return tasks, err
}

func (ts TasksService) CreateTask(createDTO CreateTaskDTO) (task Task, err error) {
	tasksFactory := TasksFactory{Validate: ts.Validate, TasksRepository: ts.TasksRepository}
	createAction := CreateTaskAction{
		DTO:             createDTO,
		TasksRepository: ts.TasksRepository,
		TasksFactory:    tasksFactory,
	}
	return createAction.Execute()
}

func (ts TasksService) UpdateTask(updateDTO UpdateDTO) (task Task, err error) {
	updateAction := UpdateTaskAction{
		DTO:        updateDTO,
		Repository: ts.TasksRepository,
	}
	return updateAction.Execute()
}

func (ts TasksService) ChangeTasksColumn(changeTasksColumnDTO ChangeTasksColumnDTO) error {
	changeTasksColumnAction := ChangeTasksColumnAction{
		DTO:        changeTasksColumnDTO,
		Repository: ts.TasksRepository,
		Logger:     ts.Logger,
	}
	return changeTasksColumnAction.Execute()
}

func (ts TasksService) DeleteColumnTasks(deleteDTO DeleteColumnTasksDTO) error {
	deleteTasksAction := DeleteColumnTasksAction{
		DTO:             deleteDTO,
		TasksRepository: ts.TasksRepository,
		CommentsService: ts.CommentsService,
	}
	return deleteTasksAction.Execute()
}

func (ts TasksService) DeleteTasks(deleteDTO DeleteTasksDTO) error {
	deleteTasksAction := DeleteTasksAction{
		DTO:             deleteDTO,
		TasksRepository: ts.TasksRepository,
		CommentsService: ts.CommentsService,
	}
	return deleteTasksAction.Execute()
}
