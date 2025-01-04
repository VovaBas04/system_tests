package database

import (
	"os"
)

func SwitchDbByEnv() ConfigDb {
	switch os.Getenv("ENV") {
	case "PROD":
		return &LocalhostDb{}
	case "DEV":
		return &DockerDb{}
	default:
		return &LocalhostDb{}
	}
}
