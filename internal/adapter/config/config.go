package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const MONGO_PREFIX = "MONGO"
const JWT_PREFIX = "JWT"
const HTTP_PREFIX = "HTTP"

type Config struct {
	MongoConfig      MongoConfig
	JwtConfig        JwtConfig
	HttpServerConfig HttpServerConfig
	GrpcServerConfig GrpcServerConfig
}

type MongoConfig struct {
	URI               string        `envconfig:"URI" required:"true"`
	DatabaseName      string        `envconfig:"DATABASE_NAME" required:"true"`
	ConnectionTimeout time.Duration `envconfig:"CONNECTION_TIMEOUT" default:"5s"`
	ReadTimeout       time.Duration `envconfig:"READ_TIMEOUT" default:"5s"`
}

type JwtConfig struct {
	Secret  string        `envconfig:"SECRET" required:"true"`
	Expired time.Duration `envconfig:"EXPIRE" default:"30m"`
}

type HttpServerConfig ServerConfig
type GrpcServerConfig ServerConfig
type ServerConfig struct {
	Port         string        `envconfig:"SERVER_PORT" required:"true"`
	ReadTimeout  time.Duration `envconfig:"SERVER_READ_TIMEOUT" default:"10s"`
	WriteTimeout time.Duration `envconfig:"SERVER_WRITE_TIMEOUT" default:"15s"`
	IdleTimeout  time.Duration `envconfig:"SERVER_IDLE_TIMEOUT" default:"60s"`
}

func newMongoConfig() MongoConfig {
	var cfg MongoConfig
	envconfig.MustProcess(MONGO_PREFIX, &cfg)

	return cfg
}

func newJwtConfig() JwtConfig {
	var cfg JwtConfig
	envconfig.MustProcess(JWT_PREFIX, &cfg)

	return cfg
}

func newHttpServerConfig() HttpServerConfig {
	var cfg ServerConfig
	envconfig.MustProcess(HTTP_PREFIX, &cfg)

	return HttpServerConfig(cfg)
}

func NewConfig() Config {
	return Config{
		MongoConfig:      newMongoConfig(),
		JwtConfig:        newJwtConfig(),
		HttpServerConfig: newHttpServerConfig(),
	}
}
