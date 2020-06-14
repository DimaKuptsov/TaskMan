package column

import "github.com/google/uuid"

const (
	WithDeletedColumns    = true
	WithoutDeletedColumns = false
)

type ColumnsRepository interface {
	FindById(id uuid.UUID) (Column, error)
	FindForProject(projectID uuid.UUID, withDeleted bool) (ColumnsCollection, error)
	Save(column Column) error
	Update(column Column) error
	BatchUpdate(columns ColumnsCollection) error
}
