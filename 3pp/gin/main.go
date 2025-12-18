package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":10010")
}
