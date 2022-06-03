package config

type ConfigServiceInterface interface {
	LoadConfig() error
	GetConfig(key string) string
}
