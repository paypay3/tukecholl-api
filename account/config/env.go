package config

import (
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var Env ENV

func init() {
	env := os.Getenv("GO_ENV")

	if err := envconfig.Process(env, &Env); err != nil {
		log.Fatalf("%+v", err)
	}
}

type ENV struct {
	Server
	RDB
	Log
}

type Server struct {
	Port int `envconfig:"SERVER_PORT" required:"true"`
}

type RDB struct {
	Dsn             string        `envconfig:"MYSQL_DSN"               required:"true"`
	MaxConn         int           `envconfig:"MYSQL_MAX_CONN"          default:"25"`
	MaxIdleConn     int           `envconfig:"MYSQL_MAX_IDLE"          default:"25"`
	MaxConnLifetime time.Duration `envconfig:"MYSQL_MAX_CONN_LIFETIME" default:"300s"`
}

type Log struct {
	Debug bool `envconfig:"LOG_DEBUG" default:"false"`
}
