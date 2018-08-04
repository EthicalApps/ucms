package v1

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrResponse response
type ErrResponse struct {
	Err            error `json:"-"`          // low-level runtime error
	HTTPStatusCode int   `json:"statusCode"` // http response status code

	StatusText string      `json:"status"`           // user-level status message
	AppCode    int64       `json:"code,omitempty"`   // application-specific error code
	ErrorText  string      `json:"error,omitempty"`  // application-level error message, for debugging
	Fields     interface{} `json:"fields,omitempty"` // field validation errors
}

// Render runs before response is written
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrNotFound error
func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     http.StatusText(http.StatusNotFound),
		ErrorText:      err.Error(),
	}
}

// ErrBadRequest error
func ErrBadRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     http.StatusText(http.StatusBadRequest),
		ErrorText:      err.Error(),
	}
}

// ErrInternalServerError error
func ErrInternalServerError(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     http.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}
