package comment

import "github.com/google/uuid"

type CreateDTO struct {
	TaskID uuid.UUID
	Text   string
}

type EditCommentDTO struct {
	CommentID  uuid.UUID
	EditedText string
}

type DeleteTasksCommentsDTO struct {
	TasksIDs []uuid.UUID
}
