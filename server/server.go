package server

import (
	"net/http"
)

type BacklogStore interface {
	GetBacklog(path string) string
}

type KanbanServer struct {
	Store BacklogStore
}

func (k *KanbanServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backlog := k.Store.GetBacklog(r.URL.Path)
	w.Write([]byte(backlog))
}
