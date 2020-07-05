package comment

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
	"testing"
)

func TestSortByCreateTime(t *testing.T) {
	var tests = getTestComments()
	for _, testComments := range tests {
		collection := CommentsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("CommentsCollection.SortByCreateTime: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testComments...)
		collection.SortByCreateTime()
		for i, testComment := range collection.Comments {
			nextCommentIndex := i + 1
			if nextCommentIndex < collection.Len() {
				if testComment.GetCreateTime().After(collection.Comments[nextCommentIndex].GetCreateTime()) {
					t.Errorf("CommentsCollection.SortByCreateTime: invalid comments sort")
				}
			}
		}
	}
}

func TestLen(t *testing.T) {
	var tests = getTestComments()
	for _, testComments := range tests {
		collection := CommentsCollection{}
		if !assert.IsEqual(0, collection.Len()) {
			t.Errorf("CommentsCollection.Len: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testComments...)
		if !assert.IsEqual(len(testComments), collection.Len()) {
			t.Errorf("CommentsCollection.Len: expected len %v. Got %v", len(testComments), collection.Len())
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = getTestComments()
	for _, testComments := range tests {
		collection := CommentsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("CommentsCollection.Add: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testComments...)
		if !assert.IsEqual(len(testComments), collection.Len()) {
			t.Errorf("CommentsCollection.Add: expected collection with len %v. Got with len %v", len(testComments), collection.Len())
		}
		firstCommentFromTestComments := testComments[0]
		firstCommentFromCollection := collection.Comments[0]
		if !assert.IsEqual(firstCommentFromTestComments.ID.String(), firstCommentFromCollection.ID.String()) {
			t.Errorf("CommentsCollection.Add: expected comments with equal ids. Got %s and %s", firstCommentFromTestComments.ID, firstCommentFromCollection.ID)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = getTestComments()
	for _, testComments := range tests {
		collection := CommentsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("CommentsCollection.IsEmpty: expected empty collection. Got with %v len", collection.Len())
		}
		collection.Add(testComments...)
		if collection.IsEmpty() {
			t.Errorf("CommentsCollection.IsEmpty: expected not empty collection. Got empty")
		}
	}
}

func TestCommentsCollectionMarshalJSON(t *testing.T) {
	var tests = getTestComments()
	for _, testComments := range tests {
		collection := CommentsCollection{}
		if !collection.IsEmpty() {
			t.Errorf("CommentsCollection.MarshalJSON: expected empty collection. Got with %v len", collection.Len())
		}
		emptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("CommentsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		if string(emptyJson) != "{}" {
			t.Errorf("CommentsCollection.MarshalJSON: expected marshal empty collection to '{}'. Got %s", string(emptyJson))
		}
		collection.Add(testComments...)
		if collection.IsEmpty() {
			t.Errorf("CommentsCollection.MarshalJSON: expected not empty collection. Got empty")
		}
		notEmptyJson, err := collection.MarshalJSON()
		if err != nil {
			t.Errorf("CommentsCollection.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("CommentsCollection.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
	}
}

func getTestComments() [][]Comment {
	testsNumber := 5
	var tests [][]Comment
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		generateCommentsNumber := helpers.GenerateIntBetween(1, 10)
		var comments []Comment
		for commentIndex := 0; commentIndex < generateCommentsNumber; commentIndex++ {
			text := Text{Text: helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 5000), "characters")}
			comment := Comment{ID: uuid.New(), Text: text}
			comments = append(comments, comment)
		}
		tests = append(tests, comments)
	}
	return tests
}
