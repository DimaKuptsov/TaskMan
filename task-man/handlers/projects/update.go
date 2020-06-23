package projects

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/project"
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
	id := r.Form.Get(ProjectIDField)
	if id == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil || projectID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	name := r.Form.Get(ProjectNameField)
	description := r.Form.Get(ProjectDescriptionField)
	if name == "" && description == "" {
		err = errors.New("no parameters for updating")
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	updateDTO := project.UpdateDTO{
		ID:          projectID,
		Name:        name,
		Description: description,
	}
	updatedProject, err := app.GetAppService().ProjectsService().UpdateProject(updateDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, updatedProject)
}
