package config

type serverConfigInterface interface {
	HttpServerConfig | GrpcServerConfig
}

func ToServerConfig[T serverConfigInterface](config T) ServerConfig {
	return ServerConfig(config)
}
