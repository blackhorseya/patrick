//go:build wireinject

package repo

import (
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo() (IProjectRepo, error) {
	panic(wire.Build(testProviderSet))
}
