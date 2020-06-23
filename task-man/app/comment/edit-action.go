package comment

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
)

type EditCommentAction struct {
	DTO                EditCommentDTO
	CommentsRepository CommentsRepository
}

func (action EditCommentAction) Execute() (editedComment Comment, err error) {
	editedComment, err = action.CommentsRepository.FindById(action.DTO.CommentID)
	if err != nil {
		return
	}
	editedText := Text{action.DTO.EditedText}
	err = validator.New().Struct(editedComment)
	if err != nil {
		return editedComment, appErrors.ValidationError{Field: TextField, Message: err.Error()}
	}
	err = editedComment.ChangeText(editedText)
	if err != nil {
		return
	}
	err = action.CommentsRepository.Save(editedComment)
	return editedComment, err
}
