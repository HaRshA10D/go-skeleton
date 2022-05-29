package config

type AppConfig struct {
	LogLevel      string
	NetworkClient NetworkClientConfig
}

type NetworkClientConfig struct {
	TimeoutInMS int
}
