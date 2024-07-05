package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("iniciando servidor")

	r := gin.Default()

	r.LoadHTMLGlob("./dist/index.html")
	r.Static("/assets", "./dist/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.Run(":8080")
}
