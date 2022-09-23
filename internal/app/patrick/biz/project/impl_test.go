package project

import (
	"testing"

	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
	"github.com/blackhorseya/patrick/internal/pkg/entity/project"
	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	repo *repo.MockIProjectRepo
	biz  IProjectBiz
}

func (s *SuiteTest) SetupTest() {
	s.repo = new(repo.MockIProjectRepo)

	biz, err := CreateBiz(s.repo)
	if err != nil {
		panic(err)
	}
	s.biz = biz
}

func (s *SuiteTest) TearDownTest() {
}

func TestSuiteTest(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

func (s *SuiteTest) Test_impl_InitProject() {
	type args struct {
		prj  *project.Info
		mock func()
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			if tt.args.mock != nil {
				tt.args.mock()
			}

			if err := s.biz.InitProject(tt.args.prj); (err != nil) != tt.wantErr {
				t.Errorf("InitProject() error = %v, wantErr %v", err, tt.wantErr)
			}

			s.repo.AssertExpectations(t)
		})
	}
}
