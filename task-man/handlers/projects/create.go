package projects

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/project"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	name := r.Form.Get(ProjectNameField)
	if name == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", ProjectNameField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	description := r.Form.Get(ProjectDescriptionField)
	createDTO := project.CreateDTO{
		Name:        name,
		Description: description,
	}
	createdProject, err := app.GetAppService().ProjectsService().CreateProject(createDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusCreated, createdProject)
}
