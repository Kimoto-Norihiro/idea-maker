package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type IUserUseCase interface {
	CreateUser(c *gin.Context, m model.User) error
	ShowUser(c *gin.Context, uid string) (model.User, error)
}

type IThemeUseCase interface {
	CreateTheme(c *gin.Context, m model.Theme) error
	IndexTheme(c *gin.Context, uid string) ([]model.Theme, error)
}

type IIdeaUseCase interface {
	CreateIdea(c *gin.Context, m model.Idea) error
}

type IElementUseCase interface {
	CreateElement(c *gin.Context, m model.Element) error
}
