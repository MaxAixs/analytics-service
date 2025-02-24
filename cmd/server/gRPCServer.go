package server

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	GrpcServer *grpc.Server
	port       string
}

func NewGRPCServer(port string) *GRPCServer {
	return &GRPCServer{
		GrpcServer: grpc.NewServer(),
		port:       port,
	}
}

func (s *GRPCServer) RunServer() error {
	listen, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}

	logrus.Printf("listening on port %s", s.port)

	if err := s.GrpcServer.Serve(listen); err != nil {
		fmt.Printf("failed to run server %v", err)
	}

	return nil
}

func (s *GRPCServer) StopServer() {
	s.GrpcServer.GracefulStop()
	logrus.Printf("Server stopped")
}
