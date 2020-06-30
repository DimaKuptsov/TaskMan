package tasks

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/task"
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
	id := r.Form.Get(ColumnIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", ColumnIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	columnID, err := uuid.Parse(id)
	if err != nil || columnID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ColumnIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	name := r.Form.Get(TaskNameField)
	if name == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", TaskNameField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	description := r.Form.Get(TaskDescriptionField)
	createDTO := task.CreateTaskDTO{
		ColumnID:    columnID,
		Name:        name,
		Description: description,
	}
	createdTask, err := app.GetAppService().TasksService().CreateTask(createDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusCreated, createdTask)
}
