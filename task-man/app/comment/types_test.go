package comment

import (
	"fmt"
	"github.com/DimaKuptsov/task-man/helpers"
	"testing"
)

func TestString(t *testing.T) {
	tests := getTestTexts()
	for _, testText := range tests {
		stringRepresentation := fmt.Sprintf("%s", testText)
		if testText.Text != stringRepresentation {
			t.Errorf("Text.String: string representation not equal with text field. Representation %s. Field %s", stringRepresentation, testText.Text)
		}
	}
}

func TestScan(t *testing.T) {
	tests := getTestTexts()
	for _, testText := range tests {
		newText := helpers.GenerateRandomString(500, "zxy")
		err := testText.Scan([]byte(newText))
		if err != nil {
			t.Errorf("Text.Scan: not scan new text. Error %s", err.Error())
		}
		if testText.Text != newText {
			t.Errorf("Text.Scan: not scan new text. Given name %s. Field %s", newText, testText.Text)
		}
	}
}

func TestValue(t *testing.T) {
	tests := getTestTexts()
	for _, testText := range tests {
		value, err := testText.Value()
		if err != nil {
			t.Errorf("Text.Value: got error %s", err.Error())
		}
		if testText.Text != value {
			t.Errorf("Text.Value: returned not valid value. Returned value %s. Field %s", value, testText.Text)
		}
	}
}

func getTestTexts() []Text {
	testsNumber := 5
	var tests []Text
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		text := Text{
			Text: helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 5000), "abc"),
		}
		tests = append(tests, text)
	}
	return tests
}
