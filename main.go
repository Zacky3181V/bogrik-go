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

	r.Run(":8085")
}