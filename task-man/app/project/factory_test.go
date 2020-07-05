package project

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestCreate(t *testing.T) {
	validate := validator.New()
	factory := ProjectsFactory{Validate: validate}
	createDTO := CreateDTO{}
	_, err := factory.Create(createDTO)
	if validateErr, ok := err.(appErrors.ValidationError); !ok || validateErr.GetField() != NameField {
		t.Error("ProjectsFactory.Create: expected name validation err.")
	}
	createDTO.Name = helpers.GenerateRandomString(100, "name")
	createDTO.Description = helpers.GenerateRandomString(1001, "invalid description")
	_, err = factory.Create(createDTO)
	if validateErr, ok := err.(appErrors.ValidationError); !ok || validateErr.GetField() != DescriptionField {
		t.Error("ProjectsFactory.Create: expected description validation err.")
	}
	createDTO.Description = helpers.GenerateRandomString(1000, "valid description")
	createdProject, err := factory.Create(createDTO)
	if err != nil {
		t.Errorf("ProjectsFactory.Create: expected no errors. Got %s", err.Error())
	}
	if createdProject.Name.String() != createDTO.Name {
		t.Errorf("ProjectsFactory.Create: expected project with name %s. Got with name %s", createDTO.Name, createdProject.Name)
	}
	if createdProject.Description.String() != createDTO.Description {
		t.Errorf("ProjectsFactory.Create: expected project with desctiprion %s. Got with desctiprion %s", createDTO.Description, createdProject.Description)
	}
}
