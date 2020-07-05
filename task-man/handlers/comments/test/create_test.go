package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	"github.com/DimaKuptsov/task-man/handlers/comments"
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

func TestCommentsCreateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"text=TestComment", comments.TaskIDField},
		{fmt.Sprintf("taskId=%s", uuid.New().String()), comments.CommentTextField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/comments/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Comments.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Comments.CreateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Comments.CreateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestCommentsCreateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"taskId=1&text=TestText", comments.TaskIDField},
		{"taskId=qwe&text=SomeText", comments.TaskIDField},
		{"taskId=OOQE012-12312qe-123&text=OtherText", comments.TaskIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/comments/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Comments.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Comments.CreateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Comments.CreateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestCommentsCreateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.CreateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithValidParams = []struct {
		body string
	}{
		{fmt.Sprintf("taskId=%s&text=%s", uuid.New().String(), helpers.GenerateRandomString(50, "text"))},
		{fmt.Sprintf("taskId=%s&text=%s", uuid.New().String(), helpers.GenerateRandomString(50, "text"))},
		{fmt.Sprintf("taskId=%s&text=%s", uuid.New().String(), helpers.GenerateRandomString(50, "text"))},
	}
	for _, test := range testsWithValidParams {
		requestBody := strings.NewReader(test.body)
		res, err := http.Post(srv.URL+"/comments/create", "application/x-www-form-urlencoded", requestBody)
		if err != nil {
			t.Errorf("Comments.CreateRoute: expected request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.CreateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Comments.CreateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Comments.CreateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}
