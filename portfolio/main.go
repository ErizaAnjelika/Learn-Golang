package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// initialize gin router
	router := gin.Default()

	// midleware logger
	router.Use(gin.Logger())

	// midleware recovery
	router.Use(gin.Recovery())

	// router definition
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello, world",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello, " + name,
		})
	})

	router.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		// ShouldBindJSON adalah function yang digunakan untuk mengambil data dari request body

		// disini anda dapat melakukan validasi login, misalnya memeriksa di databse dll
		// contoh sederhana: memerikas apakah email dan password cocok
		if loginData.Email == "example@example.com" && loginData.Password == "password123" {
			c.JSON(200, gin.H{
				"message": "Login successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid credentials",
			})
		}
	})

	// menambahkan endpoint untuk mengambuk parameter query
	router.GET("/user", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			c.JSON(400, gin.H{
				"error": "name parameter is missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello " + name + "!",
		})
	})

	// run the server
	router.Run(":8080")
}
