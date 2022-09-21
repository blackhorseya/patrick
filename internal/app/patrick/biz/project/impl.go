package project

import (
	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/blackhorseya/patrick/internal/pkg/entity/project"
	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
	repo   repo.IProjectRepo
}

// NewImpl return IProjectBiz
func NewImpl(logger *zap.Logger, repo repo.IProjectRepo) IProjectBiz {
	return &impl{
		logger: logger.With(zap.String("type", "ProjectBiz")),
		repo:   repo,
	}
}

func (i *impl) InitProject(prj *project.Info) error {
	// TODO implement me
	panic("implement me")
}
