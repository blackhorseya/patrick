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

	// create Makefile
	makeFile, err := os.Create(fmt.Sprintf("%s/Makefile", p.AbsolutePath))
	if err != nil {
		return err
	}
	defer makeFile.Close()

	makefileTemplate := template.Must(template.New("Makefile").Parse(string(tpl.MakefileTemplate())))
	err = makefileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	// todo: 2022/9/22|sean|create gitignore
	// todo: 2022/9/22|sean|create .golangci.yaml
	// todo: 2022/9/22|sean|create standard project layout folders
	// todo: 2022/9/22|sean|create scripts/go.test.sh
	// todo: 2022/9/22|sean|create .pre-commit-config.yaml
	// todo: 2022/9/22|sean|create .cz.yaml
	// todo: 2022/9/22|sean|create Dockerfile

	return nil
}
