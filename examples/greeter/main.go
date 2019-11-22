package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"github.com/webhippie/protoc-gen-microweb/examples/greeter/proto"
)

func main() {
	grpc := micro.NewService(
		micro.Name("go.micro.api.hello"),
	)

	proto.RegisterGreeterHandler(
		grpc.Server(),
		&proto.Greeter{},
	)

	go func(svc micro.Service) {
		if err := svc.Init(); err != nil {
			log.Fatal(err)
		}

		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	}(grpc)

	http := web.NewService(
		web.Name("go.micro.web.hello"),
	)

	proto.RegisterGreeterWeb(
		http,
		grpc.Name(),
		grpc.Client(),
	)

	go func(svc web.Service) {
		if err := svc.Init(); err != nil {
			log.Fatal(err)
		}

		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	}(http)
}
