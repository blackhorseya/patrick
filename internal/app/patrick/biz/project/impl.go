package project

import (
	"fmt"
	"os"

	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/blackhorseya/patrick/internal/pkg/entity/project"
	"github.com/blackhorseya/patrick/internal/pkg/tpl"
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
	// check if AbsolutePath exists
	if _, err := os.Stat(prj.AbsolutePath); os.IsNotExist(err) {
		// create directory
		err = os.MkdirAll(prj.AbsolutePath, 0754)
		if err != nil {
			return err
		}
	}

	i.logger.Info("Starting create project...")

	filesMap := map[string][]byte{
		"Makefile":                tpl.MakefileTemplate(),
		".gitignore":              tpl.GitignoreTemplate(),
		".golangci.yaml":          tpl.GolangCITemplate(),
		".pre-commit-config.yaml": tpl.PreCommitConfigTemplate(),
		".cz.yaml":                tpl.CZTemplate(),
	}

	for fileName, body := range filesMap {
		err := i.repo.WriteFile(fmt.Sprintf("%s/%s", prj.AbsolutePath, fileName), body, prj, false)
		if err != nil {
			return err
		}
	}

	return nil
}
