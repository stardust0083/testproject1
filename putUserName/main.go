package main

import (
	"putUserName/handler"
	pb "putUserName/proto"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	service = "putusername"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterPutUserNameHandler(srv.Server(), new(handler.PutUserName))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
