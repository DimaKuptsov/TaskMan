package task

import "go.uber.org/zap"

type ChangeTasksColumnAction struct {
	DTO        ChangeTasksColumnDTO
	Repository TasksRepository
	Logger     *zap.Logger
}

func (action ChangeTasksColumnAction) Execute() error {
	defer action.Logger.Sync()
	tasks, err := action.Repository.FindByIds(action.DTO.TasksIDs, WithoutDeletedTasks)
	if err != nil {
		return err
	}
	tasksForUpdate := TasksCollection{}
	for _, task := range tasks.Tasks {
		if task.GetColumnID() == action.DTO.ColumnID {
			continue
		}
		err = task.ChangeColumnID(action.DTO.ColumnID)
		if err != nil {
			action.Logger.Warn(
				"not changed task for column %s",
				zap.String("new_task_id", action.DTO.ColumnID.String()),
			)
			continue
		}
		tasksForUpdate.Add(task)
	}
	return action.Repository.BatchUpdate(tasksForUpdate)
}
