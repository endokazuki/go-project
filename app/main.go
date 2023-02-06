package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-project/app/calculate"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	fmt.Println(calculate.Sqrt(4))
	r := setupRouter()
	r.Run(":8888")
}
