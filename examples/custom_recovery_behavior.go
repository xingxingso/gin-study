package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	r.GET("/panic", func(c *gin.Context) {
		// panic with a string -- the custom middleware could save this to a database or report it to the user
		panic("foo")
		// panic(true)
	})

	r.GET("/", func(c *gin.Context) {
		s := []int{1, 2, 3}
		x := 3
		c.String(http.StatusOK, fmt.Sprintf("x=%d", s[x]))

		// c.String(http.StatusOK, "ohai")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

// go run examples/custom_recovery_behavior.go
// GET http://localhost:8080/panic
// GET http://localhost:8080/