package main

import (
	"time"

	"github.com/didip/tollbooth"
	ratelimit "github.com/feixiao/gin-ratelimit"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())

	lmt := tollbooth.NewLimiter(10, nil)

	r.Use(ratelimit.RateLimit(lmt))

	r.GET("/ping", func(c *gin.Context) {

		time.Sleep(200 * time.Millisecond)

		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.Run(":12345")
}
