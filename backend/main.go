package main

import (
	"github.com/Kimoto-Norihiro/idea-maker/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  config := cors.DefaultConfig()
  config.AllowOrigins = []string{"http://localhost:3000"}
  config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
  config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
  config.AllowCredentials = true
  r := gin.Default()
  r.Use(cors.New(config))

  r.POST("/signup", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  r.POST("/signin", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  authTest := r.Group("/auth")
  authTest.Use(middleware.Auth())
  {
    authTest.GET("/", func(c *gin.Context) {
      c.JSON(200, gin.H{
        "message": "pong",
      })
    })
  }

  r.Run(":8080")
}