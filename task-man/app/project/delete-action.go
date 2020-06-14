package project

import "github.com/DimaKuptsov/task-man/app/column"

type DeleteProjectAction struct {
	DTO            DeleteDTO
	Repository     ProjectsRepository
	ColumnsService column.ColumnsService
}

func (action DeleteProjectAction) Execute() error {
	project, err := action.Repository.FindById(action.DTO.ID)
	if err != nil {
		return nil
	}
	err = project.MarkDeleted()
	if err != nil {
		return err
	}
	deleteProjectColumnsDTO := column.DeleteProjectColumnsDTO{ProjectID: project.GetID()}
	err = action.Repository.Update(project)
	if err != nil {
		return err
	}
	return action.ColumnsService.DeleteProjectColumns(deleteProjectColumnsDTO)
}
