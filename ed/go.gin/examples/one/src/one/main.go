package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/v1/id/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"id": c.Param("id"),
		})
	})
	r.Run()
}
