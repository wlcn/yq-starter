package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wlcn/yq-starter/service/user"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// group: v1
	v1 := r.Group("/api/v1")
	user.Routers(v1.Group("/user"))
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
