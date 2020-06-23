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

func Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	id := r.Form.Get(TaskIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	taskID, err := uuid.Parse(id)
	if err != nil || taskID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	text := r.Form.Get(CommentTextField)
	if text == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", CommentTextField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	createDTO := comment.CreateDTO{
		TaskID: taskID,
		Text:   text,
	}
	createdComment, err := app.GetAppService().CommentsService().CreateComment(createDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusCreated, createdComment)
}
