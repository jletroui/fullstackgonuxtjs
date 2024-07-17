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

func (suite *PostgresTaskRepositoryTestSuite) SetupTest() {
	_, err := Db.Exec("TRUNCATE tasks")
	suite.NoError(err)
}

func (suite *PostgresTaskRepositoryTestSuite) TestWhenNoTaskThenCountReturn0() {
	count, err := suite.sut.Count()
	suite.NoError(err)
	suite.Equal(0, count)
}

func (suite *PostgresTaskRepositoryTestSuite) TestWhen2TasksThenCountReturn2() {
	err := suite.sut.CreateTask("abc")
	suite.NoError(err)
	err = suite.sut.CreateTask("def")
	suite.NoError(err)

	count, err := suite.sut.Count()
	suite.NoError(err)
	suite.Equal(2, count)
}
