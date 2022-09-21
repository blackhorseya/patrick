package repo

import (
	"github.com/google/wire"
)

// IProjectRepo declare project functions
//
//go:generate mockery --all --inpackage
type IProjectRepo interface {
	// WriteFile serve caller to given path and body to write file
	WriteFile(path string, body []byte, overwrite bool) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
