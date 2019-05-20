package main 

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("你好，世界！")

	r:= gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // Default Port 8080
}