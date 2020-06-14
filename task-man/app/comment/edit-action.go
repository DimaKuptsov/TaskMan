package comment

import "github.com/go-playground/validator/v10"

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
		return
	}
	err = editedComment.ChangeText(editedText)
	if err != nil {
		return
	}
	err = action.CommentsRepository.Save(editedComment)
	return editedComment, err
}
