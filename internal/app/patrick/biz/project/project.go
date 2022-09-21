package project

import (
	"github.com/google/wire"
)

// IProjectBiz declare project functions
//
//go:generate mockery --all --inpackage
type IProjectBiz interface {
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
