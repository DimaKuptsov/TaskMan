package columns

import (
	"errors"
	"fmt"
	"github.com/DimaKuptsov/task-man/app"
	"github.com/DimaKuptsov/task-man/app/column"
	appErrors "github.com/DimaKuptsov/task-man/app/error"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, ColumnIDField)
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
	deleteDTO := column.DeleteDTO{
		ID: columnID,
	}
	err = app.GetAppService().ColumnsService().DeleteColumn(deleteDTO)
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
