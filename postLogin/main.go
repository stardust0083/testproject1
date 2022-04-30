package main

import (
	"postLogin/handler"
	pb "postLogin/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
)

var (
	service = "go.micro.srv.PostLogin"
	version = "latest"
)

func main() {
	// Create service
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterPostLoginHandler(srv.Server(), new(handler.PostLogin))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
