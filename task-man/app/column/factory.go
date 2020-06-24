package column

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type ColumnsFactory struct {
	Validate          *validator.Validate
	ColumnsRepository ColumnsRepository
}

func (f ColumnsFactory) Create(createDTO CreateDTO) (column Column, err error) {
	projectId := createDTO.ProjectID
	if projectId.String() == "" {
		err = errors.New("invalid project id")
	}
	columnName := Name{createDTO.Name}
	err = f.Validate.Struct(columnName)
	if err != nil {
		return
	}
	projectColumns, err := f.ColumnsRepository.FindForProject(projectId, WithoutDeletedColumns)
	if err != nil {
		return
	}
	for _, existColumn := range projectColumns.Columns {
		if existColumn.GetName().String() == columnName.String() {
			err = errors.New("column with the same name already exists")
			break
		}
	}
	if err != nil {
		return
	}
	column = Column{
		id:        uuid.New(),
		projectID: projectId,
		name:      columnName,
		priority:  projectColumns.Len(),
		createdAt: time.Now(),
	}
	return column, err
}
