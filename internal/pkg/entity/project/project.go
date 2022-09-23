package project

// Info define a project information
type Info struct {
	PkgName      string `json:"pkg_name"`
	AbsolutePath string `json:"absolute_path"`
	AppName      string `json:"app_name"`
}
