package main

import (
	"log"
	"net/http"

	"github.com/brettearle/kanbanServer/server"
)

type InMemoryBacklogStore struct{}

func (i *InMemoryBacklogStore) GetBacklog(path string) string {
	return "hello world from get backlog"
}

func main() {
	server := &server.KanbanServer{Store: &InMemoryBacklogStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
