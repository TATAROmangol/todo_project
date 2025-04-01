package gv1

import (
	"auth/pkg/logger"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct{
	ctx context.Context
	cfg Config
	server *grpc.Server
}

func New(ctx context.Context, cfg Config, l *logger.Logger, service Auth) *Server{
	server := grpc.NewServer(
		grpc.UnaryInterceptor(LoggerInterceptor(ctx, l)),
	)
	Register(server, service)

	return &Server{ctx: ctx, cfg: cfg, server: server}
}

func (s *Server) Run() error{
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", s.cfg.Host, s.cfg.Port))
	if err != nil{
		return fmt.Errorf("failed create listener from grpc: %v", err)
	}

	err = s.server.Serve(lis)
	if err != nil{
		return fmt.Errorf("failed in grpc server: %v", err)
	}

	return nil
}

func (s *Server) GracefulStop() {
	s.server.GracefulStop()
}