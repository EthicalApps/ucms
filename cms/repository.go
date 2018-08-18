package cms

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// Repository represents a single repository
type Repository struct {
	Name  string
	Store Store
}

const schemaNamespace = "_schema"
const contentTypeNamespace = "_content_type"

// NewRepository returns a Repository for the given store/bucket
func NewRepository(name string) (*Repository, error) {
	if err := store.CreateBucket(name); err != nil {
		return nil, err
	}

	repo := &Repository{name, store}

	if err := repo.PutSchema("type", []byte(ContentTypeSchema)); err != nil {
		return nil, err
	}

	return repo, nil
}

// PutType stores a ContentType in the repository
func (r *Repository) PutType(id string, content []byte) error {
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		return err
	}

	schema, err := r.GetSchema("type")
	if err != nil {
		return err
	}

	typeSchema := gojsonschema.NewStringLoader(string(schema))
	typeDocument := gojsonschema.NewGoLoader(result)

	res, err := gojsonschema.Validate(typeSchema, typeDocument)
	if err != nil {
		return err
	}

	if res.Valid() {
		// save a version of the document
		// update the JSON schema for this type
		contentType := &ContentType{}
		if err := json.Unmarshal(content, &contentType); err != nil {
			return err
		}

		schema, err := GenerateSchema(contentType)
		if err != nil {
			return err
		}

		if err := r.CreateIndex(id); err != nil {
			return err
		}

		if err := r.PutSchema(id, schema); err != nil {
			return err
		}

		return r.Store.PutDocument(r.Name, contentTypeNamespace, id, content)
	}

	fmt.Printf("The document is not valid. see errors :\n")
	for _, desc := range res.Errors() {
		fmt.Printf("- %s\n", desc)
	}
	return errors.New("Type document didn't validate")
}

// GetType returns a type from the repository
func (r *Repository) GetType(id string) ([]byte, error) {
	return r.Store.GetDocument(r.Name, contentTypeNamespace, id)
}

// PutSchema stores a schema in the repository
func (r *Repository) PutSchema(id string, content []byte) error {
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		return err
	}

	// load the schema to check it is valid
	loader := gojsonschema.NewStringLoader(string(content))
	_, err := gojsonschema.NewSchema(loader)
	if err != nil {
		return errors.New("[PutSchema] " + "'" + id + "'" + " is not valid: " + err.Error() + "\n" + string(content))
	}

	return r.Store.PutDocument(r.Name, schemaNamespace, id, content)
}

// GetSchema gets a schema from the Repository
func (r *Repository) GetSchema(id string) ([]byte, error) {
	return r.Store.GetDocument(r.Name, schemaNamespace, id)
}

// PutDocument stores a document in the Repository
func (r *Repository) PutDocument(typeID, id string, content []byte) error {
	var result map[string]interface{}
	if err := json.Unmarshal(content, &result); err != nil {
		return err
	}

	schema, err := r.GetSchema(typeID)
	if err != nil {
		return err
	}
	if schema == nil {
		return errors.New("[PutDocument] couldn't save " + "'" + id + "' : " + "'" + typeID + "'" + " schema not found")
	}

	typeSchema := gojsonschema.NewStringLoader(string(schema))
	typeDocument := gojsonschema.NewGoLoader(result)

	res, err := gojsonschema.Validate(typeSchema, typeDocument)
	if err != nil {
		return err
	}

	if res.Valid() {
		// save a version of the document
		if err := r.IndexDocument(typeID, id, content); err != nil {
			return err
		}

		return r.Store.PutDocument(r.Name, typeID, id, content)
	}

	fmt.Printf("The document is not valid. see errors :\n")
	for _, desc := range res.Errors() {
		fmt.Printf("- %s\n", desc)
	}
	return errors.New("Document document didn't validate")
}

// GetDocument gets a document from the Repository
func (r *Repository) GetDocument(typeID, id string) ([]byte, error) {
	return r.Store.GetDocument(r.Name, typeID, id)
}

// ListDocuments returns a collection of documents of the specified type
func (r *Repository) ListDocuments(typeID string) ([][]byte, error) {
	return r.Store.ListDocuments(r.Name, typeID)
}
