package comment

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type CommentsFactory struct {
	Validate *validator.Validate
}

func (f CommentsFactory) Create(createDTO CreateDTO) (comment Comment, err error) {
	taskId := createDTO.TaskID
	if taskId.String() == "" {
		err = appErrors.ValidationError{Field: TaskIDField, Message: "task id should be in the uuid format"}
	}
	commentText := Text{createDTO.Text}
	err = f.Validate.Struct(commentText)
	if err != nil {
		return comment, appErrors.ValidationError{Field: TextField, Message: err.Error()}
	}
	comment = Comment{
		ID:        uuid.New(),
		TaskID:    taskId,
		Text:      commentText,
		CreatedAt: time.Now(),
	}
	return comment, err
}
