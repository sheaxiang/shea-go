package examples

import (
	"log"
	"net/http"
	"github.com/sheaxiang/shea-go"
	"time"
)

func onlyForV1() shea.HandlerFunc {
	return func(c *shea.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main()  {
	r := shea.Default()

	r.GET("/index", func(c *shea.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	v1.Use(onlyForV1())
	{
		v1.GET("/", func(c *shea.Context) {
			c.HTML(http.StatusOK, "<h1>Hello shea</h1>")
		})

		v1.GET("/hello", func(c *shea.Context) {
			names := []string{"geektutu"}
			c.String(http.StatusOK, names[100])

			// expect /hello?name=shea
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *shea.Context) {
			// expect /hello/shea
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *shea.Context) {
			c.JSON(http.StatusOK, shea.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}

