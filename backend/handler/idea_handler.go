package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/model"
	"github.com/Kimoto-Norihiro/idea-maker/usecase"
)

type IdeaHandler struct {
	usecase usecase.IIdeaUseCase
}

func NewIdeaHandler(u usecase.IIdeaUseCase) *IdeaHandler {
	return &IdeaHandler{u}
}

func getThemeID(c *gin.Context) (uint, error) {
	idStr := c.Param("theme_id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (ih *IdeaHandler) CreateIdea(c *gin.Context) {
	var m model.Idea
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	m.Name = "New Idea"

	if err := ih.usecase.CreateIdea(c, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func (ih *IdeaHandler) UpdateIdea(c *gin.Context) {
	var m model.Idea
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ih.usecase.UpdateIdea(c, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}