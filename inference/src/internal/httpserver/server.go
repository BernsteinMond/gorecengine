package httpserver

import (
	"github.com/BernsteinMond/gorecengine/inference/src/internal/core"
	"github.com/BernsteinMond/gorecengine/inference/src/internal/httpserver/internal/handler"
	"net/http"
)

func New(addr string, service core.Service) *http.Server {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":" + addr,
		Handler: mux,
	}

	handler.SetupRoutes(mux, service)

	return srv
}
