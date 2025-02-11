package api_gateway

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigUrl interface {
	GetUrlTasks() string
	GetUrlCodeChecker() string
}
type Url struct {
	CodeChecker string `yaml:"codechecker"`
	Tasks       string `yaml:"tasks"`
}

type Docker struct {
	Url Url `yaml:"docker"`
}

type Localhost struct {
	Url Url `yaml:"localhost"`
}

func (entity *Localhost) GetUrlCodeChecker() string {
	return entity.Url.CodeChecker
}

func (entity *Localhost) GetUrlTasks() string {
	return entity.Url.Tasks
}

func (entity *Docker) GetUrlCodeChecker() string {
	return entity.Url.CodeChecker
}

func (entity *Docker) GetUrlTasks() string {
	return entity.Url.Tasks
}

func getConfig(configName string) (ConfigUrl, error) {
	config := initConfig()
	data, err := os.ReadFile(configName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func initConfig() ConfigUrl {
	switch os.Getenv("ENV") {
	case "PROD":
		return &Localhost{}
	case "DEV":
		return &Docker{}
	default:
		return &Localhost{}
	}
}
