package comment

import (
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/google/uuid"
	"strings"
	"testing"
)

func TestGetID(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if testComment.GetID() != testComment.ID {
			t.Errorf("Comment.GetID: not equal id. From method %s. ID field %s", testComment.GetID(), testComment.ID)
		}
	}
}

func TestGetTaskID(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		testComment.TaskID = uuid.New()
		if testComment.GetTaskID() != testComment.TaskID {
			t.Errorf("Comment.GetTaskID: not equal task id. From method %s. TaskID field %s", testComment.GetTaskID(), testComment.TaskID)
		}
	}
}

func TestGetText(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if testComment.GetText().String() != testComment.Text.String() {
			t.Errorf("Comment.GetText: not equal text. From method %s. Text field %s", testComment.GetText(), testComment.Text)
		}
		newText := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 255), "abc")
		err := testComment.ChangeText(Text{Text: newText})
		if err != nil {
			t.Errorf("Comment.GetText: text not changed to valid text")
		}
		if testComment.GetText().String() != testComment.Text.String() {
			t.Errorf("Comment.GetText: not equal text. From method %s. Text field %s.", testComment.GetText(), testComment.Text)
		}
	}
}

func TestChangeText(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		emptyText := Text{}
		err := testComment.ChangeText(emptyText)
		if err == nil {
			t.Errorf("Comment.ChangeText: text changed to empty")
		}
		toLongText := Text{Text: helpers.GenerateRandomString(5001, "abc")}
		err = testComment.ChangeText(toLongText)
		if err == nil {
			t.Errorf("Comment.ChangeText: text changed to longer than allowed")
		}
		validText := helpers.GenerateRandomString(helpers.GenerateIntBetween(1, 5000), "abc")
		text := Text{Text: validText}
		err = testComment.ChangeText(text)
		if err != nil || testComment.Text.String() != text.String() {
			t.Errorf("Comment.ChangeText: text not changed to valid text")
		}
	}
}

func TestGetCreateTime(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if testComment.GetCreateTime() != testComment.CreatedAt {
			t.Errorf("Comment.GetCreateTime: not equal create date. From method %s. CreatedAt field %s.", testComment.GetCreateTime(), testComment.CreatedAt)
		}
	}
}

func TestIsDeleted(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if testComment.IsDeleted() {
			t.Errorf("Comment.IsDeleted: expected not deleted comment. Got deleted at %s", testComment.DeletedAt)
		}
		testComment.MarkDeleted()
		if !testComment.IsDeleted() {
			t.Errorf("Comment.IsDeleted: expected deleted comment. Got not deleted")
		}
		if testComment.IsDeleted() != !testComment.DeletedAt.IsZero() {
			t.Errorf("Comment.IsDeleted: status is not defined correctly")
		}
	}
}

func TestMarkDeleted(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if !testComment.DeletedAt.IsZero() {
			t.Errorf("Comment.MarkDeleted: expected not deleted comment. Got deleted at %s", testComment.DeletedAt)
		}
		testComment.MarkDeleted()
		if testComment.DeletedAt.IsZero() {
			t.Errorf("Comment.MarkDeleted: expected deleted comment. Got not deleted")
		}
	}
}

func TestMarkUpdated(t *testing.T) {
	tests := getTestComment()
	for _, testComment := range tests {
		if !testComment.UpdatedAt.IsZero() {
			t.Errorf("Comment.MarkUpdated: expected not updated comment. Got updated at %s", testComment.UpdatedAt)
		}
		testComment.markUpdated()
		if testComment.UpdatedAt.IsZero() {
			t.Errorf("Comment.MarkUpdated: expected updated comment. Got not updated")
		}
	}
}

func TestCommentMarshalJSON(t *testing.T) {
	var tests = getTestComment()
	for _, testComment := range tests {
		notEmptyJson, err := testComment.MarshalJSON()
		if err != nil {
			t.Errorf("Comment.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString := string(notEmptyJson)
		if len(notEmptyJsonString) == 0 || notEmptyJsonString == "{}" || notEmptyJsonString == "null" {
			t.Errorf("Comment.MarshalJSON: expected marshal not empty collection to not empty string, not '{}' and not 'null'. Got %s", notEmptyJsonString)
		}
		if !strings.Contains(notEmptyJsonString, "\"id\"") {
			t.Errorf("Comment.MarshalJSON: test comment has id and json must contain 'id' field but not contain. Got string %s", notEmptyJsonString)
		}
		testComment.markUpdated()
		notEmptyJson, err = testComment.MarshalJSON()
		if err != nil {
			t.Errorf("Comment.MarshalJSON: expected marshal without error. Got with error: %s", err.Error())
		}
		notEmptyJsonString = string(notEmptyJson)
		if !strings.Contains(notEmptyJsonString, "\"updated_at\"") {
			t.Errorf("Comment.MarshalJSON: test comment has update date and json must contain 'updated_at' field but not contain. Got string %s", notEmptyJsonString)
		}
	}
}

func getTestComment() []Comment {
	testsNumber := 5
	var tests []Comment
	for testIndex := 0; testIndex < testsNumber; testIndex++ {
		comment := Comment{ID: uuid.New()}
		tests = append(tests, comment)
	}
	return tests
}
