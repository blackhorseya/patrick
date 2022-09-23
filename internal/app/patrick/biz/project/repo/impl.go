package repo

import (
	"html/template"
	"os"
	"path"
)

type impl struct {
}

// NewImpl return IProjectRepo
func NewImpl() IProjectRepo {
	return &impl{}
}

func (i *impl) WriteFile(filePath string, tpl []byte, data any, overwrite bool) error {
	_, err := os.Stat(filePath)
	if os.IsExist(err) && !overwrite {
		return nil
	}

	if os.IsNotExist(err) || overwrite {
		if _, err = os.Stat(path.Dir(filePath)); os.IsNotExist(err) {
			_ = os.MkdirAll(path.Dir(filePath), 0751)
		}

		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		err = template.Must(template.New(path.Base(filePath)).Parse(string(tpl))).Execute(file, data)
		if err != nil {
			return err
		}
	}

	return nil
}
