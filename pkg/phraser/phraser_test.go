package phraser_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/augmentable-opensource/phraser/pkg/phraser"
	badger_backend "github.com/augmentable-opensource/phraser/pkg/phraser/backends/badger"
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

	backend := badger_backend.NewBadgerBackend(db)
	p = phraser.NewPhraser(backend)
	defer p.Close()
	os.Exit(m.Run())
}

func TestBasic(t *testing.T) {

}
