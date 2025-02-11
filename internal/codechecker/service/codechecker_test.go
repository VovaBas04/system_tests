package service

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CodeCheckerServiceHandlerTestSuite struct {
	service *Service
	suite.Suite
}

type RepositoryMock struct {
	mock.Mock
}

func (repository *RepositoryMock) Save() error {
	return repository.Called().Error(0)
}

func (s *CodeCheckerServiceHandlerTestSuite) SetupTest() {
	repository := &RepositoryMock{}
	s.service = &Service{repository: repository}
}

func (s *CodeCheckerServiceHandlerTestSuite) TestTrueCode() {
	isTrue, err := s.service.CheckCode("a = int(input())\nb = int(input())\nprint(a+b)", []map[string]string{
		{
			"a":      "5\n",
			"b":      "5\n",
			"answer": "10\n",
		},
		{
			"a":      "10\n",
			"b":      "10\n",
			"answer": "20\n",
		},
	})
	fmt.Println(err)
	s.Assert().Nil(err)
	s.Assert().Equal(isTrue, true)
}

func (s *CodeCheckerServiceHandlerTestSuite) TestErrorCode() {
	isTrue, err := s.service.CheckCode("a = int(input())d\nb = int(input())\nprint(a+b+c)", []map[string]string{
		{
			"a":      "5\n",
			"b":      "5\n",
			"answer": "10\n",
		},
		{
			"a":      "10\n",
			"b":      "10\n",
			"answer": "20\n",
		},
	})
	fmt.Println(err)
	s.Assert().Nil(err)
	s.Assert().Equal(isTrue, false)
}

func (s *CodeCheckerServiceHandlerTestSuite) TestWrongCode() {
	isTrue, err := s.service.CheckCode("a = int(input())\nb = int(input())\nprint(a+b)", []map[string]string{
		{
			"a":      "5\n",
			"b":      "6\n",
			"answer": "10\n",
		},
		{
			"a":      "10\n",
			"b":      "10\n",
			"answer": "20\n",
		},
	})
	fmt.Println(err)
	s.Assert().Nil(err)
	s.Assert().Equal(isTrue, false)
}

func TestRun(t *testing.T) {
	suite.Run(t, new(CodeCheckerServiceHandlerTestSuite))
}
