package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homePage(c *gin.Context) {
	c.HTML(http.StatusOK, "test.html", gin.H{
		"title": "Main website",
	})

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("content/*")
	r.GET("/", homePage)
	r.Run()
}
