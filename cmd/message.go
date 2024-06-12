package cmd

import (
	blackboxv3 "blackboxv3/api/apigen"
	"blackboxv3/internal/models"
	"blackboxv3/internal/services"
	"blackboxv3/pkg/utils"
	"context"
	"io"
	"os"
	"strconv"

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

func (s *GrpcServer) SendMediaMessage(stream blackboxv3.BlackBoxService_SendMediaMessageServer) error {
	var newFile *os.File
	defer newFile.Close()
	var filesize uint32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&blackboxv3.SendMediaMessageResponse{
				Success: true,
				Msg:     "Media message sent successfully",
			})
		}
		if err != nil {
			return err
		}
		if req.GetFileMetaData() != nil {
			err = services.SaveMediaMessageMetadata(uint(req.GetChannelId()), req.GetFileMetaData().GetName())
			if err != nil {
				return err
			}
		}
		uploadsDir := os.Getenv("UPLOADSDIR")
		channel, err := services.GetChannelByParam("id", strconv.Itoa(int(req.GetChannelId())))
		if err != nil {
			return err
		}
		newFilePath := uploadsDir + "/" + channel.Name + "/" + req.GetFileMetaData().GetName()
		newFile, err = os.Create(newFilePath)
		if err != nil {
			return err
		}
		chunk := req.GetChunk()
		filesize += uint32(len(chunk))
		if _, err := newFile.Write(chunk); err != nil {
			return err
		}
	}
}

// func SaveFile(file []byte, fileName, channelName string) error {
// 	uploadsDir := os.Getenv("UPLOADSDIR")
// 	newFilePath := uploadsDir + "/" + channelName + "/" + fileName
// 	newFile, err := os.Create(newFilePath)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = newFile.Write(file)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
