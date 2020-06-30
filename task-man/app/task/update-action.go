package task

import appErrors "github.com/DimaKuptsov/task-man/app/error"

type UpdateTaskAction struct {
	DTO        UpdateDTO
	Repository TasksRepository
}

func (action UpdateTaskAction) Execute() (updatedTask Task, err error) {
	if action.DTO.ID.String() == "" {
		return updatedTask, appErrors.ValidationError{Field: IDField, Message: "task id should be in the uuid format"}
	}
	updatedTask, err = action.Repository.FindById(action.DTO.ID, WithoutDeletedTasks)
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
	if action.DTO.Priority != 0 && action.DTO.Priority != updatedTask.GetPriority() {
		columnTasks, err := action.Repository.FindForColumn(updatedTask.GetColumnID(), WithoutDeletedTasks)
		if err != nil {
			return updatedTask, err
		}
		for _, task := range columnTasks.Tasks {
			if task.GetPriority() == action.DTO.Priority {
				task.ChangePriority(updatedTask.GetPriority())
				tasksForUpdate.Add(task)
				break
			}
		}
		updatedTask.ChangePriority(action.DTO.Priority)
	}
	tasksForUpdate.Add(updatedTask)
	err = action.Repository.BatchUpdate(tasksForUpdate)
	if err != nil {
		return
	}
	return updatedTask, err
}
