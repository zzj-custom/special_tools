package config

import (
	"special_tools/backend/pkg/email"
	"special_tools/backend/pkg/iMysql"
	"special_tools/backend/pkg/iRedis"
	"special_tools/backend/pkg/sms"
)

type Config struct {
	Database map[string]*iMysql.Database `toml:"database"`
	Redis    []*iRedis.MultiDialConfig   `toml:"redis"`
	Email    *email.Config               `toml:"email"`
	Sms      *sms.Config                 `toml:"sms"`
}
