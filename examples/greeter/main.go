package main

import (
	"log"

	"github.com/go-chi/chi"
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

	mux := chi.NewMux()

	mux.Route("/", func(r chi.Router) {
		proto.RegisterGreeterWeb(
			r,
			&proto.Greeter{},
		)
	})

	go func(svc web.Service) {
		if err := svc.Init(); err != nil {
			log.Fatal(err)
		}

		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	}(http)
}
