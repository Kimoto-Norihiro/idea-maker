package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type IdeaRepository struct {
	db *gorm.DB
}

func NewIdeaRepository(db *gorm.DB) *IdeaRepository {
	return &IdeaRepository{db}
}

func (ir *IdeaRepository) CreateIdea(m model.Idea) error {
	return ir.db.Create(&m).Error
}

func (ir *IdeaRepository) UpdateIdea(m model.Idea) error {
	return ir.db.Save(&m).Error
}

func (ir *IdeaRepository) DeleteIdea(m model.Idea) error {
	return ir.db.Delete(&m).Error
}