package projects

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/project"
	dbErrors "github.com/DimaKuptsov/task-man/db/errors"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/DimaKuptsov/task-man/logger"
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
	name := r.Form.Get(ProjectNameField)
	if name == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", ProjectNameField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	description := r.Form.Get(ProjectDescriptionField)
	createDTO := project.CreateDTO{
		Name:        name,
		Description: description,
	}
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	createdProject, err := appService.ProjectsService().CreateProject(createDTO)
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
	responseSender.SendResponse(w, http.StatusCreated, createdProject)
}
