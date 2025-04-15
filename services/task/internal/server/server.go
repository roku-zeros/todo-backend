package server

import (
	"net/http"

	"backend-todo-list/services/task/internal/providers"
)

type Server struct {
	taskProvider providers.TaskProvider
}

func New(marketProvider providers.TaskProvider) *Server {
	return &Server{
		taskProvider: marketProvider,
	}
}

func (s *Server) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /v1/ping", s.Ping)
	mux.HandleFunc("POST /v1/tasks", s.CreateTask)
	mux.HandleFunc("GET /v1/tasks", s.ListTasks)
	mux.HandleFunc("GET /v1/tasks/", s.GetTask)
	mux.HandleFunc("PUT /v1/tasks/", s.UpdateTask)
	mux.HandleFunc("DELETE /v1/tasks/", s.DeleteTask)
	mux.HandleFunc("GET /v1/tasks/date", s.ListTasksByDate)
}
