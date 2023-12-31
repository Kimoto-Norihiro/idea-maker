package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/repository"
	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type ThemeUseCase struct {
	repository repository.IThemeRepository
	validate *validator.Validate
}

func NewThemeUseCase(r repository.IThemeRepository) *ThemeUseCase {
	return &ThemeUseCase{
		repository: r,
		validate: validator.New(),
	}
}

func (mu *ThemeUseCase) CreateTheme(c *gin.Context, m model.Theme) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.CreateTheme(m)
}

func (mu *ThemeUseCase) UpdateTheme(c *gin.Context, m model.Theme) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.UpdateTheme(m)
}

func (mu *ThemeUseCase) ShowTheme(c *gin.Context, uid string, themeId uint) (model.Theme, error) {
	return mu.repository.ShowTheme(uid, themeId)
}

func (mu *ThemeUseCase) DeleteTheme(c *gin.Context, m model.Theme) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.DeleteTheme(m)
}