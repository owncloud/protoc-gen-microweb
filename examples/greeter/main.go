package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/go-chi/chi/v5"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/web"
	"github.com/owncloud/protoc-gen-microweb/examples/greeter/proto"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	grpc := micro.NewService(
		micro.Name("go.micro.api.hello"),
	)

	proto.RegisterGreeterHandler(
		grpc.Server(),
		&Greeter{},
	)

	go func(svc micro.Service) {
		defer wg.Done()
		svc.Init()

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
			&Greeter{},
		)
	})

	go func(svc web.Service) {
		defer wg.Done()
		svc.Init()

		if err := svc.Run(); err != nil {
			log.Fatal(err)
		}
	}(http)

	wg.Wait()
}

type Greeter struct{}

func (g *Greeter) Say(ctx context.Context, in *proto.SayRequest, out *proto.SayResponse) error {
	name := "World"

	if in.Name != "" {
		name = in.Name
	}

	out.Message = fmt.Sprintf("Hello %s!", name)
	return nil
}

func (g *Greeter) SayAnything(ctx context.Context, in *empty.Empty, out *proto.SayResponse) error {
	out.Message = "Saying Anything!"
	return nil
}
