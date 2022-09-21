package tpl

// PreCommitConfigTemplate define a .pre-commit-config.yaml template
func PreCommitConfigTemplate() []byte {
	return []byte(`repos:
- hooks:
  - id: commitizen
  repo: https://github.com/commitizen-tools/commitizen
  rev: v2.29.4`)
}
