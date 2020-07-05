package helpers

import (
	"strings"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	var tests = []struct {
		len   int
		chars string
	}{
		{1, "a"},
		{45, "random"},
		{521, "someCharacters"},
	}
	for _, testCase := range tests {
		generatedString := GenerateRandomString(testCase.len, testCase.chars)
		if len(generatedString) != testCase.len {
			t.Errorf("GenerateRandomString: expected generated string with len %v. Got with len %v", testCase.len, len(generatedString))
		}
		for _, char := range generatedString {
			if !strings.Contains(testCase.chars, string(char)) {
				t.Errorf("GenerateRandomString: generated string contain not expected char %s. Expected chars: %s", string(char), testCase.chars)
			}
		}
	}
}

func TestGenerateIntBetween(t *testing.T) {
	var tests = []struct {
		min, max int
	}{
		{1, 2},
		{1, 25},
		{5211, 12391},
		{0, 9999},
	}
	for _, testCase := range tests {
		generatedNumber := GenerateIntBetween(testCase.min, testCase.max)
		if generatedNumber < testCase.min {
			t.Errorf("GenerateIntBetween: expected generated number greater then %v. Got %v", testCase.min, generatedNumber)
		}
		if generatedNumber > testCase.max {
			t.Errorf("GenerateIntBetween: expected generated number greate then %v. Got %v", testCase.max, generatedNumber)
		}
	}
}
