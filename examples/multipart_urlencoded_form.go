package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	router.Run(":8080")
}

// go run examples/multipart_urlencoded_form.go
// POST http://localhost:8080/form_post
