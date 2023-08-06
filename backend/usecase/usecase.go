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
	UpdateTheme(c *gin.Context, m model.Theme) error
	ShowTheme(c *gin.Context, uid string, themeId uint) (model.Theme, error)
	DeleteTheme(c *gin.Context, m model.Theme) error
}

type IIdeaUseCase interface {
	CreateIdea(c *gin.Context, m model.Idea) error
	UpdateIdea(c *gin.Context, m model.Idea) error
	DeleteIdea(c *gin.Context, m model.Idea) error
}

type IElementUseCase interface {
	CreateElement(c *gin.Context, m model.Element) error
	UpdateElement(c *gin.Context, m model.Element) error
	DeleteElement(c *gin.Context, m model.Element) error
}
