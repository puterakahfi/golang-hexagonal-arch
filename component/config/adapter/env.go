package adapter

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type configEnv struct {
	envFile string
}

func NewEnvConfig() *configEnv {
	envFile := ".env"
	return &configEnv{envFile: envFile}
}

func (cenv *configEnv) LoadConfig() error {
	log.Println("Config Env init")

	err := godotenv.Load(cenv.envFile)
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return err
	}

	return nil

}

func (c *configEnv) GetConfig(key string) string {
	return os.Getenv(key)
}
