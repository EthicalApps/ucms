package cms

// Store is the generic interface to a uCMS store
type Store interface {
	CreateBucket(name string) error
	PutDocument(bucket, namespace, id string, data []byte) error
	GetDocument(bucket, namespace, id string) ([]byte, error)
	ListDocuments(bucket, namespace string) ([][]byte, error)
	Close() error
}

var (
	dir   string
	store Store
)

// Init initializes the CMS with the given dir & store
func Init(d string, s Store) {
	dir = d
	store = s
}
