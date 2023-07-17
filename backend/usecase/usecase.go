package usecase

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type UseCase interface {
	CreateUser(c *gin.Context, m model.User) error
	ShowUser(c *gin.Context, uid string) (model.User, error)
}
