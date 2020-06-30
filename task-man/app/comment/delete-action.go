package comment

type DeleteCommentAction struct {
	DTO                DeleteCommentDTO
	CommentsRepository CommentsRepository
}

func (action DeleteCommentAction) Execute() error {
	comment, err := action.CommentsRepository.FindById(action.DTO.CommentID, WithoutDeletedComments)
	if err != nil {
		return nil
	}
	comment.MarkDeleted()
	commentsForUpdate := CommentsCollection{}
	commentsForUpdate.Add(comment)
	return action.CommentsRepository.BatchUpdate(commentsForUpdate)
}
