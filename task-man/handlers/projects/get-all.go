package projects

import (
	"github.com/DimaKuptsov/task-man/app"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	projects, err := app.GetAppService().ProjectsService().GetAll()
	if err != nil {
		helper.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	helper.SendResponse(w, http.StatusOK, projects)
}
