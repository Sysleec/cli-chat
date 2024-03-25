package root

import (
	"context"

	desc "github.com/Sysleec/cli-chat/pkg/chat_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func createChat(ctx context.Context, client desc.ChatV1Client) (string, error) {

	res, err := client.CreateChat(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return res.GetChat().String(), nil
}
