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
	editedComment, err = action.CommentsRepository.FindById(action.DTO.CommentID, WithoutDeletedComments)
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
	commentsForUpdate := CommentsCollection{}
	commentsForUpdate.Add(editedComment)
	err = action.CommentsRepository.BatchUpdate(commentsForUpdate)
	return editedComment, err
}
