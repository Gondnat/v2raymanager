package usermanager

import (
	"context"

	"google.golang.org/grpc"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/common/serial"
)

// RemoveUser use to remove user
// tag should exist in inbound in config
func RemoveUser(conn *grpc.ClientConn, email string, tag string) error {
	c := command.NewHandlerServiceClient(conn)
	_, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: tag,
		Operation: serial.ToTypedMessage(&command.RemoveUserOperation{
			Email: email,
		}),
	})
	return err
}
