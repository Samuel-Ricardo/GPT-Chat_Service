package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServer struct {
  Router        chi.Router
  Handlers      map[string]http.HandlerFunc
  WebServerProt string
}


