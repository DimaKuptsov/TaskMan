package task

type CreateTaskAction struct {
	DTO             CreateTaskDTO
	TasksRepository TasksRepository
	TasksFactory    TasksFactory
}

func (action CreateTaskAction) Execute() (task Task, err error) {
	task, err = action.TasksFactory.Create(action.DTO)
	if err != nil {
		return
	}
	err = action.TasksRepository.Save(task)
	if err != nil {
		return
	}
	return task, err
}
