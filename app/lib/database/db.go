package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/HiWay-Media/go-wire-ddd/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//config, connect and  return connection

type MyDB *gorm.DB

// config *env.Configuration
func NewMyDB(config *env.Configuration) MyDB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connection on %s, err: %s", dsn, err.Error())
	}
	db, _ := conn.DB()
	idle, _ := strconv.Atoi(config.DbIdleConn)
	max, _ := strconv.Atoi(config.DbMaxConn)

	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(max)

	return conn

}
