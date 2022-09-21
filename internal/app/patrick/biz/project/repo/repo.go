package repo

import (
	"github.com/google/wire"
)

// IProjectRepo declare project functions
//
//go:generate mockery --all --inpackage
type IProjectRepo interface {
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
