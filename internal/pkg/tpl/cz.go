package tpl

// CZTemplate define a .cz.yaml template
func CZTemplate() []byte {
	return []byte(`commitizen:
  name: cz_conventional_commits
  tag_format: $version
  version: 0.0.0`)
}
