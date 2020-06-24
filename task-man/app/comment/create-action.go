package comment

type CreateCommentAction struct {
	DTO                CreateDTO
	CommentsFactory    CommentsFactory
	CommentsRepository CommentsRepository
}

func (action CreateCommentAction) Execute() (comment Comment, err error) {
	comment, err = action.CommentsFactory.Create(action.DTO)
	if err != nil {
		return
	}
	err = action.CommentsRepository.Save(comment)
	if err != nil {
		return
	}
	return comment, err
}
