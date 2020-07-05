package columns

import (
	"errors"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func GetForProject(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.GetWithContext(r.Context())
	defer appLogger.Sync()
	responseSender := helper.NewResponseSender(appLogger)
	id := chi.URLParam(r, ProjectIDField)
	if id == "" {
		err := errors.New(httpErrors.GetMissingParameterErrorMessage(ProjectIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil || projectID.String() == "" {
		err = errors.New(httpErrors.GetBadParameterErrorMessage(ProjectIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	projects, err := appService.ColumnsService().GetForProject(projectID)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(validationErr))
			return
		}
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	responseSender.SendResponse(w, http.StatusOK, projects)
}
