package grpcclient

import (
	blackboxv3 "blackboxv3/api/apigen"
	"context"
	"io"
	"os"

	"google.golang.org/grpc"
)

func SendMediaMessageTest() error {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		return err
	}
	batchsize := 1024 * 1024
	f, err := os.Open("__test_files/card.jpeg")
	if err != nil {
		return err
	}
	buf := make([]byte, batchsize)
	stream, err := blackboxv3.NewBlackBoxServiceClient(conn).SendMediaMessage(context.Background())
	if err != nil {
		return err
	}
	for {
		num, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		chunk := buf[:num]
		stream.Send(&blackboxv3.SendMediaMessageRequest{
			FileMetaData: &blackboxv3.FileMetaData{
				Name: "card.jpeg",
				Size: 0,
				Type: "jpeg",
			},
			Token:     "token",
			ChannelId: 1,
			Chunk:     chunk,
		})
	}
	_, err = stream.CloseAndRecv()
	if err != nil {
		return err
	}
	return nil
}
