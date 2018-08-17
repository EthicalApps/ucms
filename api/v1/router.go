package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/ethicalapps/ucms/cms"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Router defines the routes for the v1 API
func Router() http.Handler {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/{repository}", func(r chi.Router) {
		r.Use(repositoryCtx)

		r.Get("/type/{contentType}.json", getContentType)
		r.Get("/schema/{contentType}.json", getSchema)
		r.Get("/{contentType}/{contentID}.json", getDocument)
		r.Get("/{contentType}", listDocuments)

		r.Group(func(r chi.Router) {
			r.Use(authentication)
			r.Put("/type/{contentType}.json", putContentType)
			r.Put("/{contentType}/{contentID}.json", putDocument)
		})

		r.NotFound(notFound)
		r.MethodNotAllowed(methodNotAllowed)
	})

	return r
}

type contextKey string

var (
	contextKeyRepo      = contextKey("repo")
	urlParamRepository  = "repository"
	urlParamContentType = "contentType"
	urlParamContentID   = "contentID"
)

func repositoryCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repositoryName := chi.URLParam(r, urlParamRepository)
		repo, err := cms.NewRepository(repositoryName)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyRepo, repo)
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

func notFound(w http.ResponseWriter, r *http.Request) {
	err := errors.New("route not found")
	render.Render(w, r, ErrNotFound(err))
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	err := errors.New("method not allowed")
	render.Render(w, r, ErrNotFound(err))
}
