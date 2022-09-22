package biz

import (
	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project"
	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(
	project.ProviderSet,
)
