package udp

import (
	"context"
	"io"

	"github.com/imannamdari/v2ray-core/v5/common"
	"github.com/imannamdari/v2ray-core/v5/common/buf"
	"github.com/imannamdari/v2ray-core/v5/common/net"
)

type DispatcherI interface {
	common.Closable
	Dispatch(ctx context.Context, destination net.Destination, payload *buf.Buffer)
}

var DispatcherConnectionTerminationSignalReceiverMark = "DispatcherConnectionTerminationSignalReceiverMark"

type DispatcherConnectionTerminationSignalReceiver interface {
	io.Closer
}
