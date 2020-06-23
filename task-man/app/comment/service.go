package comment

import (
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CommentsService struct {
	Validate           *validator.Validate
	CommentsRepository CommentsRepository
}

func (cs CommentsService) GetById(commentID uuid.UUID) (comment Comment, err error) {
	if commentID.String() == "" {
		return comment, appErrors.ValidationError{Field: IDField, Message: "comment id should be in the uuid format"}
	}
	return cs.CommentsRepository.FindById(commentID)
}

func (cs CommentsService) GetForTask(taskID uuid.UUID) (comments CommentsCollection, err error) {
	if taskID.String() == "" {
		return comments, appErrors.ValidationError{Field: TaskIDField, Message: "task id should be in the uuid format"}
	}
	tasksIDs := []uuid.UUID{taskID}
	comments, err = cs.CommentsRepository.FindForTasks(tasksIDs, WithoutDeletedComments)
	if err != nil {
		return comments, err
	}
	comments.SortByCreateTime()
	return comments, err
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

func (cs CommentsService) DeleteComment(deleteDTO DeleteCommentDTO) error {
	deleteAction := DeleteCommentAction{
		DTO:                deleteDTO,
		CommentsRepository: cs.CommentsRepository,
	}
	return deleteAction.Execute()
}

func (cs CommentsService) DeleteTaskComments(deleteDTO DeleteTasksCommentsDTO) error {
	deleteTaskCommentsAction := DeleteTasksCommentsAction{
		DTO:                deleteDTO,
		CommentsRepository: cs.CommentsRepository,
	}
	return deleteTaskCommentsAction.Execute()
}
