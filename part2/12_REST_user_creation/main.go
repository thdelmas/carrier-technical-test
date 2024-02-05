package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type DataType struct {
	Email string `json:"email"`
}

var i int

func createUser(c *gin.Context) {
	var dataType DataType

	err := c.ShouldBindJSON(&dataType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	response := gin.H{"id": i, "email": dataType.Email}
	i += 1

	c.JSON(http.StatusOK, response)
}

func main() {
	i = 0
	router := gin.Default()

	router.POST("/api/users", createUser)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
