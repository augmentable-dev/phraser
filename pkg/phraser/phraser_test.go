package phraser_test

import (
	"log"
	"os"
	"testing"

	"github.com/augmentable-opensource/phraser/pkg/phraser"
	badger_backend "github.com/augmentable-opensource/phraser/pkg/phraser/backends/badger"
)

var (
	p        *phraser.Phraser
	backends map[string]phraser.Backend
)

func setupBadgerBackend() {
	dir := "/tmp/phraser-badger-backend-test"
	backend, err := badger_backend.NewBadgerBackendFromDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	backends["Badger"] = backend
}

func TestMain(m *testing.M) {
	backends = make(map[string]phraser.Backend)
	setupBadgerBackend()
	os.Exit(m.Run())
}

func TestBasic(t *testing.T) {
	for name, backend := range backends {
		t.Run(name, func(t *testing.T) {
			collection := "test-collection"
			path := "test/path/hello"
			value := "world"
			phrase, err := backend.SetPhrase(collection, path, value)
			if err != nil {
				t.Fatal(err)
			}
			if phrase.Collection != collection {
				t.Fatalf("expected %s got %s", collection, phrase.Collection)
			}
			if phrase.Path != path {
				t.Fatalf("expected %s got %s", path, phrase.Path)
			}
			if phrase.Value != value {
				t.Fatalf("expected %s got %s", value, phrase.Value)
			}

			phrase, err = backend.GetPhrase(collection, path)
			if err != nil {
				t.Fatal(err)
			}
			if phrase.Collection != collection {
				t.Fatalf("expected %s got %s", collection, phrase.Collection)
			}
			if phrase.Path != path {
				t.Fatalf("expected %s got %s", path, phrase.Path)
			}
			if phrase.Value != value {
				t.Fatalf("expected %s got %s", value, phrase.Value)
			}

			newValue := "world!"
			_, err = backend.SetPhraseIfNotExists(collection, path, newValue)
			if err == nil {
				t.Fatal("expected error for phrase that already exists")
			}
		})
	}
}
