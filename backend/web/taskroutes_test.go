package web_test

import (
	"backend/config"
	"backend/mocks"
	"backend/web"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TaskRoutesTestSuite struct {
	suite.Suite
	router          *gin.Engine
	tasksRepository *mocks.TaskRepository
	recorder        *httptest.ResponseRecorder
}

func TestTaskRoutes(t *testing.T) {
	suite.Run(t, new(TaskRoutesTestSuite))
}

func (suite *TaskRoutesTestSuite) SetupSuite() {
	suite.tasksRepository = mocks.NewTaskRepository(suite.T())
	suite.router = config.CreateWebServer(Cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes, suite.tasksRepository)
	})
}

func (suite *TaskRoutesTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
}

func (suite *TaskRoutesTestSuite) TestWhenTaskCount() {
	suite.tasksRepository.On("Count").Return(7, nil).Once()

	req, _ := http.NewRequest("GET", "/api/tasks/count", nil)
	suite.router.ServeHTTP(suite.recorder, req)

	suite.Equal(200, suite.recorder.Code)
	expected, _ := json.Marshal(gin.H{"count": 7})
	suite.Equal(string(expected), suite.recorder.Body.String())
}
