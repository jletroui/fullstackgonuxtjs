package logic_test

import (
	"backend/config"
	"database/sql"
	"fmt"
	"testing"
)

var Cfg *config.Config
var Db *sql.DB

func TestMain(m *testing.M) {
	var err error
	Cfg, err = config.LoadConfigFromEnv("test")
	if err != nil {
		panic(fmt.Sprintf("Cannot load config: %s", err))
	}
	err = config.MigrateDatabase(Cfg)
	if err != nil {
		panic(fmt.Sprintf("Cannot migrate database: %s", err))
	}
	Db, err = config.NewDatabase(Cfg)
	if err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %s", err))
	}
	defer Db.Close()
	m.Run()
}
