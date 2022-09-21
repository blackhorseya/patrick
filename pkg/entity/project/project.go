package project

import (
	"fmt"
	"os"
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

	// todo: 2022/9/22|sean|impl me

	return nil
}
