package task

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestGetID(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if testTask.GetID() != testTask.ID {
			t.Errorf("Task.GetID: not equal id. From method %s. ID field %s", testTask.GetID(), testTask.ID)
		}
	}
}

func TestGetColumnID(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		testTask.ColumnID = uuid.New()
		if testTask.GetColumnID() != testTask.ColumnID {
			t.Errorf("Task.GetColumnID: not equal column id. From method %s. ColumnID field %s", testTask.GetColumnID(), testTask.ColumnID)
		}
	}
}

func TestChangeColumnID(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		columnId := uuid.New()
		err := testTask.ChangeColumnID(columnId)
		if err != nil || testTask.ColumnID != columnId {
			t.Errorf("Task.ChangeColumnID: column id not changed to valid uuid")
		}
	}
}

func TestGetName(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if testTask.GetName().String() != testTask.Name.String() {
			t.Errorf("Task.GetName: not equal name. From method %s. Name field %s", testTask.GetName(), testTask.Name)
		}
		newName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 500), "abc")
		err := testTask.ChangeName(Name{Name: newName})
		if err != nil {
			t.Errorf("Task.GetName: name not changed to valid name")
		}
		if testTask.GetName().String() != testTask.Name.String() {
			t.Errorf("Task.GetName: not equal name. From method %s. Name field %s.", testTask.GetName(), testTask.Name)
		}
	}
}

func TestChangeName(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		emptyName := Name{}
		err := testTask.ChangeName(emptyName)
		if err == nil {
			t.Errorf("Task.ChangeName: name changed to empty")
		}
		toLongName := Name{Name: helpers.GenerateRandomString(501, "abc")}
		err = testTask.ChangeName(toLongName)
		if err == nil {
			t.Errorf("Task.ChangeName: name changed to longer than allowed")
		}
		validName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 500), "abc")
		name := Name{Name: validName}
		err = testTask.ChangeName(name)
		if err != nil || testTask.Name.String() != name.String() {
			t.Errorf("Task.ChangeName: name not changed to valid name")
		}
	}
}

func TestGetDescription(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if testTask.GetDescription().String() != testTask.Description.String() {
			t.Errorf("Task.GetDescription: not equal description. From method %s. Description field %s", testTask.GetDescription(), testTask.Description)
		}
		newDescription := helpers.GenerateRandomString(helpers.GenerateIntBetween(0, 5000), "abc")
		err := testTask.ChangeDescription(Description{Description: newDescription})
		if err != nil {
			t.Errorf("Task.GetDescription: description not changed to valid description")
		}
		if testTask.GetDescription().String() != testTask.Description.String() {
			t.Errorf("Task.GetDescription: not equal description. From method %s. Description field %s.", testTask.GetDescription(), testTask.Description)
		}
	}
}

func TestChangeDescription(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		emptyDescription := Description{}
		err := testTask.ChangeDescription(emptyDescription)
		if err != nil {
			t.Errorf("Task.ChangeDescription: description not changed to empty")
		}
		toLongDescription := Description{Description: helpers.GenerateRandomString(5001, "abc")}
		err = testTask.ChangeDescription(toLongDescription)
		if err == nil {
			t.Errorf("Task.ChangeDescription: description changed to longer than allowed")
		}
		validDescription := helpers.GenerateRandomString(helpers.GenerateIntBetween(0, 5000), "abc")
		description := Description{Description: validDescription}
		err = testTask.ChangeDescription(description)
		if err != nil || testTask.Description.String() != description.String() {
			t.Errorf("Task.ChangeDescription: description not changed to valid description")
		}
	}
}

func TestGetPriority(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if testTask.GetPriority() != testTask.Priority {
			t.Errorf("Task.GetPriority: not equal priority. From method %v. Priority field %v.", testTask.GetPriority(), testTask.Priority)
		}
		newPriority := testTask.Priority + helpers.GenerateIntBetween(1, 20)
		testTask.ChangePriority(newPriority)
		if testTask.GetPriority() != testTask.Priority {
			t.Errorf("Task.GetPriority: not equal priority. From method %v. Priority field %v.", testTask.GetPriority(), testTask.Priority)
		}
	}
}

func TestChangePriority(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		priority := testTask.Priority
		testTask.ChangePriority(priority + 1)
		if testTask.Priority == priority {
			t.Errorf("Task.ChangePriority: not changed priority.")
		}
	}
}

func TestIsDeleted(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if testTask.IsDeleted() {
			t.Errorf("Task.IsDeleted: expected not deleted task. Got deleted at %s", testTask.DeletedAt)
		}
		testTask.MarkDeleted()
		if !testTask.IsDeleted() {
			t.Errorf("Task.IsDeleted: expected deleted task. Got not deleted")
		}
		if testTask.IsDeleted() != !testTask.DeletedAt.IsZero() {
			t.Errorf("Task.IsDeleted: status is not defined correctly")
		}
	}
}

func TestMarkDeleted(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if !testTask.DeletedAt.IsZero() {
			t.Errorf("Task.MarkDeleted: expected not deleted task. Got deleted at %s", testTask.DeletedAt)
		}
		testTask.MarkDeleted()
		if testTask.DeletedAt.IsZero() {
			t.Errorf("Task.MarkDeleted: expected deleted task. Got not deleted")
		}
	}
}

func TestMarkUpdated(t *testing.T) {
	tests := getTestTask()
	for _, testTask := range tests {
		if !testTask.UpdatedAt.IsZero() {
			t.Errorf("Task.MarkUpdated: expected not updated task. Got updated at %s", testTask.UpdatedAt)
		}
		testTask.markUpdated()
		if testTask.UpdatedAt.IsZero() {
			t.Errorf("Task.MarkUpdated: expected updated task. Got not updated")
		}
	}
}

func TestTaskMarshalJSON(t *testing.T) {
	var tests = getTestTask()
	for _, testTask := range tests {
		notEmptyJson, err := testTask.MarshalJSON()
		if err != nil {
			t.Errorf("Task.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("Task.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
		if !strings.Contains(notEmptyJsonString, "\"id\"") {
			t.Errorf("Task.MarshalJSON: test task has id and json must contain 'id' field but not contain. Got string %s", notEmptyJsonString)
		}
		testTask.markUpdated()
		notEmptyJson, err = testTask.MarshalJSON()
		if err != nil {
			t.Errorf("Task.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString = string(notEmptyJson)
		if !strings.Contains(notEmptyJsonString, "\"updated_at\"") {
			t.Errorf("Task.MarshalJSON: test task has update date and json must contain 'updated_at' field but not contain. Got string %s", notEmptyJsonString)
		}
	}
}

func getTestTask() []Task {
	testsNumber := 5
	var tests []Task
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		task := Task{ID: uuid.New()}
		tests = append(tests, task)
	}
	return tests
}
