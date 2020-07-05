package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProjectsGetAllExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.GetAllRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	getUrl := fmt.Sprintf("%s/projects/all", srv.URL)
	res, err := http.Get(getUrl)
	if err != nil {
		t.Errorf("Projects.GetAllRoute: expected send request without errors. Got %s", err.Error())
	}
	reader := bufio.NewReader(res.Body)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		t.Errorf("Projects.GetAllRoute: expected response reading without errors. Got %s", err.Error())
	}
	if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
		t.Errorf("Projects.GetAllRoute: expected internal server error response. Got %s", line)
	}
	if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
		t.Errorf("Projects.GetAllRoute: expected internal server error description. Got %s", line)
	}
	res.Body.Close()
}
