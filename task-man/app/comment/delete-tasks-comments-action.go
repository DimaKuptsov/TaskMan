package comment

type DeleteTasksCommentsAction struct {
	DTO                DeleteTasksCommentsDTO
	CommentsRepository CommentsRepository
}

func (action DeleteTasksCommentsAction) Execute() error {
	taskComments, err := action.CommentsRepository.FindForTasks(action.DTO.TasksIDs, WithoutDeletedComments)
	if err != nil {
		return nil
	}
	commentsForUpdate := CommentsCollection{}
	for _, comment := range taskComments.Comments {
		err = comment.MarkDeleted()
		if err != nil {
			//TODO log
			continue
		}
		commentsForUpdate.Add(comment)
	}
	return action.CommentsRepository.BatchUpdate(commentsForUpdate)
}
