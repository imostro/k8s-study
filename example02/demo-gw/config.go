package main

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Env        string `yaml:"env"`
	ServerAddr string `yaml:"server_addr"`
	Greeting   string `yaml:"greeting"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config := Default()
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func Default() *Config {
	return &Config{Server: Server{
		Env:        "DEBUG",
		ServerAddr: ":9000",
		Greeting:   "hello gay",
	}}
}
