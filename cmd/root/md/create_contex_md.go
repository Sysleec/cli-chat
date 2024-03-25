package md

import (
	"context"

	"google.golang.org/grpc/metadata"
)

func CreateNewAuthMd(token string) context.Context {
	ctx := context.Background()
	md := metadata.New(map[string]string{"authorization": "Bearer " + token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}
