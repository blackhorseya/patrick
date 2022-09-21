//go:build wireinject

package repo

import (
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateRepo(logger *zap.Logger) (IProjectRepo, error) {
	panic(wire.Build(testProviderSet))
}
