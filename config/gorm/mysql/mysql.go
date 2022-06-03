package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(configDsn string) (*gorm.DB, error) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db, err := gorm.Open(mysql.Open(configDsn), &gorm.Config{})

	return db, err
}
