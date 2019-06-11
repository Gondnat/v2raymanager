package states

import (
	"context"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	statsService "v2ray.com/core/app/stats/command"
)

// GetStatsForAll get all stats
func GetStatsForAll(conn *grpc.ClientConn) (string, error) {
	request := "pattern: \"\" reset: false"

	client := statsService.NewStatsServiceClient(conn)

	r := &statsService.QueryStatsRequest{}
	if err := proto.UnmarshalText(request, r); err != nil {
		return "", err
	}
	resp, err := client.QueryStats(context.Background(), r)
	if err != nil {
		return "", err
	}

	return proto.MarshalTextString(resp), nil
}
