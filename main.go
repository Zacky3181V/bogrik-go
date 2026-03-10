package main

import (
	"github.com/Zacky3181V/bogrik-go/backend"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"Title": "About",
			"Scripts": []string{
				"/assets/js/parallax.js",
				"/assets/js/form-submission.js",
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
			},
		})
	})

	r.GET("/portfolio", func(c *gin.Context) {
		c.HTML(200, "portfolio.html", gin.H{
			"Title":     "Portfolio",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
			},
		})
	})

	r.GET("/order", func(c *gin.Context) {
		c.HTML(200, "store.html", gin.H{
			"Title":     "Order",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/form-submission.js",
				"/assets/js/store-disabled.js",
			},
		})
	})

	r.POST("/order", backend.HandleForm)

	r.GET("/digital-painting", func(c *gin.Context) {
		c.HTML(200, "digital-painting.html", gin.H{
			"Title":     "Digital Painting",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
			
		})
	})
	r.GET("/malutka-pig", func(c *gin.Context) {
		c.HTML(200, "malutka-pig.html", gin.H{
			"Title":     "Malutka Pig",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})
	r.GET("/weird", func(c *gin.Context) {
		c.HTML(200, "weird.html", gin.H{
			"Title":     "Weird",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})
	r.GET("/random", func(c *gin.Context) {
		c.HTML(200, "random.html", gin.H{
			"Title":     "Random",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})
	r.GET("/moral", func(c *gin.Context) {
		c.HTML(200, "moral.html", gin.H{
			"Title":     "Moral",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})
	r.GET("/marvin", func(c *gin.Context) {
		c.HTML(200, "marvin.html", gin.H{
			"Title":     "Marvin",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})
	r.GET("/unibo", func(c *gin.Context) {
		c.HTML(200, "unibo.html", gin.H{
			"Title":     "UniBo",
			"Scripts": []string{
				"/assets/js/footer-animation.js",
				"/assets/js/store-disabled.js",
				"/assets/js/image-overlay.js",
			},
		})
	})

	r.Run(":8086")
}
