package database

import "github.com/dhuki/go-template/internal/infra/configloader"

type connectionDBClient struct {
	conf *configloader.DatabaseConfig
}

func NewConnectionDBClient(conf *configloader.DatabaseConfig) *connectionDBClient {
	return &connectionDBClient{
		conf: conf,
	}
}
