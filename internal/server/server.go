package server

import (
	"context"

	v1 "github.com/speza/grpc-gateway-example/gen/go/api/v1"
)

// HelloWorldServer creates represents the hello world server.
type HelloWorldServer struct {
	v1.UnimplementedHelloWorldServiceServer
}

// SayHello represents the SayHello handler.
func (h *HelloWorldServer) SayHello(ctx context.Context, message *v1.Message) (*v1.Message, error) {
	return &v1.Message{Value: message.Value}, nil
}
