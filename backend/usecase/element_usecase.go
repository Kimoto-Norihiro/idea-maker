package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Kimoto-Norihiro/idea-maker/repository"
	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type ElementUseCase struct {
	repository repository.IElementRepository
	validate   *validator.Validate
}

func NewElementUseCase(r repository.IElementRepository) *ElementUseCase {
	return &ElementUseCase{
		repository: r,
		validate:   validator.New(),
	}
}

func (mu *ElementUseCase) CreateElement(c *gin.Context, m model.Element) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.CreateElement(m)
}

func (mu *ElementUseCase) UpdateElement(c *gin.Context, m model.Element) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.UpdateElement(m)
}