package project

import (
	"github.com/blackhorseya/patrick/pkg/entity/project"
	"github.com/google/wire"
)

// IProjectBiz declare project functions
//
//go:generate mockery --all --inpackage
type IProjectBiz interface {
	// InitProject serve caller to initialize project
	InitProject(prj *project.Info) error
}

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet()
