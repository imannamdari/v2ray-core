package grpc

import (
	"github.com/imannamdari/v2ray-core/v5/common"
	"github.com/imannamdari/v2ray-core/v5/transport/internet"
)

const protocolName = "gun"

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
