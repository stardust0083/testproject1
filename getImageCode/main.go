package main

import (
	"getImageCode/handler"
	pb "getImageCode/proto"

	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/transport/grpc/v4"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	Service = "go.micro.srv.GetImageCode"
	Version = "latest"
)

func main() {
	reg := consul.NewRegistry()
	ser := grpc.NewTransport()
	srv := micro.NewService(
		micro.Registry(reg),
		micro.Transport(ser),
		micro.Name(Service),
		micro.Version(Version),
	)
	// Register handler
	pb.RegisterGetImageCodeHandler(srv.Server(), new(handler.GetImageCode))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
