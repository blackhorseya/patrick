//go:build wireinject

package project

import (
	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(logger *zap.Logger, repo repo.IProjectRepo) (IProjectBiz, error) {
	panic(wire.Build(testProviderSet))
}
