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
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.GetWithContext(r.Context())
	defer appLogger.Sync()
	responseSender := helper.NewResponseSender(appLogger)
	id := chi.URLParam(r, CommentIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", CommentIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	commentID, err := uuid.Parse(id)
	if err != nil || commentID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", CommentIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	deleteDTO := comment.DeleteCommentDTO{
		CommentID: commentID,
	}
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	err = appService.CommentsService().DeleteComment(deleteDTO)
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
	responseSender.SendResponse(w, http.StatusOK, struct{}{})
}
