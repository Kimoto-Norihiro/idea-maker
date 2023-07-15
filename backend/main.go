package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
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

  r.Run(":8080")
}