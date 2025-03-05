package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pranav514/curd-operations-golang/initializers"
)
func init(){
	initializers.LoadEnv()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := os.Getenv("PORT")
	fmt.Println(port)
	r.Run(port) 
}