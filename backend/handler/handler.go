package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateUser(c *gin.Context)
	ShowUser(c *gin.Context)
}