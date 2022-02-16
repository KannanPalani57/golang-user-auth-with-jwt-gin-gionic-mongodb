package main

import (
	"github.com/KannanPalani57/database"
	"github.com/KannanPalani57/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	
	router := gin.New()

	router.Use(gin.Logger())

	database.ConnectDB()

	routes.UserRoute(router)	

	router.GET("/sri-rama-jeyam", func (c *gin.Context){
		c.JSON(200, gin.H{"success": "Welcome to Golang!"})
	})


	router.Run("localhost:8000")

	
}