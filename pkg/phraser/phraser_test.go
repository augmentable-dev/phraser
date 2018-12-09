package phraser_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/augmentable-opensource/phraser/pkg/phraser"
	"github.com/dgraph-io/badger"
)

var p *phraser.Phraser

func TestMain(m *testing.M) {
	opts := badger.DefaultOptions
	dir := "/tmp/phraser-test-" + time.Now().String()
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	p = phraser.NewPhraser(db)
	defer p.Close()
	os.Exit(m.Run())
}

func TestBasic(t *testing.T) {
	collection := "test-collection"
	path := "hello"
	value := "world"

	_, err := p.SetPhrase(collection, path, value)
	if err != nil {
		t.Fatal(err)
	}

	phrase, err := p.GetPhrase(collection, path)
	if err != nil {
		t.Fatal(err)
	}

	if phrase.Collection != collection {
		t.Fatalf("expected collection %s; got %s", collection, phrase.Collection)
	}
	if phrase.Path != path {
		t.Fatalf("expected path %s; got %s", path, phrase.Path)
	}
	if phrase.Value != value {
		t.Fatalf("expected value %s; got %s", value, phrase.Value)
	}

	newValue := "world!"
	_, err = p.SetPhraseIfNotExists(collection, path, newValue)
	if err == nil {
		t.Fatal("expected error for phrase that already exists")
	}
}
