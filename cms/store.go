package cms

// Store is the generic interface to a uCMS store
type Store interface {
	CreateBucket(name string) error
	PutDocument(bucket, namespace, id string, data []byte) error
	GetDocument(bucket, namespace, id string) ([]byte, error)
	ListDocuments(bucket, namespace string) ([][]byte, error)
	Close() error
}

var store Store

// Init initializes the CMS with the given store
func Init(s Store) {
	store = s
}
