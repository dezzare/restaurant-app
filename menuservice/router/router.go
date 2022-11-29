package router

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Setup Router
func SetupRouter() (*gin.Engine, error) {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "pong")
  })
  return r, nil
}

