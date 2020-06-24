package mock

import (
	"github.com/DimaKuptsov/task-man/app/column"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ColumnsRepositoryMock struct {
	columns map[string]column.Column
}

func ColumnsRepoMock(projectsRepoMock ProjectsRepositoryMock) ColumnsRepositoryMock {
	projects := projectsRepoMock.All()
	fakeRepo := ColumnsRepositoryMock{make(map[string]column.Column)}
	factory := column.ColumnsFactory{Validate: validator.New(), ColumnsRepository: fakeRepo}
	fakeColumns := make(map[string]column.Column)
	for _, fakeProject := range projects {
		firstColumn, _ := factory.Create(column.CreateDTO{ProjectID: fakeProject.GetID(), Name: "First column " + fakeProject.GetID().String()})
		secondColumn, _ := factory.Create(column.CreateDTO{ProjectID: fakeProject.GetID(), Name: "Second column " + fakeProject.GetID().String()})
		thirdColumn, _ := factory.Create(column.CreateDTO{ProjectID: fakeProject.GetID(), Name: "Third column " + fakeProject.GetID().String()})
		fakeColumns[firstColumn.GetID().String()] = firstColumn
		fakeColumns[secondColumn.GetID().String()] = secondColumn
		fakeColumns[thirdColumn.GetID().String()] = thirdColumn
	}
	return ColumnsRepositoryMock{fakeColumns}
}

func (r ColumnsRepositoryMock) All() map[string]column.Column {
	return r.columns
}

func (r ColumnsRepositoryMock) FindById(id uuid.UUID) (fakeColumn column.Column, err error) {
	for columnId, existColumn := range r.columns {
		if columnId == id.String() {
			fakeColumn = existColumn
			break
		}
	}
	return fakeColumn, err
}

func (r ColumnsRepositoryMock) FindForProject(projectID uuid.UUID, withDeleted bool) (columns column.ColumnsCollection, err error) {
	for _, existColumn := range r.columns {
		if existColumn.IsDeleted() && withDeleted == column.WithoutDeletedColumns {
			continue
		}
		if existColumn.GetProjectID() == projectID {
			columns.Add(existColumn)
		}
	}
	return columns, err
}

func (r ColumnsRepositoryMock) Save(column column.Column) error {
	r.columns[column.GetID().String()] = column
	return nil
}

func (r ColumnsRepositoryMock) Update(column column.Column) error {
	for _, existColumn := range r.columns {
		if existColumn.GetID() == column.GetID() {
			_ = existColumn.ChangeName(column.GetName())
			_ = existColumn.ChangePriority(column.GetPriority())
			if column.IsDeleted() {
				_ = existColumn.MarkDeleted()
			}
			r.columns[existColumn.GetID().String()] = existColumn
			break
		}
	}
	return nil
}

func (r ColumnsRepositoryMock) BatchUpdate(columns column.ColumnsCollection) error {
	for _, columnForUpdate := range columns.Columns {
		_ = r.Update(columnForUpdate)
	}
	return nil
}
