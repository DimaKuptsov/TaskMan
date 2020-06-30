package column

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
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
		return column, appErrors.ValidationError{Field: ProjectIDField, Message: "project id should be in the uuid format"}
	}
	columnName := Name{createDTO.Name}
	err = f.Validate.Struct(columnName)
	if err != nil {
		return column, appErrors.ValidationError{Field: NameField, Message: err.Error()}
	}
	projectColumns, err := f.ColumnsRepository.FindForProject(projectId, WithoutDeletedColumns)
	if err != nil {
		return
	}
	for _, existColumn := range projectColumns.Columns {
		if existColumn.GetName().String() == columnName.String() {
			err = appErrors.ValidationError{Field: NameField, Message: "column with the same name already exists"}
			break
		}
	}
	if err != nil {
		return
	}
	priority := projectColumns.Len() + 1
	column = Column{
		ID:        uuid.New(),
		ProjectID: projectId,
		Name:      columnName,
		Priority:  priority,
		CreatedAt: time.Now(),
	}
	return column, err
}
