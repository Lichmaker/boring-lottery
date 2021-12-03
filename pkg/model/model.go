package model

import (
	"fmt"

	"github.com/lichmaker/boring-lottery/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var (
		host     = config.Viper.GetString("database.mysql.host")
		port     = config.Viper.GetString("database.mysql.port")
		database = config.Viper.GetString("database.mysql.database")
		username = config.Viper.GetString("database.mysql.username")
		password = config.Viper.GetString("database.mysql.password")
		charset  = config.Viper.GetString("database.mysql.charset")
	)
	// dsn 格式可以参考 https://github.com/go-sql-driver/mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")

	// gormConfig := mysql.New(mysql.Config{
	// 	DSN: dsn,
	// })

	var level logger.LogLevel
	if config.Viper.GetBool("app.debug") {
		level = logger.Warn
	} else {
		level = logger.Error
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level),
	})
	if err != nil {
		panic(err)
	}
	return DB
}
