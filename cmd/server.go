package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/internal/services"
	"context"
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
	err := services.CreateChannel(in.Name)
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
