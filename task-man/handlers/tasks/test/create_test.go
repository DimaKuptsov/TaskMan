package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/handlers/tasks"
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTasksCreateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"name=TestName", tasks.ColumnIDField},
		{fmt.Sprintf("columnId=%s", uuid.New().String()), tasks.TaskNameField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/tasks/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Tasks.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Tasks.CreateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Tasks.CreateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestTasksCreateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"columnId=1&name=TestName", tasks.ColumnIDField},
		{"columnId=qwe&name=TestName", tasks.ColumnIDField},
		{"columnId=OOQE012-12312qe-123&name=TestName", tasks.ColumnIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/tasks/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Tasks.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Tasks.CreateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Tasks.CreateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestTasksCreateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithValidParams = []struct {
		body string
	}{
		{fmt.Sprintf("columnId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
		{fmt.Sprintf("columnId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
		{fmt.Sprintf("columnId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
	}
	for _, test := range testsWithValidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/tasks/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Tasks.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Tasks.CreateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Tasks.CreateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}
