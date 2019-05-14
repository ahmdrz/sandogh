package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	SecretKey     string `envconfig:"secret_key" required:"true"`
	ListenAddr    string `envconfig:"listen_addr" default:"0.0.0.0:8080"`
	BaseDirectory string `envconfig:"base_directory" required:"true"`
}

func Read() (*Configuration, error) {
	config := &Configuration{}
	err := envconfig.Process("SERVICE", config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
