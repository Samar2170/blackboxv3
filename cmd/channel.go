package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/internal/services"
	"blackboxv3/pkg/utils"
	"context"
	"strconv"
)

func (s *GrpcServer) CreateChannel(ctx context.Context, in *blackboxv3.CreateChannelRequest) (*blackboxv3.CreateChannelResponse, error) {
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

func (s *GrpcServer) GetChannels(ctx context.Context, in *blackboxv3.GetChannelsRequest) (*blackboxv3.GetChannelsResponse, error) {
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
