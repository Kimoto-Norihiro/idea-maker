package middleware

import (
  "fmt"

  "github.com/gin-gonic/gin"

  "github.com/Kimoto-Norihiro/idea-maker/firebase"
)

func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println("Auth middleware")
    // get token from header
    idToken := c.Request.Header.Get("Authorization")
    // init firebase app
    firebaseApp, err := firebase.InitFirebaseApp()
    if err != nil {
      fmt.Errorf("error initializing app:", err)
      c.JSON(500, gin.H{
        "message": "Internal Server Error",
      })
    }
    // verify token
    token, err := firebaseApp.VerifyIDToken(c, idToken)
    if err != nil {
      fmt.Errorf("error verifying ID token:", err)
      c.JSON(401, gin.H{
        "message": err.Error(),
      })
    }
    if token == nil {
      fmt.Errorf("token is nil")
      c.JSON(401, gin.H{
        "message": "invalid token",
      })
    }
    // set uid to context
    c.Set("uid", token.UID)

    c.Next()
  }
}