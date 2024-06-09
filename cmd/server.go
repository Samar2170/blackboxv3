package main

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/internal/services"
	"blackboxv3/pkg/utils"
	"context"
	"log"
	"net"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type grpcServer struct {
	blackboxv3.UnimplementedBlackBoxServiceServer
}

func (s *grpcServer) Signup(ctx context.Context, in *blackboxv3.SignupRequest) (*blackboxv3.SignupResponse, error) {
	err := services.Signup(in.Username, in.Pin, in.Email)
	if err != nil {
		return &blackboxv3.SignupResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	return &blackboxv3.SignupResponse{
		Success: true,
		Msg:     "Signup successful",
	}, nil
}

func (s *grpcServer) Login(ctx context.Context, in *blackboxv3.LoginRequest) (*blackboxv3.LoginResponse, error) {
	token, err := services.Login(in.Username, in.Pin)
	if err != nil {
		return &blackboxv3.LoginResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	return &blackboxv3.LoginResponse{
		Success: true,
		Token:   token,
		Msg:     "Login successful",
	}, nil
}

func (s *grpcServer) CreateChannel(ctx context.Context, in *blackboxv3.CreateChannelRequest) (*blackboxv3.CreateChannelResponse, error) {
	var err error
	_, err = utils.ValidateToken(in.Token)
	if err != nil {
		return &blackboxv3.CreateChannelResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	err = services.CreateChannel(in.Name)
	if err != nil {
		return &blackboxv3.CreateChannelResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	return &blackboxv3.CreateChannelResponse{
		Success: true,
		Msg:     "Channel created successfully",
	}, nil
}

func (s *grpcServer) GetChannels(ctx context.Context, in *blackboxv3.GetChannelsRequest) (*blackboxv3.GetChannelsResponse, error) {
	var err error
	_, err = utils.ValidateToken(in.Token)
	if err != nil {
		return &blackboxv3.GetChannelsResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	channels, err := services.GetChannels()
	if err != nil {
		return &blackboxv3.GetChannelsResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	resp := &blackboxv3.GetChannelsResponse{
		Success: true,
		Msg:     "Channels retrieved successfully",
	}
	for _, channel := range channels {
		resp.Channels = append(resp.Channels, &blackboxv3.GetChannelsResponse_Channel{
			Id:   strconv.Itoa(int(channel.ID)),
			Name: channel.Name,
		})
	}
	return resp, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	log.Println("Starting server...")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	blackboxv3.RegisterBlackBoxServiceServer(s, &grpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		log.Println("Server started")
	}
}
