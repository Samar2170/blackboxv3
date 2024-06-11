package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	blackboxv3.UnimplementedBlackBoxServiceServer
}

var Server *grpc.Server

func init() {
	Server = grpc.NewServer()
}
