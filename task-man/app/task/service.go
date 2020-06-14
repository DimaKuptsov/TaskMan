package task

import (
	"github.com/DimaKuptsov/task-man/app/comment"
	"github.com/go-playground/validator/v10"
)

type TasksService struct {
	Validate        *validator.Validate
	TasksRepository TasksRepository
	CommentsService comment.CommentsService
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
