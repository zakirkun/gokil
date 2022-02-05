package config

import (
	"fmt"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DSN = "%s:%s@tcp(%s:%d)/%s?parseTime=true"
)

func NewMysql(configuration Config) *gorm.DB {
	PORT, err := strconv.Atoi(configuration.Get("DB_PORT"))
	if err != nil {
		PORT = 3306
	}

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(DSN, configuration.Get("DB_USER"), configuration.Get("DB_PASSWORD"), configuration.Get("DB_HOST"), PORT, configuration.Get("DB_NAME"))), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
