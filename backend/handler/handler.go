package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	CreateUser(c *gin.Context)
	ShowUser(c *gin.Context)
}

type ThemeHandlerInterface interface {
	CreateTheme(c *gin.Context)
	IndexTheme(c *gin.Context)
}