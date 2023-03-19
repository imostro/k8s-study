package main

import (
	context "context"
	"flag"
	"google.golang.org/grpc"
	demo_api "k8s-study-example02/api"
	"log"
	"net"
	"os"
)

func main() {
	path := flag.String("c", "/data/conf/app.conf", "config path")
	flag.Parse()

	os.Stdout, _ = os.Open("./app.log")

	log.Printf("config path:%v \n", path)
	config, err := NewConfig(*path)
	if err != nil {
		log.Println("use default config")
		config = Default()
	}
	api := Api{GreetingContext: config.Server.Greeting}
	server := grpc.NewServer()
	server.RegisterService(&demo_api.DemoApi_ServiceDesc, api)

	listen, err := net.Listen("tcp", config.Server.ServerAddr)
	if err != nil {
		panic(err)
	}
	panic(server.Serve(listen))
}

var _ demo_api.DemoApiServer = &Api{}

type Api struct {
	demo_api.UnimplementedDemoApiServer
	GreetingContext string
}

func (a Api) SayHello(ctx context.Context, request *demo_api.HelloRequest) (*demo_api.HelloReply, error) {
	name := request.GetName()
	log.Printf("user:%v say hello", name)
	if name == "" {
		return &demo_api.HelloReply{Message: "hi, who are you?"}, nil
	}
	return &demo_api.HelloReply{Message: "hi, " + name}, nil
}

func (a Api) Greeting(ctx context.Context, req *demo_api.GreetingReq) (*demo_api.GreetingReply, error) {
	log.Printf("user:%v greeting", req.GetName())
	return &demo_api.GreetingReply{Greeting: a.GreetingContext}, nil
}
