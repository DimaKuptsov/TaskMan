package project

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/app/column"
	"go.uber.org/zap"
)

type CreateProjectAction struct {
	DTO           CreateDTO
	Repository    ProjectsRepository
	Factory       ProjectsFactory
	ColumnService column.ColumnsService
	Logger        *zap.Logger
}

func (action CreateProjectAction) Execute() (newProject Project, err error) {
	newProject, err = action.Factory.Create(action.DTO)
	if err != nil {
		return
	}
	err = action.Repository.Save(newProject)
	if err != nil {
		return
	}
	createColumnDTO := column.CreateDTO{ProjectID: newProject.GetID(), Name: column.DefaultColumnName}
	_, createDefaultColumnErr := action.ColumnService.CreateColumn(createColumnDTO)
	if createDefaultColumnErr != nil {
		errMsg := fmt.Sprintf("not created default column for project %s: %s", newProject.GetID().String(), createDefaultColumnErr.Error())
		action.Logger.Warn(errMsg)
	}
	return newProject, err
}
