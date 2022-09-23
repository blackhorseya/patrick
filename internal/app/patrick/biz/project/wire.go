//go:build wireinject

package project

import (
	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/google/wire"
)

var testProviderSet = wire.NewSet(NewImpl)

func CreateBiz(repo repo.IProjectRepo) (IProjectBiz, error) {
	panic(wire.Build(testProviderSet))
}
