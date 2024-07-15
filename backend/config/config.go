package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/ejson"
)

type Config struct {
	Port                  int    `json:"_port"`
	PostgresWaitTimeoutMs int    `json:"_postgresWaitTimeoutMs"`
	PostgresHost          string `json:"_postgresHost"`
	PostgresDatabase      string `json:"_postgresDatabase"`
	PostgresAdminUser     string `json:"_postgresAdminUser"`
	PostgresAdminPassword string `json:"postgresAdminPassword"`
	PostgresAppUser       string `json:"_postgresAppUser"`
	PostgresAppPassword   string `json:"postgresAppPassword"`
}

var encryptedConfigEnvs = map[string]struct{}{
	"production": {},
	"staging":    {},
}

func LoadFromEnv(defaultEnv string) (*Config, error) {
	env, envExists := os.LookupEnv("ENV")
	if !envExists {
		env = defaultEnv
	}
	_, isEncrypted := encryptedConfigEnvs[env]
	var data []byte
	var err error
	if isEncrypted {
		var configFilePath = fmt.Sprintf("config/backend.%s.ejson", env)
		data, err = ejson.DecryptFile(configFilePath, "/opt/ejson/keys", "")
		if err != nil {
			return nil, err
		}
	} else {
		var configFilePath = fmt.Sprintf("config/backend.%s.json", env)
		data, err = os.ReadFile(configFilePath)
		if err != nil {
			return nil, err
		}
	}
	var res = new(Config)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}
