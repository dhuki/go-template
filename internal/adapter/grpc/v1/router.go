package v1

import (
	"context"
	"fmt"
	"net"

	grpcHandler "github.com/dhuki/go-template/internal/adapter/grpc"
	protoHC "github.com/dhuki/go-template/internal/adapter/grpc/v1/pb"
	"google.golang.org/grpc"
)

type svc struct {
	grpcServer *grpc.Server
	addr       string
	handler    grpcHandler.Handler

	// register default unimplemented server
	protoHC.UnimplementedHealthCheckServer
}

func NewGRPCHandlerV1(handler grpcHandler.Handler, port int) *svc {
	grpcServer := grpc.NewServer()
	return &svc{handler: handler, grpcServer: grpcServer, addr: fmt.Sprintf(":%d", port)}
}

func (s *svc) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	protoHC.RegisterHealthCheckServer(s.grpcServer, s)
	return s.grpcServer.Serve(lis)
}

func (s *svc) Stop() {
	s.grpcServer.GracefulStop()
}
