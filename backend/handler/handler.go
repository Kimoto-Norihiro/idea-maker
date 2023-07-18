package handler

import (
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	CreateUser(c *gin.Context)
	ShowUser(c *gin.Context)
}

type IThemeHandler interface {
	CreateTheme(c *gin.Context)
	IndexTheme(c *gin.Context)
}

type IIdeaHandler interface {
	CreateIdea(c *gin.Context)
}

type IElementHandler interface {
	CreateElement(c *gin.Context)
}