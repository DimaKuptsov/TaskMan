package mock

import (
	"errors"
	"github.com/DimaKuptsov/task-man/app/project"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProjectsRepositoryMock struct {
	projects map[string]project.Project
}

func ProjectRepoMock() ProjectsRepositoryMock {
	factory := project.ProjectsFactory{Validate: validator.New()}
	firstProject, _ := factory.Create(project.CreateDTO{Name: "First project"})
	secondProject, _ := factory.Create(project.CreateDTO{Name: "Second project"})
	thirdProject, _ := factory.Create(project.CreateDTO{Name: "Third project"})
	fourthProject, _ := factory.Create(project.CreateDTO{Name: "Fourth project"})
	fifthProject, _ := factory.Create(project.CreateDTO{Name: "Fifth project"})
	sixthProject, _ := factory.Create(project.CreateDTO{Name: "Sixth project"})
	seventhProject, _ := factory.Create(project.CreateDTO{Name: "Seventh project"})
	fakeProjects := map[string]project.Project{
		firstProject.GetID().String():   firstProject,
		secondProject.GetID().String():  secondProject,
		thirdProject.GetID().String():   thirdProject,
		fourthProject.GetID().String():  fourthProject,
		fifthProject.GetID().String():   fifthProject,
		sixthProject.GetID().String():   sixthProject,
		seventhProject.GetID().String(): seventhProject,
	}
	return ProjectsRepositoryMock{
		projects: fakeProjects,
	}
}

func (r ProjectsRepositoryMock) All() map[string]project.Project {
	return r.projects
}

func (r ProjectsRepositoryMock) FindNotDeleted() (project.ProjectsCollection, error) {
	notDeleted := project.ProjectsCollection{}
	for _, fakeProject := range r.projects {
		if !fakeProject.IsDeleted() {
			notDeleted.Add(fakeProject)
		}
	}
	return notDeleted, nil
}
func (r ProjectsRepositoryMock) FindById(id uuid.UUID) (project project.Project, err error) {
	exist := false
	for projectId, fakeProject := range r.projects {
		if projectId == id.String() {
			project = fakeProject
			exist = true
			break
		}
	}
	if !exist {
		err = errors.New("not exist")
		return project, err
	}
	return project, err
}

func (r ProjectsRepositoryMock) Save(project project.Project) error {
	r.projects[project.GetID().String()] = project
	return nil
}

func (r ProjectsRepositoryMock) Update(project project.Project) error {
	for _, existProject := range r.projects {
		if existProject.GetID() == project.GetID() {
			_ = existProject.ChangeName(project.GetName())
			_ = existProject.ChangeDescription(project.GetDescription())
			if project.IsDeleted() {
				_ = existProject.MarkDeleted()
			}
			r.projects[existProject.GetID().String()] = existProject
			break
		}
	}
	return nil
}
