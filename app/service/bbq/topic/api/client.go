package api

import (
	"context"
	"go-common/library/net/rpc/warden"
	"google.golang.org/grpc"
)

// AppID unique app id for service discovery
const AppID = "bbq.service.topic"

// NewClient new member grpc client
func NewClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (TopicClient, error) {
	client := warden.NewClient(cfg, opts...)
	conn, err := client.Dial(context.Background(), "discovery://default/"+AppID)
	if err != nil {
		return nil, err
	}
	return NewTopicClient(conn), nil
}
