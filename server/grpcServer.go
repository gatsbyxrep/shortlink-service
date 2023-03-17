package server

import (
	"main/links"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	server   *grpc.Server
	errCh    chan error
	listener net.Listener
}

func NewGrpcServer(service links.LinksServiceServer, port string) (GrpcServer, error) {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return GrpcServer{}, err
	}
	server := grpc.NewServer()
	links.RegisterLinksServiceServer(server, service)
	return GrpcServer{
		server:   server,
		listener: lis,
		errCh:    make(chan error, 1),
	}, nil
}

func (g GrpcServer) Start() {
	go func() {
		if err := g.server.Serve(g.listener); err != nil {
			g.errCh <- err
		}
	}()
}

func (g GrpcServer) Stop() {
	g.server.GracefulStop()
}

func (g GrpcServer) Error() chan error {
	return g.errCh
}
