package project

import "github.com/DimaKuptsov/task-man/app/column"

type DeleteProjectAction struct {
	DTO            DeleteDTO
	Repository     ProjectsRepository
	ColumnsService column.ColumnsService
}

func (action DeleteProjectAction) Execute() error {
	project, err := action.Repository.FindById(action.DTO.ID, WithoutDeletedProjects)
	if err != nil {
		return nil
	}
	project.MarkDeleted()
	deleteProjectColumnsDTO := column.DeleteProjectColumnsDTO{ProjectID: project.GetID()}
	err = action.Repository.Update(project)
	if err != nil {
		return err
	}
	return action.ColumnsService.DeleteProjectColumns(deleteProjectColumnsDTO)
}
