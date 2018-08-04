package v1

import (
	"encoding/json"
	"net/http"
)

// JSONResponse is the response payload
type JSONResponse struct {
	data     []byte
	Document *map[string]interface{} `json:"document"`
}

// Render is called before the response is written
func (res *JSONResponse) Render(w http.ResponseWriter, r *http.Request) error {
	var doc map[string]interface{}
	if err := json.Unmarshal(res.data, &doc); err != nil {
		return err
	}
	res.Document = &doc
	return nil
}
