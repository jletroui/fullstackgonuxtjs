package web_test

import (
	"backend/config"
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var Cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	fmt.Println("Loading test config...")
	env, envExists := os.LookupEnv("ENV")
	if envExists && env == "ci" {
		Cfg, err = config.LoadConfigFromEnv("ci")

	} else {
		Cfg, err = config.LoadConfigFromEnv("test")
	}
	if err != nil {
		panic(fmt.Sprintf("Cannot load config: %s", err))
	}
	fmt.Println("Loaded test config.")
	m.Run()
}

type TestSessionVerifier struct {
	failNextVerification bool
	userID               string
}

func (tsv *TestSessionVerifier) VerifySession(c *gin.Context) {
	if tsv.failNextVerification {
		// Roughly simulate SuperTokens behaviour
		c.AbortWithStatus(401)
	} else {
		c.Next()
	}
}

func (tsv *TestSessionVerifier) GetUserID(c *gin.Context) string {
	return tsv.userID
}

func (tsv *TestSessionVerifier) FailNextVerification() {
	tsv.failNextVerification = true
}

func (tsv *TestSessionVerifier) SetUserID(userID string) {
	tsv.userID = userID
}
