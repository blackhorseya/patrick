package repo

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

const (
	TempPath = "/tmp/patrick"
)

type SuiteTest struct {
	suite.Suite
	repo IProjectRepo
}

func (s *SuiteTest) SetupTest() {
	repo, err := CreateRepo()
	if err != nil {
		panic(err)
	}
	s.repo = repo

	createFolder(TempPath)
}

func createFolder(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, 0751)
	}
}

func (s *SuiteTest) TearDownTest() {
	_ = os.RemoveAll(TempPath)
}

func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

func (s *SuiteTest) Test_impl_WriteFile() {
	type args struct {
		path      string
		body      []byte
		data      any
		overwrite bool
		mock      func()
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		wantBody string
	}{
		{
			name:     "write to a not exists file then ok",
			args:     args{path: TempPath + "/test.txt", body: []byte("test"), overwrite: false},
			wantErr:  false,
			wantBody: "test",
		},
		{
			name: "write to exists file with not overwrite then ok",
			args: args{path: TempPath + "/test.txt", body: []byte("modified"), overwrite: false, mock: func() {
				file, err := os.Create(TempPath + "/test.txt")
				if err != nil {
					panic(err)
				}
				defer file.Close()

				_, _ = file.Write([]byte("test"))
			}},
			wantErr:  false,
			wantBody: "test",
		},
		{
			name: "write to exists file with overwrite then ok",
			args: args{path: TempPath + "/test.txt", body: []byte("modified"), overwrite: true, mock: func() {
				file, err := os.Create(TempPath + "/test.txt")
				if err != nil {
					panic(err)
				}
				defer file.Close()

				_, _ = file.Write([]byte("test"))
			}},
			wantErr:  false,
			wantBody: "modified",
		},
		{
			name: "not exists folder then write file success",
			args: args{path: TempPath + "/l1/test.txt", body: []byte("test"), overwrite: true, mock: func() {
				filePath := TempPath + "/l1/test.txt"
				if _, err := os.Stat(path.Dir(filePath)); os.IsNotExist(err) {
					_ = os.MkdirAll(path.Dir(filePath), 0751)
				}
				file, err := os.Create(filePath)
				if err != nil {
					panic(err)
				}
				defer file.Close()

				_, _ = file.Write([]byte("test"))
			}},
			wantErr:  false,
			wantBody: "test",
		},
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.repo.WriteFile(tt.args.path, tt.args.body, tt.args.data, tt.args.overwrite); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			gotBody, _ := os.ReadFile(tt.args.path)
			if string(gotBody) != tt.wantBody {
				t.Errorf("WriteFile() getBody = %v, wantBody %s", string(gotBody), tt.wantBody)
			}

			_ = os.RemoveAll(TempPath + "*")
		})
	}
}
