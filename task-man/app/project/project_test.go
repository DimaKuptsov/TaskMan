package project

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestGetID(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if testProject.GetID() != testProject.ID {
			t.Errorf("Project.GetID: not equal id. From method %s. ID field %s", testProject.GetID(), testProject.ID)
		}
	}
}

func TestGetName(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if testProject.GetName().String() != testProject.Name.String() {
			t.Errorf("Project.GetName: not equal name. From method %s. Name field %s", testProject.GetName(), testProject.Name)
		}
		newName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 500), "abc")
		err := testProject.ChangeName(Name{Name: newName})
		if err != nil {
			t.Errorf("Project.GetName: name not changed to valid name")
		}
		if testProject.GetName().String() != testProject.Name.String() {
			t.Errorf("Project.GetName: not equal name. From method %s. Name field %s.", testProject.GetName(), testProject.Name)
		}
	}
}

func TestChangeName(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		emptyName := Name{}
		err := testProject.ChangeName(emptyName)
		if err == nil {
			t.Errorf("Project.ChangeName: name changed to empty")
		}
		toLongName := Name{Name: helpers.GenerateRandomString(501, "abc")}
		err = testProject.ChangeName(toLongName)
		if err == nil {
			t.Errorf("Project.ChangeName: name changed to longer than allowed")
		}
		validName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 500), "abc")
		name := Name{Name: validName}
		err = testProject.ChangeName(name)
		if err != nil || testProject.Name.String() != name.String() {
			t.Errorf("Project.ChangeName: name not changed to valid name")
		}
	}
}

func TestGetDescription(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if testProject.GetDescription().String() != testProject.Description.String() {
			t.Errorf("Project.GetDescription: not equal description. From method %s. Description field %s", testProject.GetDescription(), testProject.Description)
		}
		newDescription := helpers.GenerateRandomString(helpers.GenerateIntBetween(0, 1000), "abc")
		err := testProject.ChangeDescription(Description{Description: newDescription})
		if err != nil {
			t.Errorf("Project.GetDescription: description not changed to valid desctription")
		}
		if testProject.GetDescription().String() != testProject.Description.String() {
			t.Errorf("Project.GetDescription: not equal description. From method %s. Description field %s.", testProject.GetDescription(), testProject.Description)
		}
	}
}

func TestChangeDescription(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		emptyDescription := Description{}
		err := testProject.ChangeDescription(emptyDescription)
		if err != nil {
			t.Errorf("Project.ChangeDescription: description not changed to empty")
		}
		toLongDescription := Description{Description: helpers.GenerateRandomString(1001, "abc")}
		err = testProject.ChangeDescription(toLongDescription)
		if err == nil {
			t.Errorf("Project.ChangeDescription: description changed to longer than allowed")
		}
		validDescription := helpers.GenerateRandomString(helpers.GenerateIntBetween(0, 1000), "abc")
		description := Description{Description: validDescription}
		err = testProject.ChangeDescription(description)
		if err != nil || testProject.Description.String() != description.String() {
			t.Errorf("Project.ChangeDescription: description not changed to valid description")
		}
	}
}

func TestGetCreateTime(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if testProject.GetCreateTime() != testProject.CreatedAt {
			t.Errorf("Project.GetCreateTime: not equal create date. From method %s. CreatedAt field %s.", testProject.GetCreateTime(), testProject.CreatedAt)
		}
	}
}

func TestIsDeleted(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if testProject.IsDeleted() {
			t.Errorf("Project.IsDeleted: expected not deleted project. Got deleted at %s", testProject.DeletedAt)
		}
		testProject.MarkDeleted()
		if !testProject.IsDeleted() {
			t.Errorf("Project.IsDeleted: expected deleted project. Got not deleted")
		}
		if testProject.IsDeleted() != !testProject.DeletedAt.IsZero() {
			t.Errorf("Project.IsDeleted: status is not defined correctly")
		}
	}
}

func TestMarkDeleted(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if !testProject.DeletedAt.IsZero() {
			t.Errorf("Project.MarkDeleted: expected not deleted project. Got deleted at %s", testProject.DeletedAt)
		}
		testProject.MarkDeleted()
		if testProject.DeletedAt.IsZero() {
			t.Errorf("Project.MarkDeleted: expected deleted project. Got not deleted")
		}
	}
}

func TestMarkUpdated(t *testing.T) {
	tests := getTestProject()
	for _, testProject := range tests {
		if !testProject.UpdatedAt.IsZero() {
			t.Errorf("Project.MarkUpdated: expected not updated project. Got updated at %s", testProject.UpdatedAt)
		}
		testProject.markUpdated()
		if testProject.UpdatedAt.IsZero() {
			t.Errorf("Project.MarkUpdated: expected updated project. Got not updated")
		}
	}
}

func TestProjectMarshalJSON(t *testing.T) {
	var tests = getTestProject()
	for _, testProject := range tests {
		notEmptyJson, err := testProject.MarshalJSON()
		if err != nil {
			t.Errorf("Project.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("Project.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
		if !strings.Contains(notEmptyJsonString, "\"id\"") {
			t.Errorf("Project.MarshalJSON: test project has id and json must contain 'id' field but not contain. Got string %s", notEmptyJsonString)
		}
		testProject.markUpdated()
		notEmptyJson, err = testProject.MarshalJSON()
		if err != nil {
			t.Errorf("Project.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString = string(notEmptyJson)
		if !strings.Contains(notEmptyJsonString, "\"updated_at\"") {
			t.Errorf("Project.MarshalJSON: test project has update date and json must contain 'updated_at' field but not contain. Got string %s", notEmptyJsonString)
		}
	}
}

func getTestProject() []Project {
	testsNumber := 5
	var tests []Project
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		project := Project{ID: uuid.New()}
		tests = append(tests, project)
	}
	return tests
}
