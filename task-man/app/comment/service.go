package comment

import (
	"github.com/go-playground/validator/v10"
)

type CommentsService struct {
	Validate           *validator.Validate
	CommentsRepository CommentsRepository
}

func (cs CommentsService) CreateComment(createDTO CreateDTO) (comment Comment, err error) {
	commentsFactory := CommentsFactory{Validate: cs.Validate}
	createAction := CreateCommentAction{
		DTO:                createDTO,
		CommentsFactory:    commentsFactory,
		CommentsRepository: cs.CommentsRepository,
	}
	return createAction.Execute()
}

func (cs CommentsService) EditComment(editDTO EditCommentDTO) (comment Comment, err error) {
	editAction := EditCommentAction{
		DTO:                editDTO,
		CommentsRepository: cs.CommentsRepository,
	}
	return editAction.Execute()
}

func (cs CommentsService) DeleteTaskComments(deleteDTO DeleteTasksCommentsDTO) error {
	deleteTaskCommentsAction := DeleteTasksCommentsAction{
		DTO:                deleteDTO,
		CommentsRepository: cs.CommentsRepository,
	}
	return deleteTaskCommentsAction.Execute()
}
