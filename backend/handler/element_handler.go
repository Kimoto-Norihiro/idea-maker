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
	ThemeID, err := getThemeID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	var m model.Element
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	m.ThemeID = ThemeID

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