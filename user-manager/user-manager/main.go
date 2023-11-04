package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/user", getUsers)
	router.GET("/user/:email", userByEmail)
	//router.PUT("/user", createUser) //used to edit
	router.POST("/user", createUser)
	//router.PATCH("/edit", manageUser)
	router.Run("localhost:3000")
}
