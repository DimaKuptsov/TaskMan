package columns

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	"github.com/DimaKuptsov/task-man/app/column"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
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
	id := r.Form.Get(ColumnIDField)
	if id == "" {
		err = errors.New(fmt.Sprintf("missing required field \"%s\"", ColumnIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	projectID, err := uuid.Parse(id)
	if err != nil || projectID.String() == "" {
		err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ColumnIDField))
		responseSender.SendErrorResponse(w, httpErrors.NewUnprocessableEntityError(err))
		return
	}
	name := r.Form.Get(ColumnNameField)
	updateDTO := column.UpdateDTO{
		ID:   projectID,
		Name: name,
	}
	priorityStr := r.Form.Get(ColumnPriorityField)
	if name == "" && priorityStr == "" {
		err = errors.New("no parameters for updating")
		responseSender.SendErrorResponse(w, httpErrors.NewBadRequestError(err))
		return
	}
	if priorityStr != "" {
		priority, err := strconv.Atoi(priorityStr)
		if err != nil {
			err = errors.New(fmt.Sprintf("invalid parameter \"%s\"", ColumnPriorityField))
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
	updatedColumn, err := appService.ColumnsService().UpdateColumn(updateDTO)
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
	responseSender.SendResponse(w, http.StatusOK, updatedColumn)
}
