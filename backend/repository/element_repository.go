package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type ElementRepository struct {
	db *gorm.DB
}

func NewElementRepository(db *gorm.DB) *ElementRepository {
	return &ElementRepository{db}
}

func (er *ElementRepository) CreateElement(m model.Element) error {
	return er.db.Create(&m).Error
}