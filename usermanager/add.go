package usermanager

import (
	"context"

	"google.golang.org/grpc"
	"v2ray.com/core/app/proxyman/command"
	"v2ray.com/core/common/protocol"
	"v2ray.com/core/common/serial"
	"v2ray.com/core/proxy/vmess"
)

// AddUser is used to add vmess user on runtime, will lost after v2ray restarted
// tag should exist in inbound in config
func AddUser(conn *grpc.ClientConn, UUID string, EMAIL string, alterID uint32, tag string) error {
	c := command.NewHandlerServiceClient(conn)
	_, err := c.AlterInbound(context.Background(), &command.AlterInboundRequest{
		Tag: tag,
		Operation: serial.ToTypedMessage(&command.AddUserOperation{
			User: &protocol.User{
				Email: EMAIL,
				Account: serial.ToTypedMessage(&vmess.Account{
					Id:      UUID,
					AlterId: alterID,
				}),
			},
		}),
	})
	return err
}
