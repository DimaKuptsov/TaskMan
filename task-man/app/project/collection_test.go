package project

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"testing"
)

func TestSortByCreateTime(t *testing.T) {
	var tests = getTestProjects()
	for _, testProjects := range tests {
		collection := ProjectsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("ProjectsCollection.SortByCreateTime: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testProjects...)
		collection.SortByCreateTime()
		for i, testProject := range collection.Projects {
			nextProjectIndex := i + 1
			if nextProjectIndex < collection.Len() {
				if testProject.GetCreateTime().After(collection.Projects[nextProjectIndex].GetCreateTime()) {
					t.Errorf("ProjectsCollection.SortByCreateTime: invalid projects sort")
				}
			}
		}
	}
}

func TestLen(t *testing.T) {
	var tests = getTestProjects()
	for _, testProjects := range tests {
		collection := ProjectsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("ProjectsCollection.Len: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testProjects...)
		if !assert.IsEqual(len(testProjects), collection.Len()) {
			t.Errorf("ProjectsCollection.Len: expected len %v. Got %v", len(testProjects), collection.Len())
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = getTestProjects()
	for _, testProjects := range tests {
		collection := ProjectsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ProjectsCollection.Add: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testProjects...)
		if !assert.IsEqual(len(testProjects), collection.Len()) {
			t.Errorf("ProjectsCollection.Add: expected collection with len %v. Got with len %v", len(testProjects), collection.Len())
		}
		firstProjectFromTestProjects := testProjects[0]
		firstProjectFromCollection := collection.Projects[0]
		if !assert.IsEqual(firstProjectFromTestProjects.ID.String(), firstProjectFromCollection.ID.String()) {
			t.Errorf("ProjectsCollection.Add: expected projects with equal ids. Got %s and %s", firstProjectFromTestProjects.ID, firstProjectFromCollection.ID)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = getTestProjects()
	for _, testProjects := range tests {
		collection := ProjectsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ProjectsCollection.IsEmpty: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testProjects...)
		if collection.IsEmpty() {
			t.Errorf("ProjectsCollection.IsEmpty: expected not empty collection. Got empty")
		}
	}
}

func TestProjectsCollectionMarshalJSON(t *testing.T) {
	var tests = getTestProjects()
	for _, testProjects := range tests {
		collection := ProjectsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ProjectsCollection.MarshalJSON: expected empty collection. Got with %v len", collection.Len())
		}
		emptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("ProjectsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		if string(emptyJson) != "{}" {
			t.Errorf("ProjectsCollection.MarshalJSON: expected marshal empty collection to '{}'. Got %s", string(emptyJson))
		}
		collection.Add(testProjects...)
		if collection.IsEmpty() {
			t.Errorf("ProjectsCollection.MarshalJSON: expected not empty collection. Got empty")
		}
		notEmptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("ProjectsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("ProjectsCollection.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
	}
}

func getTestProjects() [][]Project {
	testsNumber := 5
	var tests [][]Project
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		generateProjectsNumber := helpers.GenerateIntBetween(1, 10)
		var projects []Project
		for projectIndex := 0; projectIndex < generateProjectsNumber; projectIndex++ {
			project := Project{ID: uuid.New()}
			projects = append(projects, project)
		}
		tests = append(tests, projects)
	}
	return tests
}
