package main

import (
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	// logLevel = int32(0)
	mutex sync.RWMutex
)

func main() {
	r := gin.Default()
	client := http.Client{}
	r.GET("/ping", func(c *gin.Context) {
		rsp, _ := client.Get("http://localhost:5001/ping")
		bd, _ := ioutil.ReadAll(rsp.Body)
		defer func() { rsp.Body.Close() }()

		c.JSON(200, gin.H{
			"message": string(bd),
		})
	})
	r.Run(":5000") // listen and serve on 0.0.0.0:8080
}
