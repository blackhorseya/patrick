package project

import (
	"testing"

	"github.com/blackhorseya/patrick/internal/app/patrick/biz/project/repo"
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
