package web_test

import (
	"backend/config"
	"fmt"
	"testing"
)

var Cfg *config.Config

func TestMain(m *testing.M) {
	var err error
	Cfg, err = config.LoadConfigFromEnv("test")
	if err != nil {
		panic(fmt.Sprintf("Cannot load config: %s", err))
	}
	m.Run()
}
