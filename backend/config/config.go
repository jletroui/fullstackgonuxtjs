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
	BasePath              string
}

// When deployed, the config is encrypted
var deployedEnvs = map[string]struct{}{
	"production": {},
	"staging":    {},
}

func LoadConfigFromEnv(defaultEnv string) (*Config, error) {
	env, envExists := os.LookupEnv("ENV")
	if !envExists {
		env = defaultEnv
	}
	basePath := basePath(env)
	_, isDeployed := deployedEnvs[env]
	var data []byte
	var err error
	if isDeployed {
		data, err = ejson.DecryptFile(fmt.Sprintf("%sconfig/backend.%s.ejson", basePath, env), "/opt/ejson/keys", "")
		if err != nil {
			return nil, err
		}
	} else {
		data, err = os.ReadFile(fmt.Sprintf("%sconfig/backend.%s.json", basePath, env))
		if err != nil {
			return nil, err
		}
	}
	var res = new(Config)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	res.BasePath = basePath
	return res, nil
}

func basePath(env string) string {
	// This is a bit of logic allowing to execute from the project root or from within backend/.
	// Useful to not have to fiddle with config when executing tests, debugging the project, etc...
	switch env {
	case "dev":
		return "../" // Running app in dev, we are in backend/
	case "test", "ci":
		// This is britle but sufficient. A stronger (but over engineered) technique would be to search recursively in each parent directory.
		return "../../" // Running tests, we are in backend/somepackage
	case "production", "staging":
		return ""
	default:
		panic(fmt.Sprintf("Don't know env '%s'.", env))
	}
}
