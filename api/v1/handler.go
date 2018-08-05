package v1

import (
	"errors"
	"net/http"

	"github.com/ethicalapps/ucms/cms"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func getContentType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value(contextKeyRepo).(*cms.Repository)
	if !ok {
		err := errors.New(string(contextKeyRepo) + " not found in context")
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	contentType := chi.URLParam(r, urlParamContentType)

	data, err := repo.GetType(contentType)
	if err != nil {
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	if data == nil {
		err = errors.New("type " + "'" + contentType + "'" + " not found")
		render.Render(w, r, ErrNotFound(err))
		return
	}

	if err := render.Render(w, r, &JSONResponse{data: data}); err != nil {
		render.Render(w, r, ErrInternalServerError(err))
	}
}

func getSchema(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value(contextKeyRepo).(*cms.Repository)
	if !ok {
		err := errors.New(string(contextKeyRepo) + " not found in context")
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	contentType := chi.URLParam(r, urlParamContentType)

	data, err := repo.GetSchema(contentType)
	if err != nil {
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	if data == nil {
		err = errors.New("schema " + "'" + contentType + "'" + " not found")
		render.Render(w, r, ErrNotFound(err))
		return
	}

	if err := render.Render(w, r, &JSONResponse{data: data}); err != nil {
		render.Render(w, r, ErrInternalServerError(err))
	}
}

func getDocument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value(contextKeyRepo).(*cms.Repository)
	if !ok {
		err := errors.New(string(contextKeyRepo) + " not found in context")
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	contentType := chi.URLParam(r, urlParamContentType)
	contentID := chi.URLParam(r, urlParamContentID)

	data, err := repo.GetDocument(contentType, contentID)
	if err != nil {
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	if data == nil {
		err = errors.New("document " + "'" + contentType + "/" + contentID + "'" + " not found")
		render.Render(w, r, ErrNotFound(err))
		return
	}

	if err := render.Render(w, r, &JSONResponse{data: data}); err != nil {
		render.Render(w, r, ErrInternalServerError(err))
	}
}

func putContentType(w http.ResponseWriter, r *http.Request) {
}

func putDocument(w http.ResponseWriter, r *http.Request) {
}
