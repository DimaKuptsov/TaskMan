package comments

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	"github.com/DimaKuptsov/task-man/app/comment"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	dbErrors "github.com/DimaKuptsov/task-man/db/errors"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.GetWithContext(r.Context())
	defer appLogger.Sync()
	responseSender := helper.NewResponseSender(appLogger)
	err := r.ParseForm()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	id := r.Form.Get(TaskIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", TaskIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	taskID, err := uuid.Parse(id)
	if err != nil || taskID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	text := r.Form.Get(CommentTextField)
	if text == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", CommentTextField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	createDTO := comment.CreateDTO{
		TaskID: taskID,
		Text:   text,
	}
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	createdComment, err := appService.CommentsService().CreateComment(createDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		if notFoundErr, ok := err.(dbErrors.NoRowsFoundError); ok {
			responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(notFoundErr))
			return
		}
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	responseSender.SendResponse(w, http.StatusCreated, createdComment)
}
