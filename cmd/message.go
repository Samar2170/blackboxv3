package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/internal/models"
	"blackboxv3/internal/services"
	"blackboxv3/pkg/utils"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GrpcServer) SendMessage(ctx context.Context, in *blackboxv3.SendMessageRequest) (*blackboxv3.SendMessageResponse, error) {
	var err error
	_, err = utils.ValidateToken(in.Token)
	if err != nil {
		return &blackboxv3.SendMessageResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	err = services.SendMessage(uint(in.ChannelId), in.Msg)
	if err != nil {
		return &blackboxv3.SendMessageResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	return &blackboxv3.SendMessageResponse{
		Success: true,
		Msg:     "Message sent successfully",
	}, nil
}

func (s *GrpcServer) GetMessages(ctx context.Context, in *blackboxv3.GetMessagesRequest) (*blackboxv3.GetMessagesResponse, error) {
	var err error
	_, err = utils.ValidateToken(in.Token)
	if err != nil {
		return &blackboxv3.GetMessagesResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	msgs, err := services.GetChannelMessages(uint(in.ChannelId))
	if err != nil {
		return &blackboxv3.GetMessagesResponse{
			Success: false,
			Msg:     err.Error(),
		}, nil
	}
	resp := &blackboxv3.GetMessagesResponse{
		Success: true,
		Msg:     "Messages retrieved successfully",
	}
	for _, msg := range msgs {
		tmpMsg := msg.(models.TextMessage)
		resp.Messages = append(resp.Messages, &blackboxv3.TextMessage{
			Id:     int32(tmpMsg.ID),
			Msg:    tmpMsg.Text,
			SentAt: timestamppb.New(tmpMsg.CreatedAt),
		})
	}
	return resp, nil
}

// func (s *GrpcServer) SendMediaMessage(ctx context.Context, in *blackboxv3.SendMediaMessageRequest) (*blackboxv3.SendMediaMessageResponse, error) {
// 	var err error
// 	_, err = utils.ValidateToken(in.GetToken())
// 	if err != nil {
// 		return &blackboxv3.SendMediaMessageResponse{
// 			Success: false,
// 			Msg:     err.Error(),
// 		}, nil
// 	}
// err = services.SendMediaMessage(uint(in.ChannelId), in.GetFile())
// if err != nil {
// 	return &blackboxv3.SendMediaMessageResponse{
// 		Success: false,
// 		Msg:     err.Error(),
// 	}, nil
// }
// 	return &blackboxv3.SendMediaMessageResponse{
// 		Success: true,
// 		Msg:     "Message sent successfully",
// 	}, nil
// }
