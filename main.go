package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello/:name", helloHandler)
	r.Run()
}

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hola %s", c.Param("name")),
	})
}
