package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Kimoto-Norihiro/idea-maker/database"
	"github.com/Kimoto-Norihiro/idea-maker/handler"
	"github.com/Kimoto-Norihiro/idea-maker/middleware"
	"github.com/Kimoto-Norihiro/idea-maker/repository"
	"github.com/Kimoto-Norihiro/idea-maker/usecase"
)

func main() {
  r := gin.Default()
  r.Use(middleware.Cors())

  v1 := r.Group("/v1")
  {
    const dns = "n000r111:password@tcp(localhost:3306)/idea_maker?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := database.NewMySql(dns)
    if err != nil {
      panic(err)
    }
    ur := repository.NewUserRepository(db)
    uu := usecase.NewUserUseCase(ur)
    uh := handler.NewUserHandler(uu)

    v1.Use(middleware.Auth())
    v1.POST("/signup", uh.CreateUser)
    v1.GET("/signin", uh.ShowUser)
  }

  r.Run(":8080")
}