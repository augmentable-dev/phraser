package phraser

import (
	"github.com/dgraph-io/badger"
)

// Phraser is a key-value store for your application's phrases
type Phraser struct {
	DB *badger.DB
}

// Phrase is a phrase
type Phrase struct {
	Collection string
	Path       string
	Value      string
}

// NewPhraser creates a new phraser
func NewPhraser(db *badger.DB) *Phraser {
	return &Phraser{DB: db}
}

// Close closes a phraser
func (phraser *Phraser) Close() error {
	return phraser.DB.Close()
}

func buildPath(collection, path string) []byte {
	return []byte(collection + ":" + path)
}

// SetPhrase sets a phrase, no matter if it exists or not
func (phraser *Phraser) SetPhrase(collection, path, value string) (*Phrase, error) {
	err := phraser.DB.Update(func(txn *badger.Txn) error {
		key := buildPath(collection, path)
		return txn.Set(key, []byte(value))
	})
	if err != nil {
		return nil, err
	}
	return &Phrase{
		Collection: collection,
		Path:       path,
		Value:      value,
	}, nil
}

// SetPhraseIfNotExists only sets a phrase if it does not exist. Otherwise will return an error
func (phraser *Phraser) SetPhraseIfNotExists(collection, path, value string) (*Phrase, error) {
	err := phraser.DB.Update(func(txn *badger.Txn) error {
		key := buildPath(collection, path)
		_, err := txn.Get(key)
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return txn.Set(key, []byte(value))
			}
			return err
		}
		return ErrPhraseExists
	})
	if err != nil {
		return nil, err
	}
	return &Phrase{
		Collection: collection,
		Path:       path,
		Value:      value,
	}, nil
}

// GetPhrase gets a phrase
func (phraser *Phraser) GetPhrase(collection, path string) (*Phrase, error) {
	key := buildPath(collection, path)
	var value []byte
	err := phraser.DB.View(func(txn *badger.Txn) error {
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

	return &Phrase{
		Collection: collection,
		Path:       path,
		Value:      string(value),
	}, nil
}

// ListPhrases lists phrases
func (phraser *Phraser) ListPhrases(collection, path string) (*Phrase, error) {

	return nil, nil
}
