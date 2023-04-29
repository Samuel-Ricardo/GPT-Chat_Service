package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
  Router        chi.Router
  Handlers      map[string]http.HandlerFunc
  WebServerPort string
}

func NewWebServer(WebServerPort string) *WebServer {
  return &WebServer {
    WebServerPort: WebServerPort,
    Router: chi.NewRouter(),
    Handlers: make(map[string]http.HandlerFunc),
  }
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
  s.Handlers[path] = handler
}

func (s *WebServer) Start() {
  s.Router.Use(middleware.Logger)
}
