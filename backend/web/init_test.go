package web_test

import (
	"backend/config"
	"fmt"
	"os"
	"testing"
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
