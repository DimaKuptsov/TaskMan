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
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	id := r.Form.Get(TaskIDField)
	if id == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	taskID, err := uuid.Parse(id)
	if err != nil || taskID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	name := r.Form.Get(TaskNameField)
	description := r.Form.Get(TaskDescriptionField)
	priorityStr := r.Form.Get(TaskPriorityField)
	if name == "" && description == "" && priorityStr == "" {
		err = errors.New("no parameters for updating")
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	updateDTO := task.UpdateDTO{
		ID:          taskID,
		Name:        name,
		Description: description,
	}
	if priorityStr != "" {
		priority, err := strconv.Atoi(priorityStr)
		if err != nil {
			err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskPriorityField))
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
			return
		}
		updateDTO.Priority = priority
	}
	updatedTask, err := app.GetAppService().TasksService().UpdateTask(updateDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, updatedTask)
}
