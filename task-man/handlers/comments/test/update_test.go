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

func TestCommentsUpdateWithMissingParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()

	var testsWithMissingParams = []struct {
		body         string
		missingField string
	}{
		{"text=TestName", comments.CommentIDField},
		{fmt.Sprintf("id=%s", uuid.New()), comments.CommentTextField},
	}
	for _, test := range testsWithMissingParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/comments/update", requestBody)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Comments.UpdateRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetMissingParameterErrorMessage(test.missingField)) {
			t.Errorf("Comments.UpdateRoute: expected missing parameter %s. Got %s", test.missingField, line)
		}
		res.Body.Close()
	}
}

func TestCommentsUpdateWithInvalidParams(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.UpdateRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidParams = []struct {
		body         string
		invalidField string
	}{
		{"id=0&text=Comment", comments.CommentIDField},
		{"id=1231&text=Comment", comments.CommentIDField},
		{"id=QWE1201-1231-123-00000&text=Comment", comments.CommentIDField},
	}
	for _, test := range testsWithInvalidParams {
		requestBody := strings.NewReader(test.body)
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/comments/update", requestBody)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.UnprocessableEntityMessage) {
			t.Errorf("Comments.UpdateRoute: expected unprocessable entity request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(test.invalidField)) {
			t.Errorf("Comments.UpdateRoute: expected bad parameter %s. Got %s", test.invalidField, line)
		}
		res.Body.Close()
	}
}

func TestCommentsUpdateExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.UpdateRoute: failed to init logger. Error: %s", err.Error())
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
		request, err := http.NewRequest(http.MethodPut, srv.URL+"/comments/update", requestBody)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected create request without errors. Got %s", err.Error())
		}
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		client := &http.Client{}
		res, err := client.Do(request)
		if err != nil {
			t.Errorf("Comments.UpdateRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.UpdateRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
			t.Errorf("Comments.UpdateRoute: expected internal server error response. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
			t.Errorf("Comments.UpdateRoute: expected internal server error description. Got %s", line)
		}
		res.Body.Close()
	}
}

func generateRandomValidRequestBody() string {
	return fmt.Sprintf(
		"id=%s&text=%s",
		uuid.New().String(),
		helpers.GenerateRandomString(50, "name"),
	)
}
