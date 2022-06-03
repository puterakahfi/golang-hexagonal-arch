package adapter

import (
	"fmt"
	"log"
	"os"
)

type configJson struct {
	jsonFile string
}

func NewJsonConfig() *configJson {
	jsonFile := "config.json"
	return &configJson{jsonFile: jsonFile}
}

func (cenv *configJson) LoadConfig() error {
	log.Println("Config JSON init")
	jsonFile, err := os.Open(cenv.jsonFile)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully Opened " + cenv.jsonFile)

	defer jsonFile.Close()
	return nil
}

func (c *configJson) GetConfig(key string) string {
	return os.Getenv(key)
}
