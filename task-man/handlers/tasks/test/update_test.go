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

func TestTasksUpdateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"name=TestName&priority=2", tasks.TaskIDField},
		{"name=NewName&priority=0", tasks.TaskIDField},
		{"name=NameForUpdate&priority=1", tasks.TaskIDField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/tasks/update", requestBody)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Tasks.UpdateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Tasks.UpdateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestTasksUpdateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"id=0&name=TestName&description=TestDesc&priority=2", tasks.TaskIDField},
		{"id=1231&name=NewName&description=NewDesc&priority=0", tasks.TaskIDField},
		{"id=QWE1201-1231-123-00000&name=NameForUpdate&description=DescForUpdate&priority=1", tasks.TaskIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/tasks/update", requestBody)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Tasks.UpdateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Tasks.UpdateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestTasksUpdateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Tasks.UpdateRoute: failed to init logger. Error: %s", err.Error())
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
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/tasks/update", requestBody)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Tasks.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Tasks.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Tasks.UpdateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Tasks.UpdateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}

func generateRandomValidRequestBody() string {
	return fmt.Sprintf(
		"id=%s&name=%s&description=%s&priority=%v",
		uuid.New().String(),
		helpers.GenerateRandomString(50, "name"),
		helpers.GenerateRandomString(150, "description"),
		helpers.GenerateIntBetween(1, 100),
	)
}
