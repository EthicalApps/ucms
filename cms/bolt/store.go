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

// CreateBucket creates a new bucket if it doesn't already exist
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
func (s Store) PutDocument(bucket, key string, data []byte) error {
	return s.db.Update(func(tx *b.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), data)
		return err
	})
}

// GetDocument returns the document at the specified key in the bucket
func (s Store) GetDocument(bucket, key string) ([]byte, error) {
	var v []byte
	err := s.db.View(func(tx *b.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v = b.Get([]byte(key))
		return nil
	})
	return v, err
}

// Close closes the store
func (s Store) Close() error {
	return s.db.Close()
}
