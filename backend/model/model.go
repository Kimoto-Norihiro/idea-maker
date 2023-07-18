package model

import (
  "time"

  "gorm.io/gorm"
)

type User struct {
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  UID       string         `json:"uid" gorm:"primaryKey;type:varchar(255)"`
  Name      string         `json:"name"`
  Email     string         `json:"email"`
  Themes    []Theme        `json:"themes" gorm:"foreignKey:UserUID"`
}

type Element struct {
  gorm.Model
  Name    string `json:"name"`
  ThemeID uint   `json:"theme_id"`
  Theme   Theme  `json:"theme" gorm:"foreignKey:ThemeID"`
}

type Theme struct {
  gorm.Model
  Name     string    `json:"name"`
  Ideas    []Idea    `json:"ideas" gorm:"foreignKey:ThemeID"`
  Elements []Element `json:"elements" gorm:"foreignKey:ThemeID"`
  UserUID  string    `json:"user_uid" gorm:"type:varchar(255)"`
  User     User      `json:"user" gorm:"foreignKey:UserUID"`
}

type Idea struct {
  gorm.Model
  Name    string `json:"name"`
  ThemeID uint   `json:"theme_id"`
  Theme   Theme  `json:"theme" gorm:"foreignKey:ThemeID"`
}
