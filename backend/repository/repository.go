package repository

import (
	"github.com/Kimoto-Norihiro/idea-maker/model"
)

type Repository interface {
	CreateUser(m model.User) error
	ShowUser(uid string) (model.User, error)
}
