package bolt

import (
	"fmt"
	"time"

	b "github.com/coreos/bbolt"
	"github.com/ethicalapps/ucms/cms"
)

// Store holds the state for the store
type Store struct {
	db *b.DB
}

// New returns a new instance of a Store
func New(path string) (cms.Store, error) {
	db, err := b.Open(path, 0600, &b.Options{Timeout: 2 * time.Second})
	if err != nil {
		return nil, err
	}
	return &Store{db: db}, nil
}

// CreateBucket creates a new repository bucket if it doesn't already exist
func (s Store) CreateBucket(name string) error {
	return s.db.Update(func(tx *b.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

// PutDocument stores a document in the bucket at the specified key
func (s Store) PutDocument(bucket, namespace, id string, data []byte) error {
	return s.db.Update(func(tx *b.Tx) error {
		root := tx.Bucket([]byte(bucket))
		b, err := root.CreateBucketIfNotExists([]byte(namespace))
		if err != nil {
			return err
		}

		return b.Put([]byte(id), data)
	})
}

// GetDocument returns the document at the specified key in the bucket
func (s Store) GetDocument(bucket, namespace, id string) ([]byte, error) {
	var v []byte
	err := s.db.View(func(tx *b.Tx) error {
		root := tx.Bucket([]byte(bucket))
		b := root.Bucket([]byte(namespace))
		v = b.Get([]byte(id))
		return nil
	})
	return v, err
}

// ListDocuments returns paginated documents of the specified type
func (s Store) ListDocuments(bucket, namespace string) ([][]byte, error) {
	var documents [][]byte
	// var total int

	err := s.db.View(func(tx *b.Tx) error {
		root := tx.Bucket([]byte(bucket))
		b := root.Bucket([]byte(namespace))
		c := b.Cursor()
		n := b.Stats().KeyN
		// total = n

		// return nil if no content
		if n == 0 {
			return nil
		}

		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			documents = append(documents, v)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return documents, nil
}

// Close closes the store
func (s Store) Close() error {
	return s.db.Close()
}
