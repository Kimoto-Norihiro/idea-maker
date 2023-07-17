package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Kimoto-Norihiro/idea-maker/model"
	"github.com/Kimoto-Norihiro/idea-maker/repository"
)

type UserUseCase struct {
	repository repository.Repository
	validate   *validator.Validate
}

func NewUserUseCase(r repository.Repository) *UserUseCase {
	return &UserUseCase{
		repository: r,
		validate:   validator.New(),
	}
}

func (mu *UserUseCase) CreateUser(c *gin.Context, m model.User) error {
	if err := mu.validate.Struct(m); err != nil {
		return err
	}
	return mu.repository.CreateUser(m)
}

func (mu *UserUseCase) ShowUser(c *gin.Context, uid string) (model.User, error) {
	fmt.Printf("ShowUserUseCase")
	return mu.repository.ShowUser(uid)
}