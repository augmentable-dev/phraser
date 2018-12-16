package phraser

// Phraser is a key-value store for your application's phrases
type Phraser struct {
	Backend
}

// Backend defines a backend store for phraser
type Backend interface {
	SetPhrase(collection, path, value string) (*Phrase, error)
	SetPhraseIfNotExists(collection, path, value string) (*Phrase, error)
	GetPhrase(collection, path string) (*Phrase, error)
	Close() error
}

// Phrase is a phrase
type Phrase struct {
	Collection string
	Path       string
	Value      string
}

// NewPhraser creates a new phraser
func NewPhraser(backend Backend) *Phraser {
	return &Phraser{Backend: backend}
}

// BuildPath constructs a KV prefix path from a collection and path
func BuildPath(collection, path string) []byte {
	return []byte(collection + ":" + path)
}

// SetPhraseIfNotExists sets a phrase only if it doesn't exist yet, returns (nil, error) otherwise
func (phraser *Phraser) SetPhraseIfNotExists(collection, path, value string) (*Phrase, error) {
	return phraser.Backend.SetPhraseIfNotExists(collection, path, value)
}

// GetPhrase gets a phrase
func (phraser *Phraser) GetPhrase(collection, path string) (*Phrase, error) {
	return phraser.Backend.GetPhrase(collection, path)
}

// ListPhrases lists phrases
func (phraser *Phraser) ListPhrases(collection, path string) (*Phrase, error) {
	return nil, nil
}

// Close up the phraser
func (phraser *Phraser) Close() error {
	return phraser.Close()
}
