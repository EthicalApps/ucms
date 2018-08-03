package api

import (
	"net/http"

	"github.com/ethicalapps/ucms/cms"
	"github.com/go-chi/chi"
)

func getContentType(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value("repo").(*cms.Repository)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	contentType := chi.URLParam(r, "contentType")
	data, err := repo.GetType(contentType)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getSchema(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value("repo").(*cms.Repository)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	contentType := chi.URLParam(r, "contentType")
	data, err := repo.GetSchema(contentType)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getDocument(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	repo, ok := ctx.Value("repo").(*cms.Repository)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	contentType := chi.URLParam(r, "contentType")
	contentID := chi.URLParam(r, "contentID")
	data, err := repo.GetDocument(contentType, contentID)
	if err != nil {
		http.Error(w, http.StatusText(422), 422)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
