package comment

import (
	"errors"
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
		err = errors.New("invalid task id")
	}
	commentText := Text{createDTO.Text}
	err = f.Validate.Struct(commentText)
	if err != nil {
		return
	}
	comment = Comment{
		id:        uuid.New(),
		taskID:    taskId,
		text:      commentText,
		createdAt: time.Now(),
	}
	return comment, err
}
