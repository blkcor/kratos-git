package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDB(dsn string) error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&UserBasic{}, &RepoBasic{}, &RepoUser{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
