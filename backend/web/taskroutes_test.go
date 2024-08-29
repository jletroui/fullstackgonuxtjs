package web_test

import (
	"backend/config"
	"backend/mocks"
	"backend/web"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type TaskRoutesTestSuite struct {
	suite.Suite
	router          *gin.Engine
	tasksRepository *mocks.TaskRepository
	sessionVerifier *TestSessionVerifier
	recorder        *httptest.ResponseRecorder
}

func TestTaskRoutes(t *testing.T) {
	suite.Run(t, new(TaskRoutesTestSuite))
}

func (suite *TaskRoutesTestSuite) SetupSuite() {
	suite.tasksRepository = mocks.NewTaskRepository(suite.T())
	suite.sessionVerifier = &TestSessionVerifier{}
	router, err := config.CreateWebServer(Cfg, func(apiRoutes *gin.RouterGroup) {
		web.InstallTaskRoutes(apiRoutes, suite.sessionVerifier, suite.tasksRepository)
	})
	suite.NoError(err)
	suite.router = router
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

func (suite *TaskRoutesTestSuite) TestWhenCreateTask() {
	suite.tasksRepository.On("CreateTask", "abc").Return(nil).Once()

	req, _ := http.NewRequest("POST", "/api/tasks", strings.NewReader("{\"description\":\"abc\"}"))
	suite.router.ServeHTTP(suite.recorder, req)

	suite.Equal(200, suite.recorder.Code)
	suite.Equal("", suite.recorder.Body.String())
	suite.tasksRepository.AssertCalled(suite.T(), "CreateTask", "abc")
}

func (suite *TaskRoutesTestSuite) TestWhenCreateTaskAndNoSession() {
	suite.sessionVerifier.FailNextVerification()

	req, _ := http.NewRequest("POST", "/api/tasks", strings.NewReader("{\"description\":\"abc\"}"))
	suite.router.ServeHTTP(suite.recorder, req)

	suite.Equal(401, suite.recorder.Code)
}

// TODO: in a real app, also test errors from the TaskRepository return values
