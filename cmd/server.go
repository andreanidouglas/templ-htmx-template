package main

import (
	"log"

	"andreanidr.com/templ-go/components"
	"github.com/gin-gonic/gin"
)

type Hello struct {
	Name string `form:"name"`
}

type GlobalState struct {
	Count int
}

var global GlobalState

func main() {
	r := gin.Default()

	r.Static("/css", "./assets/css")
	r.Static("/js", "./assets/js")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		var hello_bind Hello
		if c.ShouldBind(&hello_bind) == nil {
			log.Println(hello_bind.Name)
		}
		component := components.Hello(hello_bind.Name)
		component.Render(c, c.Writer)
	})

	r.POST("/count", func(c *gin.Context) {
		global.Count++


		c.Header("content-type", "text/html")
		count := components.Count(global.Count)
		count.Render(c, c.Writer)
	})

	r.GET("/count", func(c *gin.Context) {
		count := components.Page(global.Count, 0)
		c.Header("content-type", "text/html")
		count.Render(c, c.Writer)

	})

	r.Run()
}
