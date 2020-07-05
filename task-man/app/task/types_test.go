package task

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/helpers"
	"testing"
)

func TestNameString(t *testing.T) {
	tests := getTestNames()
	for _, testName := range tests {
		stringRepresentation := fmt.Sprintf("%s", testName)
		if testName.Name != stringRepresentation {
			t.Errorf("Name.String: string representation not equal with name field. Representation %s. Field %s", stringRepresentation, testName.Name)
		}
	}
}

func TestNameScan(t *testing.T) {
	tests := getTestNames()
	for _, testName := range tests {
		newName := helpers.GenerateRandomString(50, "zxy")
		err := testName.Scan([]byte(newName))
		if err != nil {
			t.Errorf("Name.Scan: not scan new name. Error %s", err.Error())
		}
		if testName.Name != newName {
			t.Errorf("Name.Scan: not scan new name. Given name %s. Field %s", newName, testName.Name)
		}
	}
}

func TestNameValue(t *testing.T) {
	tests := getTestNames()
	for _, testName := range tests {
		value, err := testName.Value()
		if err != nil {
			t.Errorf("Name.Value: got error %s", err.Error())
		}
		if testName.Name != value {
			t.Errorf("Name.Value: returned not valid value. Returned value %s. Field %s", value, testName.Name)
		}
	}
}

func getTestNames() []Name {
	testsNumber := 5
	var tests []Name
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		name := Name{
			Name: helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 500), "abc"),
		}
		tests = append(tests, name)
	}
	return tests
}

func TestDescriptionString(t *testing.T) {
	tests := getTestDescriptions()
	for _, testDescription := range tests {
		stringRepresentation := fmt.Sprintf("%s", testDescription)
		if testDescription.Description != stringRepresentation {
			t.Errorf("Description.String: string representation not equal with description field. Representation %s. Field %s", stringRepresentation, testDescription.Description)
		}
	}
}

func TestDescriptionScan(t *testing.T) {
	tests := getTestDescriptions()
	for _, testDescription := range tests {
		newDescription := helpers.GenerateRandomString(50, "zxy")
		err := testDescription.Scan([]byte(newDescription))
		if err != nil {
			t.Errorf("Description.Scan: not scan new description. Error %s", err.Error())
		}
		if testDescription.Description != newDescription {
			t.Errorf("Description.Scan: not scan new desctription. Given name %s. Field %s", newDescription, testDescription.Description)
		}
	}
}

func TestDescriptionValue(t *testing.T) {
	tests := getTestDescriptions()
	for _, testDescription := range tests {
		value, err := testDescription.Value()
		if err != nil {
			t.Errorf("Description.Value: got error %s", err.Error())
		}
		if testDescription.Description != value {
			t.Errorf("Description.Value: returned not valid value. Returned value %s. Field %s", value, testDescription.Description)
		}
	}
}

func getTestDescriptions() []Description {
	testsNumber := 5
	var tests []Description
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		description := Description{
			Description: helpers.GenerateRandomString(helpers.GenerateIntBetween(0, 5000), "abc"),
		}
		tests = append(tests, description)
	}
	return tests
}
