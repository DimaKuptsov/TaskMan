package task

import "github.com/DimaKuptsov/task-man/app/comment"

type DeleteColumnTasksAction struct {
	DTO             DeleteColumnTasksDTO
	TasksRepository TasksRepository
	CommentsService comment.CommentsService
}

func (action DeleteColumnTasksAction) Execute() error {
	tasksForDelete, err := action.TasksRepository.FindForColumn(action.DTO.ColumnID, WithoutDeletedTasks)
	if err != nil {
		return err
	}
	deleteTasksDTO := DeleteTasksDTO{}
	for _, task := range tasksForDelete.Tasks {
		deleteTasksDTO.AddTaskId(task.GetID())
	}
	deleteTasksAction := DeleteTasksAction{
		DTO:             deleteTasksDTO,
		CommentsService: action.CommentsService,
		TasksRepository: action.TasksRepository,
	}
	return deleteTasksAction.Execute()
}
