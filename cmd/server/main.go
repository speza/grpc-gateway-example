package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	v1 "github.com/speza/grpc-gateway-example/gen/go/api/v1"
	"github.com/speza/grpc-gateway-example/internal/server"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	v1.RegisterHelloWorldServiceServer(s, &server.HelloWorldServer{})

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterHelloWorldServiceHandlerFromEndpoint(ctx, mux, ":9090", opts); err != nil {
		log.Fatal(err)
	}

	go s.Serve(l)
	go http.ListenAndServe(":8080", mux)

	<-ctx.Done()
}
