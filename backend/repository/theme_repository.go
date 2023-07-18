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

func (tr *ThemeRepository) IndexTheme(uid string) ([]model.Theme, error) {
	var m []model.Theme
	err := tr.db.Where("user_uid = ?", uid).Find(&m).Error
	return m, err
}