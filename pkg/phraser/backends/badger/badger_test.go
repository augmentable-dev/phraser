package badger_test

import (
	"log"
	"os"
	"testing"

	badger_backend "github.com/augmentable-opensource/phraser/pkg/phraser/backends/badger"
)

var (
	DBDir         = os.Getenv("DB_DIR")
	BadgerBackend *badger_backend.Backend
)

func init() {
	if DBDir == "" {
		DBDir = "/tmp/phraser-badger-backend-test"
	}
}

func setup() {
	var err error
	BadgerBackend, err = badger_backend.NewBadgerBackendFromDir(DBDir)
	if err != nil {
		log.Fatal(err)
	}
}

func cleanup() {
	BadgerBackend.Close()
}

func TestMain(m *testing.M) {
	setup()
	defer cleanup()
	os.Exit(m.Run())
}

func TestBasic(t *testing.T) {
	collection := "test-collection"
	path := "test/path/hello"
	value := "world"
	phrase, err := BadgerBackend.SetPhrase(collection, path, value)
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

	phrase, err = BadgerBackend.GetPhrase(collection, path)
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
	_, err = BadgerBackend.SetPhraseIfNotExists(collection, path, newValue)
	if err == nil {
		t.Fatal("expected error for phrase that already exists")
	}
}
