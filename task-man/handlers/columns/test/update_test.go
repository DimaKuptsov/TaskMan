package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	"github.com/DimaKuptsov/task-man/handlers/columns"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/helpers"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestColumnsUpdateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"name=TestName&priority=2", columns.ColumnIDField},
		{"name=NewName&priority=0", columns.ColumnIDField},
		{"name=NameForUpdate&priority=1", columns.ColumnIDField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/columns/update", requestBody)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Columns.UpdateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Columns.UpdateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestColumnsUpdateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"id=0&name=TestName&priority=2", columns.ColumnIDField},
		{"id=1231&name=NewName&priority=0", columns.ColumnIDField},
		{"id=QWE1201-1231-123-00000&name=NameForUpdate&priority=1", columns.ColumnIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/columns/update", requestBody)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Columns.UpdateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Columns.UpdateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestColumnsUpdateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.UpdateRoute: failed to init logger. Error: %s", err.Error())
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
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/columns/update", requestBody)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Columns.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Columns.UpdateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Columns.UpdateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}

func generateRandomValidRequestBody() string {
	return fmt.Sprintf(
		"id=%s&name=%s&priority=%v",
		uuid.New().String(),
		helpers.GenerateRandomString(50, "name"),
		helpers.GenerateIntBetween(1, 100),
	)
}
