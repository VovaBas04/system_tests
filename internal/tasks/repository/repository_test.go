package repository

import (
	"ginProject1/pkg/database/postgres"
	"github.com/stretchr/testify/suite"
	"testing"
)

type RepositoryTestSuite struct {
	suite.Suite
	repository IRepository
	closer     func()
}

func (s *RepositoryTestSuite) SetupTest() {
	configPath := "/home/oem/GolandProjects/ginProject1/config/tasks.yaml"
	postgresDb, err := postgres.Connect(configPath)
	s.Assert().Nil(err)
	s.repository = NewTaskRepository(nil, postgresDb)
	s.closer = func() {
		postgresDb.Close()
	}
}

func (s *RepositoryTestSuite) TearDownTest() {
	s.closer()
}

func (s *RepositoryTestSuite) TestLogic() {
	err := s.repository.SaveAction("test", "test")
	s.Assert().Nil(err)
}

func TestRun(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
