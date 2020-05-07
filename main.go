package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getting(c *gin.Context) {
	fmt.Println(c.Request.Method)
	if c.Request.Method == "GET" {
		c.String(http.StatusOK, "ok")
	} else {
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Println(string(data))
		c.String(http.StatusOK, "ok")
	}
}

func main() {
	router := gin.Default()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return ""
		// 你的自定义格式
		// return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		//         param.ClientIP,
		//         param.TimeStamp.Format(time.RFC1123),
		//         param.Method,
		//         param.Path,
		//         param.Request.Proto,
		//         param.StatusCode,
		//         param.Latency,
		//         param.Request.UserAgent(),
		//         param.ErrorMessage,
		// )

	}))

	router.GET("/web", getting)
	router.POST("/web", getting)
	router.PUT("/web", getting)
	router.DELETE("/web", getting)
	router.PATCH("/web", getting)
	router.HEAD("/web", getting)
	router.OPTIONS("/web", getting)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run(":3000")
}
