package config

import pkghttp "github.com/user-name/skeleton-name/pkg/http"

type AppConfig struct {
	LogLevel      string
	Server        pkghttp.ServerConfig
	NetworkClient NetworkClientConfig
}

type NetworkClientConfig struct {
	TimeoutInMS int
}
