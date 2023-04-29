package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

