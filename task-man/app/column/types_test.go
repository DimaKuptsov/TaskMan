package column

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/helpers"
	"testing"
)

func TestString(t *testing.T) {
	tests := getTestNames()
	for _, testName := range tests {
		stringRepresentation := fmt.Sprintf("%s", testName)
		if testName.Name != stringRepresentation {
			t.Errorf("Name.String: string representation not equal with name field. Representation %s. Field %s", stringRepresentation, testName.Name)
		}
	}
}

func TestScan(t *testing.T) {
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

func TestValue(t *testing.T) {
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
			Name: helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 255), "abc"),
		}
		tests = append(tests, name)
	}
	return tests
}
