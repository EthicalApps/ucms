package cms

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/blevesearch/bleve"
)

var indexes map[string]bleve.Index

func init() {
	indexes = make(map[string]bleve.Index)
}

// CreateIndex is an experiment with indexing
func (r *Repository) CreateIndex(namespace string) error {
	mapping := bleve.NewIndexMapping()
	// mapping.UnmarshalJSON(GenerateIndexMapping(...))
	filename := r.Name + "_" + namespace + ".bleve"
	filename = filepath.Join(dir, filename)

	var index bleve.Index

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		index, err = bleve.New(filename, mapping)
		if err != nil {
			return err
		}
		index.SetName(namespace)
	} else {
		index, err = bleve.Open(filename)
		if err != nil {
			return err
		}
	}

	indexes[namespace] = index
	return nil
}

// IndexDocument indexes the document in the correct index with the given id
func (r *Repository) IndexDocument(namespace, id string, document []byte) error {
	index, ok := indexes[namespace]
	if !ok {
		return errors.New("attempt to index document in non-existant index")
	}
	return index.Index(id, string(document))
}

// Query performs a search on the given content type
func (r *Repository) Query(namespace, query string, count, offset int) ([]string, error) {
	index, ok := indexes[namespace]
	if !ok {
		return nil, errors.New("attempt to query documents in non-existant index")
	}
	q := bleve.NewQueryStringQuery(query)
	req := bleve.NewSearchRequestOptions(q, count, offset, false)
	res, err := index.Search(req)
	if err != nil {
		return nil, err
	}
	var results []string
	for _, hit := range res.Hits {
		results = append(results, hit.ID)
	}
	return results, nil
}
