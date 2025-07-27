package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/portfolio", func(c *gin.Context) {
		c.HTML(200, "portfolio.html", gin.H{})
	})

	r.GET("/digital-painting", func(c *gin.Context) {
		c.HTML(200, "digital-painting.html", gin.H{})
	})
	r.GET("/malutka-pig", func(c *gin.Context) {
		c.HTML(200, "malutka-pig.html", gin.H{})
	})
	r.GET("/weird", func(c *gin.Context) {
		c.HTML(200, "weird.html", gin.H{})
	})
	r.GET("/random", func(c *gin.Context) {
		c.HTML(200, "random.html", gin.H{})
	})
	r.GET("/moral", func(c *gin.Context) {
		c.HTML(200, "moral.html", gin.H{})
	})
	r.GET("/marvin", func(c *gin.Context) {
		c.HTML(200, "marvin.html", gin.H{})
	})

	r.Run(":8085")
}
