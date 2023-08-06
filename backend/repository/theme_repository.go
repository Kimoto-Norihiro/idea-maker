package repository

import (
	"gorm.io/gorm"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type ThemeRepository struct {
	db *gorm.DB
}

func NewThemeRepository(db *gorm.DB) *ThemeRepository {
	return &ThemeRepository{db}
}

func (tr *ThemeRepository) CreateTheme(m model.Theme) error {
	return tr.db.Create(&m).Error
}

func (tr *ThemeRepository) UpdateTheme(m model.Theme) error {
	return tr.db.Save(&m).Error
}

func (tr *ThemeRepository) ShowTheme(uid string, themeId uint) (model.Theme, error) {
	var m model.Theme
	err := tr.db.Preload("Ideas").Preload("Elements").First(&m, themeId).Error
	
	return m, err
}

func (tr *ThemeRepository) DeleteTheme(m model.Theme) error {
	return tr.db.Delete(&m).Error
}