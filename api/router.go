package api

import (
	"net/http"
	"time"

	"github.com/ethicalapps/ucms/api/v1"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Router returns a http.Handler that contains all the API routes
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultCompress)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Mount("/v1", v1.Router())

	return r
}
