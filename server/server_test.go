package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type StubBacklogStore struct {
	backlog map[string]string
}

func (s *StubBacklogStore) GetBacklog(path string) string {
	trimmedPath := strings.TrimPrefix(path, "/backlog/")
	return s.backlog[trimmedPath]
}

func TestGETBacklog(t *testing.T) {
	store := StubBacklogStore{
		map[string]string{
			"johnnytest": "johnny backlog",
			"joannetest": "joanne backlog",
		},
	}
	server := &KanbanServer{&store}
	t.Run("returns Johnny Test's backlog", func(t *testing.T) {
		request := backlogRequestBuilder("johnnytest")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "johnny backlog"
		assertResponseBody(t, got, want)
	})

	t.Run("returns Joanne's Test's backlog", func(t *testing.T) {
		request := backlogRequestBuilder("joannetest")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "joanne backlog"
		assertResponseBody(t, got, want)
	})
}

func backlogRequestBuilder(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/backlog/"+name, nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
