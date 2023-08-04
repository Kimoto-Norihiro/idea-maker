package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Kimoto-Norihiro/idea-maker/model"
	"github.com/Kimoto-Norihiro/idea-maker/repository"
)

type IdeaUseCase struct {
	repository repository.IIdeaRepository
	validate *validator.Validate
}

func NewIdeaUseCase(r repository.IIdeaRepository) *IdeaUseCase {
	return &IdeaUseCase{
		repository: r,
		validate: validator.New(),
	}
}

func (mu *IdeaUseCase) CreateIdea(c *gin.Context, m model.Idea) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.CreateIdea(m)
}

func (mu *IdeaUseCase) UpdateIdea(c *gin.Context, m model.Idea) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.UpdateIdea(m)
}