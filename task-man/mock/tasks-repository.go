package mock

import (
	"github.com/DimaKuptsov/task-man/app/task"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type TasksRepositoryMock struct {
	tasks map[string]task.Task
}

func TasksRepoMock(columnsRepoMock ColumnsRepositoryMock) TasksRepositoryMock {
	fakeTasks := make(map[string]task.Task)
	factory := task.TasksFactory{Validate: validator.New(), TasksRepository: TasksRepositoryMock{fakeTasks}}
	columns := columnsRepoMock.All()
	for _, fakeColumn := range columns {
		firstTask, _ := factory.Create(task.CreateTaskDTO{ColumnID: fakeColumn.GetID(), Name: "First task", Description: fakeColumn.GetID().String()})
		secondTask, _ := factory.Create(task.CreateTaskDTO{ColumnID: fakeColumn.GetID(), Name: "Second task", Description: fakeColumn.GetID().String()})
		thirdTask, _ := factory.Create(task.CreateTaskDTO{ColumnID: fakeColumn.GetID(), Name: "Third task", Description: fakeColumn.GetID().String()})
		fakeTasks[firstTask.GetID().String()] = firstTask
		fakeTasks[secondTask.GetID().String()] = secondTask
		fakeTasks[thirdTask.GetID().String()] = thirdTask
	}
	return TasksRepositoryMock{fakeTasks}
}

func (r TasksRepositoryMock) All() map[string]task.Task {
	return r.tasks
}

func (r TasksRepositoryMock) FindById(id uuid.UUID) (fakeTask task.Task, err error) {
	for taskId, existTask := range r.tasks {
		if taskId == id.String() {
			fakeTask = existTask
			break
		}
	}
	return fakeTask, err
}

func (r TasksRepositoryMock) FindForColumn(columnID uuid.UUID, withDeleted bool) (tasks task.TasksCollection, err error) {
	for _, existTask := range r.tasks {
		if existTask.IsDeleted() && withDeleted == task.WithoutDeletedTasks {
			continue
		}
		if existTask.GetColumnID() == columnID {
			tasks.Add(existTask)
		}
	}
	return tasks, err
}

func (r TasksRepositoryMock) FindByIds(ids []uuid.UUID) (tasks task.TasksCollection, err error) {
	for _, id := range ids {
		foundTask, _ := r.FindById(id)
		tasks.Add(foundTask)
	}
	return tasks, nil
}

func (r TasksRepositoryMock) Save(task task.Task) error {
	r.tasks[task.GetID().String()] = task
	return nil
}

func (r TasksRepositoryMock) BatchUpdate(tasks task.TasksCollection) error {
	for _, taskForUpdate := range tasks.Tasks {
		existTask, _ := r.FindById(taskForUpdate.GetID())
		_ = existTask.ChangeColumnID(taskForUpdate.GetColumnID())
		_ = existTask.ChangeName(taskForUpdate.GetName())
		_ = existTask.ChangeDescription(taskForUpdate.GetDescription())
		_ = existTask.ChangePriority(taskForUpdate.GetPriority())
		if taskForUpdate.IsDeleted() {
			_ = existTask.MarkDeleted()
		}
		r.tasks[existTask.GetID().String()] = existTask
	}
	return nil
}
