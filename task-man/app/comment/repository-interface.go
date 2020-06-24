package comment

import "github.com/google/uuid"

const (
	WithDeletedComments    = true
	WithoutDeletedComments = false
)

type CommentsRepository interface {
	FindNotDeleted() (CommentsCollection, error)
	FindById(id uuid.UUID) (Comment, error)
	FindForTasks(tasksIDs []uuid.UUID, withDeleted bool) (CommentsCollection, error)
	Save(comment Comment) error
	BatchUpdate(comments CommentsCollection) error
}
