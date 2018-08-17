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

// CollectionResponse returns a collection of documents
type CollectionResponse struct {
	data      [][]byte
	Documents []map[string]interface{} `json:"documents"`
}

// Render is called before the response is written
func (res *CollectionResponse) Render(w http.ResponseWriter, r *http.Request) error {
	for _, data := range res.data {
		var doc map[string]interface{}
		if err := json.Unmarshal(data, &doc); err != nil {
			return err
		}

		res.Documents = append(res.Documents, doc)
	}

	return nil
}
