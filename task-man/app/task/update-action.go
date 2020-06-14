package task

type UpdateTaskAction struct {
	DTO        UpdateDTO
	Repository TasksRepository
}

func (action UpdateTaskAction) Execute() (updatedTask Task, err error) {
	updatedTask, err = action.Repository.FindById(action.DTO.ID)
	if err != nil {
		return
	}
	if action.DTO.Name != updatedTask.GetName().String() {
		newName := Name{action.DTO.Name}
		err = updatedTask.ChangeName(newName)
		if err != nil {
			return
		}
	}
	if action.DTO.Description != updatedTask.GetDescription().String() {
		newDescription := Description{action.DTO.Description}
		err = updatedTask.ChangeDescription(newDescription)
		if err != nil {
			return
		}
	}
	tasksForUpdate := TasksCollection{}
	if action.DTO.Priority != updatedTask.GetPriority() {
		columnTasks, err := action.Repository.FindForColumn(updatedTask.GetColumnID(), WithoutDeletedTasks)
		if err != nil {
			return updatedTask, err
		}
		for _, task := range columnTasks.Tasks {
			if task.GetPriority() == action.DTO.Priority {
				err = updatedTask.ChangePriority(task.GetPriority())
				if err != nil {
					return updatedTask, err
				}
				err = task.ChangePriority(action.DTO.Priority)
				if err != nil {
					return updatedTask, err
				}
				tasksForUpdate.Add(task)
				break
			}
		}
	}
	tasksForUpdate.Add(updatedTask)
	err = action.Repository.BatchUpdate(tasksForUpdate)
	if err != nil {
		return
	}
	return updatedTask, err
}
