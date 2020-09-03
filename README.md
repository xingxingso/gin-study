# gin-study

## quick

```go
import "github.com/gin-gonic/gin"
```

### `*gin.Context`

```go
// c *gin.Context

name := c.Param("name")
firstname := c.DefaultQuery("firstname", "Guest")
lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
message := c.PostForm("message")
nick := c.DefaultPostForm("nick", "anonymous")

// POST /post?ids[a]=1234&ids[b]=hello HTTP/1.1
// Content-Type: application/x-www-form-urlencoded

// names[first]=thinkerou&names[second]=tianou

ids := c.QueryMap("ids")
names := c.PostFormMap("names")

file, _ := c.FormFile("file")
// Upload the file to specific dst.
c.SaveUploadedFile(file, dst)
```

```go
import (
	"net/http"
)

c.String(http.StatusOK, message)
c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

c.JSON(200, gin.H{
	"status":  "posted",
	"message": message,
	"nick":    nick,
})
```

```go
c.FullPath() == "/user/:name/*action"
```

### route 

```go
// Creates a gin router with default middleware:
// logger and recovery (crash-free) middleware
router := gin.Default()

// By default it serves on :8080 unless a
// PORT environment variable was defined.
router.Run()
// router.Run(":3000") for a hard coded port
```

```go
router.GET("/someGet", getting)
router.POST("/somePost", posting)
router.PUT("/somePut", putting)
router.DELETE("/someDelete", deleting)
router.PATCH("/somePatch", patching)
router.HEAD("/someHead", head)
router.OPTIONS("/someOptions", options)
```

> Grouping routes

```go
// Simple group: v1
v1 := router.Group("/v1")
{
	v1.POST("/login", loginEndpoint)
	v1.POST("/submit", submitEndpoint)
	v1.POST("/read", readEndpoint)
}
```

> Blank Gin Using middleware

```go
r := gin.New()
// Global middleware
// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
// By default gin.DefaultWriter = os.Stdout
r.Use(gin.Logger())

// Recovery middleware recovers from any panics and writes a 500 if there was one.
r.Use(gin.Recovery())

// Per route middleware, you can add as many as you desire.
r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

// Authorization group
// authorized := r.Group("/", AuthRequired())
// exactly the same as:
authorized := r.Group("/")
// per group middleware! in this case we use the custom created
// AuthRequired() middleware just in the "authorized" group.
authorized.Use(AuthRequired())
{
	authorized.POST("/login", loginEndpoint)
	authorized.POST("/submit", submitEndpoint)
	authorized.POST("/read", readEndpoint)

	// nested group
	testing := authorized.Group("testing")
	testing.GET("/analytics", analyticsEndpoint)
}

// Listen and serve on 0.0.0.0:8080
r.Run(":8080")
```

> Custom Recovery behavior

```go
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
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ohai")
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```

> How to write log file

```go
// Disable Console Color, you don't need console color when writing the logs to file.
gin.DisableConsoleColor()

// Logging to a file.
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)
```

> Model binding and validation

```go
// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	router := gin.Default()

	// Example for binding JSON ({"user": "manu", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		} 
		
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}
```