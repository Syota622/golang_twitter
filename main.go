package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/health_check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run() // デフォルトでは8080ポートでリッスン
}