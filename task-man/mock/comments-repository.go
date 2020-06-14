package mock

import (
	"github.com/DimaKuptsov/task-man/app/comment"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CommentsRepositoryMock struct {
	comments map[string]comment.Comment
}

func CommentsRepoMock(tasksRepositoryMock TasksRepositoryMock) CommentsRepositoryMock {
	tasks := tasksRepositoryMock.All()
	factory := comment.CommentsFactory{Validate: validator.New()}
	fakeComments := make(map[string]comment.Comment)
	for _, existTask := range tasks {
		firsComment, _ := factory.Create(comment.CreateDTO{TaskID: existTask.GetID(), Text: "First comment for " + existTask.GetID().String()})
		secondComment, _ := factory.Create(comment.CreateDTO{TaskID: existTask.GetID(), Text: "Second comment for " + existTask.GetID().String()})
		thirdComment, _ := factory.Create(comment.CreateDTO{TaskID: existTask.GetID(), Text: "Third comment for " + existTask.GetID().String()})
		fakeComments[firsComment.GetID().String()] = firsComment
		fakeComments[secondComment.GetID().String()] = secondComment
		fakeComments[thirdComment.GetID().String()] = thirdComment
	}
	return CommentsRepositoryMock{fakeComments}
}

func (r CommentsRepositoryMock) FindById(id uuid.UUID) (faceComment comment.Comment, err error) {
	for commentId, existComment := range r.comments {
		if commentId == id.String() {
			faceComment = existComment
			break
		}
	}
	return faceComment, err
}

func (r CommentsRepositoryMock) FindNotDeleted() (comments comment.CommentsCollection, err error) {
	for _, existComment := range r.comments {
		if !existComment.IsDeleted() {
			comments.Add(existComment)
		}
	}
	return comments, err
}

func (r CommentsRepositoryMock) FindForTasks(tasksIDs []uuid.UUID, withDeleted bool) (comments comment.CommentsCollection, err error) {
	for _, taskId := range tasksIDs {
		for _, existComment := range r.comments {
			if existComment.IsDeleted() && withDeleted == comment.WithoutDeletedComments {
				continue
			}
			if existComment.GetTaskID() == taskId {
				comments.Add(existComment)
			}
		}
	}
	return comments, nil
}

func (r CommentsRepositoryMock) Save(comment comment.Comment) error {
	r.comments[comment.GetID().String()] = comment
	return nil
}

func (r CommentsRepositoryMock) BatchUpdate(comments comment.CommentsCollection) error {
	for _, commentForUpdate := range comments.Comments {
		existComment, _ := r.FindById(commentForUpdate.GetID())
		_ = existComment.ChangeText(commentForUpdate.GetText())
		if commentForUpdate.IsDeleted() {
			_ = existComment.MarkDeleted()
		}
		r.comments[existComment.GetID().String()] = existComment
	}
	return nil
}
