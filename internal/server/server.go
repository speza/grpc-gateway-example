package server

import (
	"context"

	v1 "github.com/speza/grpc-gateway-test/gen/go/api/v1"
)

type HelloWorldServer struct {
	v1.UnimplementedHelloWorldServiceServer
}

func (h *HelloWorldServer) SayHello(ctx context.Context, message *v1.Message) (*v1.Message, error) {
	return &v1.Message{Value: message.Value}, nil
}
