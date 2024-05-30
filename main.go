package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router:=gin.Default()

	router.GET("/api", func(c *gin.Context){
		c.IndentedJSON(http.StatusOK,"welcome to go api")
	})

	// server := &http.Server{
	// 	Addr:           ":8888",

	//  }

	 router.Run("localhost:8000")
}