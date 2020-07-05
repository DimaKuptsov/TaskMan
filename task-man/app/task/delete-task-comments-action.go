package task

import "github.com/DimaKuptsov/task-man/app/comment"

type DeleteTasksAction struct {
	DTO             DeleteTasksDTO
	TasksRepository TasksRepository
	CommentsService comment.CommentsService
}

func (action DeleteTasksAction) Execute() error {
	err := action.deleteTasks()
	if err != nil {
		return err
	}
	deleteTasksCommentsDTO := comment.DeleteTasksCommentsDTO{TasksIDs: action.DTO.TasksIDs}
	return action.CommentsService.DeleteTaskComments(deleteTasksCommentsDTO)
}

func (action DeleteTasksAction) deleteTasks() error {
	tasksForDelete, err := action.TasksRepository.FindByIds(action.DTO.TasksIDs, WithoutDeletedTasks)
	if err != nil {
		return err
	}
	deletedTasks := TasksCollection{}
	for _, task := range tasksForDelete.Tasks {
		task.MarkDeleted()
		deletedTasks.Add(task)
	}
	return action.TasksRepository.BatchUpdate(deletedTasks)
}
