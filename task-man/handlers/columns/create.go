package columns

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	"github.com/DimaKuptsov/task-man/app/column"
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
	id := r.Form.Get(ProjectIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil || projectID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ProjectIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	name := r.Form.Get(ColumnNameField)
	if name == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", ColumnNameField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	createDTO := column.CreateDTO{
		ProjectID: projectID,
		Name:      name,
	}
	createdColumn, err := app.GetAppService().ColumnsService().CreateColumn(createDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusCreated, createdColumn)
}
