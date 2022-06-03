package config

import (
	"fmt"
	"golang-hexagonal-arch/component/config/adapter"
)

func InitConfig(configType string) (ConfigServiceInterface, error) {

	switch configType {
	case "env":
		return adapter.NewEnvConfig(), nil
	case "json":
		return adapter.NewJsonConfig(), nil
	}

	return nil, fmt.Errorf("config type " + configType + " not supported")

}
