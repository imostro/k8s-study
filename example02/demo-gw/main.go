package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	demo_api "k8s-study-example02/api"
	"log"
	"net/http"
	"time"
)

func main() {
	path := flag.String("c", "/data/conf/app.conf", "config path")
	flag.Parse()

	log.Printf("config path:%v \n", *path)
	config, err := NewConfig(*path)
	if err != nil {
		log.Println("use default config")
		config = Default()
	}
	dial, err := grpc.Dial("api-service", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := demo_api.NewDemoApiClient(dial)

	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		t := time.Now().Format(time.Layout)
		context.String(http.StatusOK, "pong ts:%v\n", t)
	})
	r.GET("/hello", func(context *gin.Context) {
		name := context.Query("name")
		reply, err := client.SayHello(context, &demo_api.HelloRequest{Name: name})
		if err != nil {
			log.Printf("greeting err:%v", err)
			context.String(http.StatusInternalServerError, "err:%v", err)
			return
		}
		context.String(http.StatusOK, reply.Message)
	})

	r.GET("/greeting", func(context *gin.Context) {
		name := context.Query("name")
		greeting, err := client.Greeting(context, &demo_api.GreetingReq{Name: name})
		if err != nil {
			log.Printf("greeting err:%v", err)
			context.String(http.StatusInternalServerError, "err:%v", err)
			return
		}
		context.String(http.StatusOK, greeting.Greeting)
	})

	if err := r.Run(config.Server.ServerAddr); err != nil {
		panic(err)
	}
}
