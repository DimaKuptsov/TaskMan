package tasks

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/task"
	dbErrors "github.com/DimaKuptsov/task-man/db/errors"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
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
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", TaskIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	taskID, err := uuid.Parse(id)
	if err != nil || taskID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	name := r.Form.Get(TaskNameField)
	description := r.Form.Get(TaskDescriptionField)
	priorityStr := r.Form.Get(TaskPriorityField)
	if name == "" && description == "" && priorityStr == "" {
		err = errors.New("no parameters for updating")
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
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
			responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
			return
		}
		updateDTO.Priority = priority
	}
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	updatedTask, err := appService.TasksService().UpdateTask(updateDTO)
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
	responseSender.SendResponse(w, http.StatusOK, updatedTask)
}
