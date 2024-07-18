package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

const port = 80

func main() {
	r := gin.Default()

	buildMode := os.Getenv("BUILD_MODE")
	if buildMode == "production" {

		r.StaticFile("/favicon.ico", "./dist/favicon.ico")
		r.Static("/assets", "./dist/assets")
		r.LoadHTMLGlob("./dist/*.html")

		r.GET("/*static", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})

		gin.SetMode(gin.ReleaseMode)

	} else if buildMode == "development" {
		resp, err := http.Get("http://host.docker.internal:5173/")
		if err != nil {
			panic(fmt.Sprintf("vue js server req error: %s", err))
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println("vue js server response", string(body))

		target, err := url.Parse("http://host.docker.internal:5173/")
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(target)

		r.GET("/*static", func(c *gin.Context) {
			proxy.ServeHTTP(c.Writer, c.Request)
		})

		// go checkServer()

	} else {
		panic("env BUILD_MODE need to be \"development\" or \"production\"")
	}

	r.Run(fmt.Sprintf(":%d", port))
}
