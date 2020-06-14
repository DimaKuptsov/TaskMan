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
	tasksForDelete, err := action.TasksRepository.FindByIds(action.DTO.TasksIDs)
	if err != nil {
		return nil
	}
	deletedTasks := TasksCollection{}
	for _, task := range tasksForDelete.Tasks {
		err = task.MarkDeleted()
		if err != nil {
			//TODO log
			continue
		}
		deletedTasks.Add(task)
	}
	return action.TasksRepository.BatchUpdate(deletedTasks)
}
