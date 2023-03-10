package udp

import (
	"github.com/imannamdari/v2ray-core/v5/common"
	"github.com/imannamdari/v2ray-core/v5/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
