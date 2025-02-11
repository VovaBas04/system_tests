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

func SwitchReplicatedDbByEnv() ReplicateConfigDb {
	switch os.Getenv("ENV") {
	case "PROD":
		return &ReplicatedLocalhostDb{}
	case "DEV":
		return &ReplicatedDockerDb{}
	default:
		return &ReplicatedLocalhostDb{}
	}
}
