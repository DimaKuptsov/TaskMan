package app

import (
	"github.com/DimaKuptsov/task-man/app/column"
	"github.com/DimaKuptsov/task-man/app/comment"
	"github.com/DimaKuptsov/task-man/app/project"
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/DimaKuptsov/task-man/mock"
	"github.com/go-playground/validator/v10"
)

type AppService struct {
	projectsService project.ProjectsService
	columnsService  column.ColumnsService
	tasksService    task.TasksService
	commentsService comment.CommentsService
}

func GetAppService() AppService {
	projectsRepository := mock.ProjectRepoMock()
	columnsRepository := mock.ColumnsRepoMock(projectsRepository)
	tasksRepository := mock.TasksRepoMock(columnsRepository)
	commentsRepository := mock.CommentsRepoMock(tasksRepository)
	validate := validator.New()
	commentsService := comment.CommentsService{
		Validate:           validate,
		CommentsRepository: commentsRepository,
	}
	tasksService := task.TasksService{
		Validate:        validate,
		TasksRepository: tasksRepository,
		CommentsService: commentsService,
	}
	columnsService := column.ColumnsService{
		Validate:          validate,
		ColumnsRepository: columnsRepository,
		TasksService:      tasksService,
	}
	projectsService := project.ProjectsService{
		Validate:       validate,
		Repository:     projectsRepository,
		ColumnsService: columnsService,
	}
	return AppService{
		projectsService: projectsService,
		columnsService:  columnsService,
		tasksService:    tasksService,
		commentsService: commentsService,
	}
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
