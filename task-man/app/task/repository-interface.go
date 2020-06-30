package task

import "github.com/google/uuid"

const (
	WithDeletedTasks    = true
	WithoutDeletedTasks = false
)

type TasksRepository interface {
	FindById(id uuid.UUID, withDeleted bool) (Task, error)
	FindByIds(ids []uuid.UUID, withDeleted bool) (TasksCollection, error)
	FindForColumn(columnID uuid.UUID, withDeleted bool) (TasksCollection, error)
	Save(task Task) error
	BatchUpdate(tasks TasksCollection) error
}
