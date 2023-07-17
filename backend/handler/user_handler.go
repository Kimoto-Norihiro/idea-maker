package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/model"
	"github.com/Kimoto-Norihiro/idea-maker/usecase"
)

type UserHandler struct {
	useCase usecase.UseCase
}

func NewUserHandler(u usecase.UseCase) *UserHandler {
	return &UserHandler{u}
}

func getUintId(c *gin.Context) (uint, error) {
	idStr := c.Param("id")
	fmt.Print(idStr)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (mh *UserHandler) CreateUser(c *gin.Context) {
	var m model.User
	uid, exist := c.Get("uid")
	if !exist {
		c.JSON(400, gin.H{"error": "uid is not exist"})
		return
	}
	err := c.BindJSON(&m)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	m.UID = uid.(string)

	err = mh.useCase.CreateUser(c, m)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func (mh *UserHandler) ShowUser(c *gin.Context) {
	uid, exist := c.Get("uid")
	if !exist {
		c.JSON(400, gin.H{
			"data":  nil,
			"error": "uid is not exist",
		})
		return
	}

	m, err := mh.useCase.ShowUser(c, uid.(string))
	if err != nil {
		c.JSON(500, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"data":  m,
		"error": nil,
	})
}

