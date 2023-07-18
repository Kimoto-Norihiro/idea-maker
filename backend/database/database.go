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
	// drop tables
	// db.Migrator().DropTable(&model.User{}, &model.Theme{}, &model.Idea{}, &model.Element{})

	db.AutoMigrate(&model.User{}, &model.Theme{}, &model.Idea{}, &model.Element{})

	return db, nil
}