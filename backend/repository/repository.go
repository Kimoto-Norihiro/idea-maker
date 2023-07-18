package repository

import (
 "github.com/Kimoto-Norihiro/idea-maker/model"
)

type IUserRepository interface {
	CreateUser(m model.User) error
	ShowUser(uid string) (model.User, error)
}

type IThemeRepository interface {
	CreateTheme(m model.Theme) error
	ShowTheme(uid string, themeId uint) (model.Theme, error)
}

type IIdeaRepository interface {
	CreateIdea(m model.Idea) error
}

type IElementRepository interface {
	CreateElement(m model.Element) error
}
