package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/usecase"
	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type ElementHandler struct {
	usecase usecase.IElementUseCase
}

func NewElementHandler(u usecase.IElementUseCase) *ElementHandler {
	return &ElementHandler{u}
}

func (ih *ElementHandler) CreateElement(c *gin.Context) {
	var m model.Element
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	m.Name = "New Element"

	if err := ih.usecase.CreateElement(c, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func (ih *ElementHandler) UpdateElement(c *gin.Context) {
	var m model.Element
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ih.usecase.UpdateElement(c, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}