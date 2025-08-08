package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets-compressed", "./assets-compressed")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/about", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/portfolio", func(c *gin.Context) {
		c.HTML(200, "portfolio.html", gin.H{
			"Title": "Portfolio",
			"MainTitle": "Bogrik",
			"SubTitle": "Portfolio",
		})
	})

	r.GET("/digital-painting", func(c *gin.Context) {
		c.HTML(200, "digital-painting.html", gin.H{
			"Title": "Digital Painting",
			"MainTitle": "Digital Painting",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/malutka-pig", func(c *gin.Context) {
		c.HTML(200, "malutka-pig.html", gin.H{
			"Title": "Malutka Pig",
			"MainTitle": "Malutka Pig",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/weird", func(c *gin.Context) {
		c.HTML(200, "weird.html", gin.H{
			"Title": "Weird",
			"MainTitle": "Weird",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/random", func(c *gin.Context) {
		c.HTML(200, "random.html", gin.H{
			"Title": "Random",
			"MainTitle": "Random",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/moral", func(c *gin.Context) {
		c.HTML(200, "moral.html", gin.H{
			"Title": "Moral",
			"MainTitle": "Moral",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/marvin", func(c *gin.Context) {
		c.HTML(200, "marvin.html", gin.H{
			"Title": "Marvin",
			"MainTitle": "Marvin",
			"SubTitle": "Bogrik",
		})
	})
	r.GET("/unibo", func(c *gin.Context) {
		c.HTML(200, "unibo.html", gin.H{
			"Title": "UniBo",
			"MainTitle": "UniBo",
			"SubTitle": "Bogrik",
		})
	})

	r.Run(":8085")
}
