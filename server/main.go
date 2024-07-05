package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./dist/index.html")
	r.Static("/assets", "./dist/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	go checkServer()

	r.Run(":80")
}

func checkServer() {
	resp, err := http.Get("http://localhost/")
	if err!=nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("server is running")
	} else {
		panic(fmt.Sprintln("server is not running as spected", resp.StatusCode))
	}
}
