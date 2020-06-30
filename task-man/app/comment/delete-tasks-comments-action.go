package comment

type DeleteTasksCommentsAction struct {
	DTO                DeleteTasksCommentsDTO
	CommentsRepository CommentsRepository
}

func (action DeleteTasksCommentsAction) Execute() error {
	taskComments, err := action.CommentsRepository.FindForTasks(action.DTO.TasksIDs, WithoutDeletedComments)
	if err != nil {
		return err
	}
	commentsForUpdate := CommentsCollection{}
	for _, comment := range taskComments.Comments {
		comment.MarkDeleted()
		commentsForUpdate.Add(comment)
	}
	return action.CommentsRepository.BatchUpdate(commentsForUpdate)
}
