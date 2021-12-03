package bootstrap

import (
	"time"

	"github.com/lichmaker/boring-lottery/app/models/prediction"
	"github.com/lichmaker/boring-lottery/app/models/results"
	"github.com/lichmaker/boring-lottery/pkg/config"
	"github.com/lichmaker/boring-lottery/pkg/model"
	"gorm.io/gorm"
)

func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.Viper.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.Viper.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.Viper.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 创建和维护数据表结构
	migration(db)
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&results.Results{},
		&prediction.Prediction{},
	)
}
