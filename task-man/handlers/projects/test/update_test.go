package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/projects"
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProjectsUpdateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"name=TestName", projects.ProjectIDField},
		{"name=NewName", projects.ProjectIDField},
		{"id=&name=NameForUpdate", projects.ProjectIDField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/projects/update", requestBody)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Projects.UpdateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Projects.UpdateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestProjectsUpdateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"id=0&name=TestName", projects.ProjectIDField},
		{"id=1231&name=NewName", projects.ProjectIDField},
		{"id=QWE1201-1231-123-00000&name=NameForUpdate", projects.ProjectIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/projects/update", requestBody)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Projects.UpdateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Projects.UpdateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestProjectsUpdateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithValidParams = []struct {
		body string
	}{
		{generateRandomValidRequestBody()},
		{generateRandomValidRequestBody()},
		{generateRandomValidRequestBody()},
	}
	for _, test := range testsWithValidParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/projects/update", requestBody)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Projects.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Projects.UpdateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Projects.UpdateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}

func generateRandomValidRequestBody() string {
	return fmt.Sprintf(
		"id=%s&name=%s",
		uuid.New().String(),
		helpers.GenerateRandomString(50, "name"),
	)
}
