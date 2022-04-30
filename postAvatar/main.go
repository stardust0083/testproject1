package main

import (
	"postAvatar/handler"
	pb "postAvatar/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
)

var (
	service = "go.micro.srv.PostAvatar"
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
	pb.RegisterPostAvatarHandler(srv.Server(), new(handler.PostAvatar))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
