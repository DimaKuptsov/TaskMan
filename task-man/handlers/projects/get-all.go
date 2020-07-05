package projects

import (
	"github.com/DimaKuptsov/task-man/app"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/helper"
	"github.com/DimaKuptsov/task-man/logger"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.GetWithContext(r.Context())
	defer appLogger.Sync()
	responseSender := helper.NewResponseSender(appLogger)
	appService, err := app.GetAppService()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	projects, err := appService.ProjectsService().GetAll()
	if err != nil {
		appLogger.Error(err.Error())
		responseSender.SendErrorResponse(w, httpErrors.NewInternalServerError(err))
		return
	}
	responseSender.SendResponse(w, http.StatusOK, projects)
}
