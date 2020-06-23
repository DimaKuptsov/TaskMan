package comments

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	"github.com/DimaKuptsov/task-man/app/comment"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/google/uuid"
	"net/http"
)

func Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	id := r.Form.Get(CommentIDField)
	if id == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", CommentIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	commentID, err := uuid.Parse(id)
	if err != nil || commentID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", CommentIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	text := r.Form.Get(CommentTextField)
	if text == "" {
		err = errors.New("no parameters for updating")
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	editDTO := comment.EditCommentDTO{
		CommentID:  commentID,
		EditedText: text,
	}
	editedComment, err := app.GetAppService().CommentsService().EditComment(editDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, editedComment)
}
