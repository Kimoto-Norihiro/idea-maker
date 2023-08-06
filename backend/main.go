package main

import (
	"github.com/gin-gonic/gin"
  
	"github.com/Kimoto-Norihiro/idea-maker/database"
	"github.com/Kimoto-Norihiro/idea-maker/handler"
	"github.com/Kimoto-Norihiro/idea-maker/middleware"
	"github.com/Kimoto-Norihiro/idea-maker/repository"
	"github.com/Kimoto-Norihiro/idea-maker/usecase"
  "github.com/Kimoto-Norihiro/idea-maker/utils"
)

func main() {
  utils.LoadEnv()
  
  r := gin.Default()
  r.Use(middleware.Cors())
  r.Use(middleware.Auth())

  const dns = "n000r111:password@tcp(localhost:3306)/idea_maker?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := database.NewMySql(dns)
  if err != nil {
    panic(err)
  }

  // user routes
  ur := repository.NewUserRepository(db)
  uu := usecase.NewUserUseCase(ur)
  uh := handler.NewUserHandler(uu)
  r.POST("/user", uh.CreateUser)
  r.GET("/user", uh.ShowUser)

  // theme routes
  tr := repository.NewThemeRepository(db)
  tu := usecase.NewThemeUseCase(tr)
  th := handler.NewThemeHandler(tu)
  r.POST("/theme", th.CreateTheme)
  r.GET("/theme/:theme_id", th.ShowTheme)
  r.PUT("/theme", th.UpdateTheme)
  r.DELETE("/theme", th.DeleteTheme)

  // idea routes
  ir := repository.NewIdeaRepository(db)
  iu := usecase.NewIdeaUseCase(ir)
  ih := handler.NewIdeaHandler(iu)
  r.POST("/idea", ih.CreateIdea)
  r.PUT("/idea", ih.UpdateIdea)
  r.DELETE("/idea", ih.DeleteIdea)

  // element routes
  er := repository.NewElementRepository(db)
  eu := usecase.NewElementUseCase(er)
  eh := handler.NewElementHandler(eu)
  r.POST("/element", eh.CreateElement)
  r.PUT("/element", eh.UpdateElement)
  r.DELETE("/element", eh.DeleteElement)

  r.Run(":8080")
}