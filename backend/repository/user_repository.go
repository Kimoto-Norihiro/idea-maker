package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (mr *UserRepository) CreateUser(m model.User) error {
	return mr.db.Create(&m).Error
}

func (mr *UserRepository) ShowUser(uid string) (model.User, error) {
	var m model.User
	err := mr.db.Preload("Themes").Preload("Themes.Ideas").Preload("Themes.Elements").Where("uid = ?", uid).First(&m).Error
	return m, err
}