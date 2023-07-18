package middleware

import (
  "fmt"
  "strings"

  "github.com/gin-gonic/gin"

  "github.com/Kimoto-Norihiro/idea-maker/firebase"
)

func Auth() gin.HandlerFunc {
  return func(c *gin.Context) {
    // get token from header
    idToken := c.Request.Header.Get("Authorization")
    // remove Bearer
    idToken = strings.Replace(idToken, "Bearer ", "", 1)
    // fmt.Printf("idToken: %v\n", idToken)
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
      return
    }
    if token == nil {
      fmt.Errorf("token is nil")
      c.JSON(401, gin.H{
        "message": "invalid token",
      })
      return
    }
    // fmt.Printf("Verified ID token: %v\n", token)
    // set uid to context
    c.Set("uid", token.UID)

    c.Next()
  }
}