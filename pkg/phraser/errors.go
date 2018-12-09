package phraser

import "errors"

var (
	// ErrPhraseExists is returned when a phrase already exists
	ErrPhraseExists = errors.New("phrase exists")
)
