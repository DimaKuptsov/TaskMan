package column

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"testing"
)

func TestSortByPriority(t *testing.T) {
	var tests = getTestColumns()
	for _, testColumns := range tests {
		collection := ColumnsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("ColumnsCollection.SortByPriority: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testColumns...)
		collection.SortByPriority()
		for i, testColumn := range collection.Columns {
			nextColumnIndex := i + 1
			if nextColumnIndex < collection.Len() {
				if testColumn.GetPriority() > collection.Columns[nextColumnIndex].GetPriority() {
					t.Errorf("ColumnsCollection.SortByPriority: invalid columns priority sort")
				}
			}
		}
	}
}

func TestLen(t *testing.T) {
	var tests = getTestColumns()
	for _, testColumns := range tests {
		collection := ColumnsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("ColumnsCollection.Len: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testColumns...)
		if !assert.IsEqual(len(testColumns), collection.Len()) {
			t.Errorf("ColumnsCollection.Len: expected len %v. Got %v", len(testColumns), collection.Len())
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = getTestColumns()
	for _, testColumns := range tests {
		collection := ColumnsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ColumnsCollection.Add: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testColumns...)
		if !assert.IsEqual(len(testColumns), collection.Len()) {
			t.Errorf("ColumnsCollection.Add: expected collection with len %v. Got with len %v", len(testColumns), collection.Len())
		}
		firstColumnFromTestColumns := testColumns[0]
		firstColumnFromCollection := collection.Columns[0]
		if !assert.IsEqual(firstColumnFromTestColumns.ID.String(), firstColumnFromCollection.ID.String()) {
			t.Errorf("ColumnsCollection.Add: expected columns with equal ids. Got %s and %s", firstColumnFromTestColumns.ID, firstColumnFromCollection.ID)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = getTestColumns()
	for _, testColumns := range tests {
		collection := ColumnsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ColumnsCollection.IsEmpty: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testColumns...)
		if collection.IsEmpty() {
			t.Errorf("ColumnsCollection.IsEmpty: expected not empty collection. Got empty")
		}
	}
}

func TestColumnsCollectionMarshalJSON(t *testing.T) {
	var tests = getTestColumns()
	for _, testColumns := range tests {
		collection := ColumnsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("ColumnsCollection.MarshalJSON: expected empty collection. Got with %v len", collection.Len())
		}
		emptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("ColumnsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		if string(emptyJson) != "{}" {
			t.Errorf("ColumnsCollection.MarshalJSON: expected marshal empty collection to '{}'. Got %s", string(emptyJson))
		}
		collection.Add(testColumns...)
		if collection.IsEmpty() {
			t.Errorf("ColumnsCollection.MarshalJSON: expected not empty collection. Got empty")
		}
		notEmptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("ColumnsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("ColumnsCollection.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
	}
}

func getTestColumns() [][]Column {
	testsNumber := 5
	var tests [][]Column
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		generateColumnsNumber := helpers.GenerateIntBetween(1, 10)
		var columns []Column
		for columnIndex := 0; columnIndex < generateColumnsNumber; columnIndex++ {
			column := Column{ID: uuid.New(), Priority: helpers.GenerateIntBetween(1, 10)}
			columns = append(columns, column)
		}
		tests = append(tests, columns)
	}
	return tests
}
