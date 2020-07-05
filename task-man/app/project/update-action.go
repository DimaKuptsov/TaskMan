package project

type UpdateProjectAction struct {
	DTO        UpdateDTO
	Repository ProjectsRepository
}

func (action UpdateProjectAction) Execute() (updatedProject Project, err error) {
	updatedProject, err = action.Repository.FindById(action.DTO.ID, WithoutDeletedProjects)
	if err != nil {
		return
	}
	if action.DTO.Name != updatedProject.GetName().String() {
		newName := Name{action.DTO.Name}
		err = updatedProject.ChangeName(newName)
		if err != nil {
			return
		}
	}
	if action.DTO.Description != updatedProject.GetDescription().String() {
		newDescription := Description{action.DTO.Description}
		err = updatedProject.ChangeDescription(newDescription)
		if err != nil {
			return
		}
	}
	err = action.Repository.Update(updatedProject)
	return updatedProject, err
}
