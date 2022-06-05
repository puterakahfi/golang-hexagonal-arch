package main

import (
	"golang-hexagonal-arch/component/config"
	"golang-hexagonal-arch/config/gorm/mysql"
	"golang-hexagonal-arch/migration"
	"golang-hexagonal-arch/server"
	"log"
)

const (
	CONFIG_TYPE    string = "env"
	SERVER_ADAPTER string = "gin"
)

func main() {

	// load configuration
	var config, err = config.InitConfig(CONFIG_TYPE)

	if err != nil {
		log.Fatal("Config error: ", err)
	}

	err = config.LoadConfig()
	if err != nil {
		log.Fatal("Config error: ", err)
	}

	db, err := mysql.Connect(config.GetConfig("DATABASE_MYSQL_DSN"))

	if err != nil {
		log.Fatal("Db connection error", err)
	}

	migration.AutoMigrate(db)

	server, _ := server.GetServer("fiber")

	//book.InitModule(db, server.GetServerInstance())

	server.Run()
}
