package main

import (
	"golang-hexagonal-arch/component/config"
	"golang-hexagonal-arch/config/gorm/mysql"
	"golang-hexagonal-arch/module/book"
	"golang-hexagonal-arch/module/book/entity"
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
	db.AutoMigrate(&entity.Book{})

	server := server.InitGinServer()

	book.InitModule(db, server.GetServerInstance())

	server.Run()
}
