package task

type ChangeTasksColumnAction struct {
	DTO        ChangeTasksColumnDTO
	Repository TasksRepository
}

func (action ChangeTasksColumnAction) Execute() error {
	tasks, err := action.Repository.FindByIds(action.DTO.TasksIDs)
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
			//TODO log
			continue
		}
		tasksForUpdate.Add(task)
	}
	return action.Repository.BatchUpdate(tasksForUpdate)
}
