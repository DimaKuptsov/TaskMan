package app

import (
	"github.com/DimaKuptsov/task-man/app/column"
	"github.com/DimaKuptsov/task-man/app/comment"
	"github.com/DimaKuptsov/task-man/app/project"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/DimaKuptsov/task-man/db/postgres/repository"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/go-playground/validator/v10"
)

type AppService struct {
	projectsService project.ProjectsService
	columnsService  column.ColumnsService
	tasksService    task.TasksService
	commentsService comment.CommentsService
}

func GetAppService() (service AppService, err error) {
	validate := validator.New()
	appLogger := logger.GetLogger()
	commentsRepository, err := repository.NewCommentsRepository()
	if err != nil {
		return service, err
	}
	commentsService := comment.CommentsService{
		Validate:           validate,
		CommentsRepository: commentsRepository,
		Logger:             appLogger,
	}
	tasksRepository, err := repository.NewTasksRepository()
	if err != nil {
		return service, err
	}
	tasksService := task.TasksService{
		Validate:        validate,
		TasksRepository: tasksRepository,
		CommentsService: commentsService,
		Logger:          appLogger,
	}
	columnsRepository, err := repository.NewColumnsRepository()
	if err != nil {
		return service, err
	}
	columnsService := column.ColumnsService{
		Validate:          validate,
		ColumnsRepository: columnsRepository,
		TasksService:      tasksService,
		Logger:            appLogger,
	}
	projectsRepository, err := repository.NewProjectsRepository()
	if err != nil {
		return service, err
	}
	projectsService := project.ProjectsService{
		Validate:       validate,
		Repository:     projectsRepository,
		ColumnsService: columnsService,
		Logger:         appLogger,
	}
	service = AppService{
		projectsService: projectsService,
		columnsService:  columnsService,
		tasksService:    tasksService,
		commentsService: commentsService,
	}
	return service, err
}

func (as AppService) ProjectsService() project.ProjectsService {
	return as.projectsService
}

func (as AppService) ColumnsService() column.ColumnsService {
	return as.columnsService
}

func (as AppService) TasksService() task.TasksService {
	return as.tasksService
}

func (as AppService) CommentsService() comment.CommentsService {
	return as.commentsService
}
