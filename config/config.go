package config

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config is a config :)
type Config struct {
	LogLevel    string `envconfig:"API_LOG_LEVEL"`
	HTTPAddr    string `envconfig:"API_HTTP_ADDR"`
	TRUEResult  string `envconfig:"API_TRUE_RESULT"`
	FALSEResult string `envconfig:"API_FALSE_RESULT"`
}

var (
	config Config
	once   sync.Once
)

// Get reads config from environment. Once.
func Get() *Config {
	once.Do(func() {
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configuration:", string(configBytes))
	})
	return &config
}
