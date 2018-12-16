package badger

import (
	"github.com/augmentable-opensource/phraser/pkg/phraser"
	"github.com/dgraph-io/badger"
)

// Backend uses the badger KV store as a backend
type Backend struct {
	DB *badger.DB
}

// NewBadgerBackend returns a new badger backend
func NewBadgerBackend(db *badger.DB) *Backend {
	return &Backend{
		DB: db,
	}
}

// NewBadgerBackendFromDir returns a new badger backend using a directory string to initialize
func NewBadgerBackendFromDir(dir string) (*Backend, error) {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = dir
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Backend{
		DB: db,
	}, nil
}

// Close closes the badger backend
func (backend *Backend) Close() error {
	return backend.DB.Close()
}

// SetPhrase implements SetsPhrase
func (backend *Backend) SetPhrase(collection, path, value string) (*phraser.Phrase, error) {
	err := backend.DB.Update(func(txn *badger.Txn) error {
		key := phraser.BuildPath(collection, path)
		return txn.Set(key, []byte(value))
	})
	if err != nil {
		return nil, err
	}
	return &phraser.Phrase{
		Collection: collection,
		Path:       path,
		Value:      value,
	}, nil
}

// GetPhrase implements GetPhrase
func (backend *Backend) GetPhrase(collection, path string) (*phraser.Phrase, error) {
	key := phraser.BuildPath(collection, path)
	var value []byte
	err := backend.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		value, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &phraser.Phrase{
		Collection: collection,
		Path:       path,
		Value:      string(value),
	}, nil
}

// SetPhraseIfNotExists implements SetPhraseIfNotExists
func (backend *Backend) SetPhraseIfNotExists(collection, path, value string) (*phraser.Phrase, error) {
	err := backend.DB.Update(func(txn *badger.Txn) error {
		key := phraser.BuildPath(collection, path)
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return txn.Set(key, []byte(value))
			}
			return err
		}
		return phraser.ErrPhraseExists
	})
	if err != nil {
		return nil, err
	}
	return &phraser.Phrase{
		Collection: collection,
		Path:       path,
		Value:      value,
	}, nil
}
