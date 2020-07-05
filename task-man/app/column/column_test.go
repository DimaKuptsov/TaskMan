package column

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestGetID(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if testColumn.GetID() != testColumn.ID {
			t.Errorf("Column.GetID: not equal id. From method %s. ID field %s", testColumn.GetID(), testColumn.ID)
		}
	}
}

func TestGetProjectID(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		testColumn.ProjectID = uuid.New()
		if testColumn.GetProjectID() != testColumn.ProjectID {
			t.Errorf("Column.GetProjectID: not equal project id. From method %s. ProjectID field %s", testColumn.GetProjectID(), testColumn.ProjectID)
		}
	}
}

func TestGetName(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if testColumn.GetName().String() != testColumn.Name.String() {
			t.Errorf("Column.GetName: not equal name. From method %s. Name field %s", testColumn.GetName(), testColumn.Name)
		}
		newName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 255), "abc")
		err := testColumn.ChangeName(Name{Name: newName})
		if err != nil {
			t.Errorf("Column.GetName: name not changed to valid name")
		}
		if testColumn.GetName().String() != testColumn.Name.String() {
			t.Errorf("Column.GetName: not equal name. From method %s. Name field %s.", testColumn.GetName(), testColumn.Name)
		}
	}
}

func TestChangeName(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		emptyName := Name{}
		err := testColumn.ChangeName(emptyName)
		if err == nil {
			t.Errorf("Column.ChangeName: name changed to empty")
		}
		toLongName := Name{Name: helpers.GenerateRandomString(256, "abc")}
		err = testColumn.ChangeName(toLongName)
		if err == nil {
			t.Errorf("Column.ChangeName: name changed to longer than allowed")
		}
		validName := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 255), "abc")
		name := Name{Name: validName}
		err = testColumn.ChangeName(name)
		if err != nil || testColumn.Name.String() != name.String() {
			t.Errorf("Column.ChangeName: name not changed to valid name")
		}
	}
}

func TestGetPriority(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if testColumn.GetPriority() != testColumn.Priority {
			t.Errorf("Column.GetPriority: not equal priority. From method %v. Priority field %v.", testColumn.GetPriority(), testColumn.Priority)
		}
		newPriority := testColumn.Priority + helpers.GenerateIntBetween(1, 20)
		testColumn.ChangePriority(newPriority)
		if testColumn.GetPriority() != testColumn.Priority {
			t.Errorf("Column.GetPriority: not equal priority. From method %v. Priority field %v.", testColumn.GetPriority(), testColumn.Priority)
		}
	}
}

func TestChangePriority(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		priority := testColumn.Priority
		testColumn.ChangePriority(priority + 1)
		if testColumn.Priority == priority {
			t.Errorf("Column.ChangePriority: not changed priority.")
		}
	}
}

func TestIsDeleted(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if testColumn.IsDeleted() {
			t.Errorf("Column.IsDeleted: expected not deleted column. Got deleted at %s", testColumn.DeletedAt)
		}
		testColumn.MarkDeleted()
		if !testColumn.IsDeleted() {
			t.Errorf("Column.IsDeleted: expected deleted column. Got not deleted")
		}
		if testColumn.IsDeleted() != !testColumn.DeletedAt.IsZero() {
			t.Errorf("Column.IsDeleted: status is not defined correctly")
		}
	}
}

func TestMarkDeleted(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if !testColumn.DeletedAt.IsZero() {
			t.Errorf("Column.MarkDeleted: expected not deleted column. Got deleted at %s", testColumn.DeletedAt)
		}
		testColumn.MarkDeleted()
		if testColumn.DeletedAt.IsZero() {
			t.Errorf("Column.MarkDeleted: expected deleted column. Got not deleted")
		}
	}
}

func TestMarkUpdated(t *testing.T) {
	tests := getTestColumn()
	for _, testColumn := range tests {
		if !testColumn.UpdatedAt.IsZero() {
			t.Errorf("Column.MarkUpdated: expected not updated column. Got updated at %s", testColumn.UpdatedAt)
		}
		testColumn.markUpdated()
		if testColumn.UpdatedAt.IsZero() {
			t.Errorf("Column.MarkUpdated: expected updated column. Got not updated")
		}
	}
}

func TestColumnMarshalJSON(t *testing.T) {
	var tests = getTestColumn()
	for _, testColumn := range tests {
		notEmptyJson, err := testColumn.MarshalJSON()
		if err != nil {
			t.Errorf("Column.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("Column.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
		if !strings.Contains(notEmptyJsonString, "\"id\"") {
			t.Errorf("Column.MarshalJSON: test column has id and json must contain 'id' field but not contain. Got string %s", notEmptyJsonString)
		}
		testColumn.markUpdated()
		notEmptyJson, err = testColumn.MarshalJSON()
		if err != nil {
			t.Errorf("Column.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString = string(notEmptyJson)
		if !strings.Contains(notEmptyJsonString, "\"updated_at\"") {
			t.Errorf("Column.MarshalJSON: test column has update date and json must contain 'updated_at' field but not contain. Got string %s", notEmptyJsonString)
		}
	}
}

func getTestColumn() []Column {
	testsNumber := 5
	var tests []Column
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		column := Column{ID: uuid.New()}
		tests = append(tests, column)
	}
	return tests
}
