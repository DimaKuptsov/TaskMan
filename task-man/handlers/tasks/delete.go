package tasks

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	"github.com/DimaKuptsov/task-man/app/task"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, TaskIDField)
	if id == "" {
		err := errors.New(fmt.Sprintf("missing required field \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	taskID, err := uuid.Parse(id)
	if err != nil || taskID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", TaskIDField))
		helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	deleteDTO := task.DeleteTasksDTO{}
	deleteDTO.AddTaskId(taskID)
	err = app.GetAppService().TasksService().DeleteTasks(deleteDTO)
	if err != nil {
		if validationErr, ok := err.(appErrors.ValidationError); ok {
			helper.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(validationErr))
			return
		}
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, struct{}{})
}
