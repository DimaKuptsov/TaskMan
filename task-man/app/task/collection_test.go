package task

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"testing"
)

func TestSortByPriority(t *testing.T) {
	var tests = getTestTasks()
	for _, testTasks := range tests {
		collection := TasksCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("TasksCollection.SortByPriority: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testTasks...)
		collection.SortByPriority()
		for i, testTask := range collection.Tasks {
			nextTaskIndex := i + 1
			if nextTaskIndex < collection.Len() {
				if testTask.GetPriority() > collection.Tasks[nextTaskIndex].GetPriority() {
					t.Errorf("TasksCollection.SortByPriority: invalid tasks priority sort")
				}
			}
		}
	}
}

func TestLen(t *testing.T) {
	var tests = getTestTasks()
	for _, testTasks := range tests {
		collection := TasksCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("TasksCollection.Len: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testTasks...)
		if !assert.IsEqual(len(testTasks), collection.Len()) {
			t.Errorf("TasksCollection.Len: expected len %v. Got %v", len(testTasks), collection.Len())
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = getTestTasks()
	for _, testTasks := range tests {
		collection := TasksCollection{}
		if !collection.IsEmpty() {
			t.Errorf("TasksCollection.Add: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testTasks...)
		if !assert.IsEqual(len(testTasks), collection.Len()) {
			t.Errorf("TasksCollection.Add: expected collection with len %v. Got with len %v", len(testTasks), collection.Len())
		}
		firstTaskFromTestTasks := testTasks[0]
		firstTaskFromCollection := collection.Tasks[0]
		if !assert.IsEqual(firstTaskFromTestTasks.ID.String(), firstTaskFromCollection.ID.String()) {
			t.Errorf("TasksCollection.Add: expected tasks with equal ids. Got %s and %s", firstTaskFromTestTasks.ID, firstTaskFromCollection.ID)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = getTestTasks()
	for _, testTasks := range tests {
		collection := TasksCollection{}
		if !collection.IsEmpty() {
			t.Errorf("TasksCollection.IsEmpty: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testTasks...)
		if collection.IsEmpty() {
			t.Errorf("TasksCollection.IsEmpty: expected not empty collection. Got empty")
		}
	}
}

func TestTasksCollectionMarshalJSON(t *testing.T) {
	var tests = getTestTasks()
	for _, testTasks := range tests {
		collection := TasksCollection{}
		if !collection.IsEmpty() {
			t.Errorf("TasksCollection.MarshalJSON: expected empty collection. Got with %v len", collection.Len())
		}
		emptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("TasksCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		if string(emptyJson) != "{}" {
			t.Errorf("TasksCollection.MarshalJSON: expected marshal empty collection to '{}'. Got %s", string(emptyJson))
		}
		collection.Add(testTasks...)
		if collection.IsEmpty() {
			t.Errorf("TasksCollection.MarshalJSON: expected not empty collection. Got empty")
		}
		notEmptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("TasksCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("TasksCollection.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
	}
}

func getTestTasks() [][]Task {
	testsNumber := 5
	var tests [][]Task
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		generateTasksNumber := helpers.GenerateIntBetween(1, 10)
		var tasks []Task
		for taskIndex := 0; taskIndex < generateTasksNumber; taskIndex++ {
			task := Task{ID: uuid.New(), Priority: helpers.GenerateIntBetween(1, 10)}
			tasks = append(tasks, task)
		}
		tests = append(tests, tasks)
	}
	return tests
}
