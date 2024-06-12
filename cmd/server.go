package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"
	"log"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	blackboxv3.UnimplementedBlackBoxServiceServer
	logger *log.Logger
}

func NewBlackboxServer(logger *log.Logger) blackboxv3.BlackBoxServiceServer {
	return &GrpcServer{logger: logger}
}

func RegisterBlackBoxServer(registrar grpc.ServiceRegistrar, server blackboxv3.BlackBoxServiceServer) {
	blackboxv3.RegisterBlackBoxServiceServer(registrar, server)
}

// type GrpcServer struct {
// 	blackboxv3.UnimplementedBlackBoxServiceServer
// }

// var Server *grpc.Server

// func init() {
// 	Server = grpc.NewServer()
// }
