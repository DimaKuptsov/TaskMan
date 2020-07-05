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

func TestColumnsCreateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"name=TestName", columns.ProjectIDField},
		{fmt.Sprintf("projectId=%s", uuid.New().String()), columns.ColumnNameField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/columns/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Columns.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Columns.CreateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Columns.CreateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestColumnsCreateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"projectId=1&name=TestName", columns.ProjectIDField},
		{"projectId=qwe&name=TestName", columns.ProjectIDField},
		{"projectId=OOQE012-12312qe-123&name=TestName", columns.ProjectIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/columns/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Columns.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Columns.CreateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Columns.CreateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestColumnsCreateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Columns.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithValidParams = []struct {
		body string
	}{
		{fmt.Sprintf("projectId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
		{fmt.Sprintf("projectId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
		{fmt.Sprintf("projectId=%s&name=%s", uuid.New().String(), helpers.GenerateRandomString(50, "name"))},
	}
	for _, test := range testsWithValidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/columns/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Columns.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Columns.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Columns.CreateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Columns.CreateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}
