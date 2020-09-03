package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()


	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
	//router.GET("/user/:name/:action", func(c *gin.Context) {  // no /
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	// For each matched request Context will hold the route definition
	router.POST("/user/:name/*action", func(c *gin.Context) {
		//c.FullPath() == "/user/:name/*action" // true
		if c.FullPath() == "/user/:name/*action" {
			c.String(http.StatusOK, c.Param("name") + " is "  + c.Param("action"))
		}
	})

	router.Run(":8080")
}

// go run examples/parameters_in_path.go
// GET http://localhost:8080/user/kant
// GET http://localhost:8080/user/kant/eat
// POST http://localhost:8080/user/kant/eat