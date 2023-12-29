package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

// InitDB initializes the database connection
func InitializesConnection() (*gorm.DB, error) {
	conf := GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DBUsername, conf.DBPassword, conf.DBHost, conf.DBPort, conf.DBName)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	Connection = db

	return Connection, nil
}


