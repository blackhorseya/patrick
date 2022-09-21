package project

import (
	"fmt"
	"html/template"
	"os"

	"github.com/blackhorseya/patrick/pkg/tpl"
)

// Info define a project information
type Info struct {
	PkgName      string `json:"pkg_name"`
	AbsolutePath string `json:"absolute_path"`
	AppName      string `json:"app_name"`
}

func (p *Info) Create() error {
	// check if AbsolutePath exists
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		// create directory
		err = os.Mkdir(p.AbsolutePath, 0754)
		if err != nil {
			return err
		}
	}

	fmt.Println("Starting create project...")

	filesMap := map[string][]byte{
		"Makefile":                tpl.MakefileTemplate(),
		".gitignore":              tpl.GitignoreTemplate(),
		".golangci.yaml":          tpl.GolangCITemplate(),
		".pre_commit_config.yaml": tpl.PreCommitConfigTemplate(),
	}

	for name, body := range filesMap {
		err := p.createFileFromTemplate(name, body)
		if err != nil {
			return err
		}
	}

	// todo: 2022/9/22|sean|create .pre-commit-config.yaml
	// todo: 2022/9/22|sean|create .cz.yaml
	// todo: 2022/9/22|sean|create Dockerfile
	// todo: 2022/9/22|sean|create standard project layout folders
	// todo: 2022/9/22|sean|create scripts/go.test.sh

	return nil
}

func (p *Info) createFileFromTemplate(name string, body []byte) error {
	path := fmt.Sprintf("%s/%s", p.AbsolutePath, name)

	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = template.Must(template.New(name).Parse(string(body))).Execute(file, p)
	if err != nil {
		return err
	}

	return nil
}
