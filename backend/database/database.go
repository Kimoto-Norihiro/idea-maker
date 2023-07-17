package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

func NewMySql(dns string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.User{})
	return db, nil
}