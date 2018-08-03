package cms

// Store is the generic interface to a uCMS store
type Store interface {
	CreateBucket(name string) error
	PutDocument(bucket, key string, data []byte) error
	GetDocument(bucket, key string) ([]byte, error)
	Close() error
}

var store Store

// Init initializes the CMS with the given store
func Init(s Store) {
	store = s
}
