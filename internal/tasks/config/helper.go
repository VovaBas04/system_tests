package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

var configPath = "config/tasks.yaml"

type TokenConfig struct {
	Token string `yaml:"secret-token"`
}
type TestConfig struct {
	TestCase int `yaml:"testcase"`
}

func GetTestCases() (int, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return 0, err
	}
	config := &TestConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return 0, err
	}

	return config.TestCase, nil
}

func GetBasicToken() (string, error) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return "", err
	}
	config := &TokenConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return "", err
	}

	return config.Token, nil
}
