package test

import (
	"bufio"
	"fmt"
	"github.com/DimaKuptsov/task-man/handlers"
	"github.com/DimaKuptsov/task-man/handlers/comments"
	httpErrors "github.com/DimaKuptsov/task-man/handlers/error"
	"github.com/DimaKuptsov/task-man/logger"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommentsGetForTaskWithInvalidID(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.GetForTaskRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	var testsWithInvalidIds = []struct {
		invalidId string
	}{
		{"0"},
		{"123"},
		{"QIEw-12312qw-sad123"},
	}
	for _, test := range testsWithInvalidIds {
		getUrl := fmt.Sprintf("%s/comments/task/%s", srv.URL, test.invalidId)
		res, err := http.Get(getUrl)
		if err != nil {
			t.Errorf("Comments.GetForTaskRoute: expected send request without errors. Got %s", err.Error())
		}
		reader := bufio.NewReader(res.Body)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Errorf("Comments.GetForTaskRoute: expected response reading without errors. Got %s", err.Error())
		}
		if !strings.Contains(line, httpErrors.BadRequestMessage) {
			t.Errorf("Comments.GetForTaskRoute: expected bad request. Got %s", line)
		}
		if !strings.Contains(line, httpErrors.GetBadParameterErrorMessage(comments.TaskIDField)) {
			t.Errorf("Comments.GetForTaskRoute: expected bad parameter %s. Got %s", comments.TaskIDField, line)
		}
		res.Body.Close()
	}
}

func TestColumnsGetForTaskExpectedInternalServerError(t *testing.T) {
	err := logger.Init()
	if err != nil {
		t.Errorf("Comments.GetForTaskRoute: failed to init logger. Error: %s", err.Error())
	}
	router := handlers.NewRouter()
	srv := httptest.NewServer(router)
	defer srv.Close()
	getUrl := fmt.Sprintf("%s/comments/task/%s", srv.URL, uuid.New().String())
	res, err := http.Get(getUrl)
	if err != nil {
		t.Errorf("Comments.GetForTaskRoute: expected send request without errors. Got %s", err.Error())
	}
	reader := bufio.NewReader(res.Body)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		t.Errorf("Comments.GetForTaskRoute: expected response reading without errors. Got %s", err.Error())
	}
	if !strings.Contains(line, httpErrors.InternalServerErrorMessage) {
		t.Errorf("Comments.GetForTaskRoute: expected internal server error response. Got %s", line)
	}
	if !strings.Contains(line, httpErrors.InternalServerErrorDescription) {
		t.Errorf("Comments.GetForTaskRoute: expected internal server error description. Got %s", line)
	}
	res.Body.Close()
}
