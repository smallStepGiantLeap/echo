package handlers

import (
	"context"

	echov1 "github.com/smallStepGiantLeap/echo/client/gen/echo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EchoService implements echo.v1.EchoService. Generated once by vikrant: this file belongs to the
// service team.
type EchoService struct {
	echov1.UnimplementedEchoServiceServer
}

// NewEchoService returns the service implementation main registers.
func NewEchoService() *EchoService { return &EchoService{} }

// Echo is unary and marked idempotent in the proto, so the mesh retries it
// on UNAVAILABLE. Keep it safe to run twice.
func (s *EchoService) Echo(ctx context.Context, req *echov1.EchoRequest) (*echov1.EchoResponse, error) {
	return nil, status.Error(codes.Unimplemented, "echo.v1.EchoService/Echo is not implemented yet")
}

// Stream streams. Return when stream.Context() is done: on shutdown the platform
// cancels it and the client reconnects to another replica.
func (s *EchoService) Stream(req *echov1.StreamRequest, stream grpc.ServerStreamingServer[echov1.StreamResponse]) error {
	return status.Error(codes.Unimplemented, "echo.v1.EchoService/Stream is not implemented yet")
}
