package project

import (
	"fmt"
	"os"

	"github.com/blackhorseya/gocommon/pkg/contextx"
	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/blackhorseya/patrick/internal/pkg/entity/project"
	"github.com/blackhorseya/patrick/internal/pkg/tpl"
)

type impl struct {
	repo repo.IProjectRepo
}

// NewImpl return IProjectBiz
func NewImpl(repo repo.IProjectRepo) IProjectBiz {
	return &impl{
		repo: repo,
	}
}

func (i *impl) InitProject(ctx contextx.Contextx, prj *project.Info) error {
	// check if AbsolutePath exists
	if _, err := os.Stat(prj.AbsolutePath); os.IsNotExist(err) {
		// create directory
		err = os.MkdirAll(prj.AbsolutePath, 0754)
		if err != nil {
			return err
		}
	}

	// ctx.Info("Starting create project...")

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
