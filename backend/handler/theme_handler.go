package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/model"
	"github.com/Kimoto-Norihiro/idea-maker/usecase"
)

type ThemeHandler struct {
	useCase usecase.IThemeUseCase
}

func NewThemeHandler(u usecase.IThemeUseCase) *ThemeHandler {
	return &ThemeHandler{u}
}

func (th *ThemeHandler) CreateTheme(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid is not exist",
		})
		return
	}
	var m model.Theme
	m.UserUID = uid.(string)
	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := th.useCase.CreateTheme(c, m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": nil,
	})
}

func (th *ThemeHandler) IndexTheme(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid is not exist",
		})
		return
	}
	m, err := th.useCase.IndexTheme(c, uid.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": m,
		"error": nil,
	})
}