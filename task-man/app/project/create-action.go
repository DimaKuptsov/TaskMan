package project

import "github.com/DimaKuptsov/task-man/app/column"

type CreateProjectAction struct {
	DTO           CreateDTO
	Repository    ProjectsRepository
	Factory       ProjectsFactory
	ColumnService column.ColumnsService
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
		//TODO log
	}
	return newProject, err
}
