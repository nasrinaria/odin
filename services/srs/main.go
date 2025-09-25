package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello, SRS Service!")
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.Writer.Write([]byte("OK"))
	})
	router.Run(":8080")
}
