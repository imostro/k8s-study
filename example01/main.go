package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		t := time.Now().Format(time.Layout)
		context.String(http.StatusOK, "pong ts:%v", t)
	})
	r.GET("/hello", func(context *gin.Context) {
		name := context.Query("name")
		if name == "" {
			name = "who are you?"
		}
		context.String(http.StatusOK, "hello, %v", name)
	})

	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
