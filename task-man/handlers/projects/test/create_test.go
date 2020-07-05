package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/projects"
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/DimaKuptsov/task-man/logger"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProjectsCreateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"description=project description", projects.ProjectNameField},
		{"description=TestDescription", projects.ProjectNameField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/projects/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Projects.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Projects.CreateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Projects.CreateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestProjectsCreateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{fmt.Sprintf("name=%s&description=%s", "", "test descr"), projects.ProjectNameField},
		{fmt.Sprintf("name=%s&description=%s", "", "other descr"), projects.ProjectNameField},
		{fmt.Sprintf("name=%s&description=%s", "", ""), projects.ProjectNameField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/projects/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Projects.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Projects.CreateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.invalidField)) {
			t.Errorf("Projects.CreateRoute: expected missing parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestProjectsCreateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Projects.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithValidParams = []struct {
		body string
	}{
		{fmt.Sprintf("name=%s&description=%s", helpers.GenerateRandomString(1, "a"), helpers.GenerateRandomString(50, "desc"))},
		{fmt.Sprintf("name=%s&description=%s", helpers.GenerateRandomString(500, "name"), helpers.GenerateRandomString(1000, "desc"))},
		{fmt.Sprintf("name=%s&description=%s", helpers.GenerateRandomString(100, "name"), "")},
	}
	for _, test := range testsWithValidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/projects/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Projects.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Projects.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Projects.CreateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Projects.CreateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}
