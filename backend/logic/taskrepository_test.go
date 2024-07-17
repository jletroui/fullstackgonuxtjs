package logic_test

import (
	"backend/logic"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PostgresTaskRepositoryTestSuite struct {
	suite.Suite
	sut logic.TaskRepository
}

func TestPostgresTaskRepository(t *testing.T) {
	suite.Run(t, new(PostgresTaskRepositoryTestSuite))
}

func (suite *PostgresTaskRepositoryTestSuite) SetupSuite() {
	suite.sut = logic.NewPostgresTaskRepository(Db)
}

func (suite *PostgresTaskRepositoryTestSuite) TestWhenNoTaskThenCountReturn0() {
	count, err := suite.sut.Count()
	suite.NoError(err)
	suite.Equal(0, count)
}
