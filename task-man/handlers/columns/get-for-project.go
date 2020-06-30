package columns

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func GetForProject(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, ProjectIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil || projectID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projects, err := app.GetAppService().ColumnsService().GetForProject(projectID)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewBadRequestError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, projects)
}
