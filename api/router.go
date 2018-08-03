package api

import (
	"context"
	"net/http"
	"time"

	"github.com/ethicalapps/ucms/cms"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Router returns a http.Handler that contains all the API routes
func Router() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/healthz"))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.DefaultCompress)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// r.Use(middleware.WithValue("repo", ))

	r.Route("/{repository}", func(r chi.Router) {
		r.Use(repositoryCtx)
		r.Get("/type/{contentType}.json", getContentType)
		r.Get("/schema/{contentType}.json", getSchema)
		r.Get("/{contentType}/{contentID}.json", getDocument)
	})

	r.Group(func(r chi.Router) {
		r.Use(middleware.Timeout(30 * time.Second))
	})

	r.Group(func(r chi.Router) {
		r.Use(authentication)
	})

	return r
}

func repositoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repositoryName := chi.URLParam(r, "repository")
		repo, err := cms.NewRepository(repositoryName)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "repo", repo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// FIXME: implement authentication
		if true {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		}
	})
}
