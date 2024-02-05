package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

type DataType struct {
	TestString string `json:"test_string"`
}

func isPalindrome(c *gin.Context) {
	var dataType DataType

	err := c.ShouldBindJSON(&dataType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	str := dataType.TestString
	lastIndex := len(str) - 1
	response := gin.H{"test_string": str, "answer": "Yes"}
	for i := 0; i < lastIndex/2 && i < (lastIndex-i); i++ {
		if str[i] != str[lastIndex-i] {
			response = gin.H{"test_string": str, "answer": "No"}
		}
	}
	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	router.POST("/api/is_palindrome", isPalindrome)

	err := router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
