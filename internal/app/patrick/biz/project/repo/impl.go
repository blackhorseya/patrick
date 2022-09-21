package repo

import (
	"html/template"
	"os"
	"path"

	"go.uber.org/zap"
)

type impl struct {
	logger *zap.Logger
}

// NewImpl return IProjectRepo
func NewImpl(logger *zap.Logger) IProjectRepo {
	return &impl{
		logger: logger.With(zap.String("type", "ProjectRepo")),
	}
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
