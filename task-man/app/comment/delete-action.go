package comment

type DeleteCommentAction struct {
	DTO                DeleteCommentDTO
	CommentsRepository CommentsRepository
}

func (action DeleteCommentAction) Execute() error {
	comment, err := action.CommentsRepository.FindById(action.DTO.CommentID)
	if err != nil {
		return nil
	}
	err = comment.MarkDeleted()
	if err != nil {
		return err
	}
	commentsForUpdate := CommentsCollection{}
	commentsForUpdate.Add(comment)
	return action.CommentsRepository.BatchUpdate(commentsForUpdate)
}
