package jsonv4

import "github.com/imannamdari/v2ray-core/v5/main/commands/all/api"

func init() {
	api.CmdAPI.Commands = append(api.CmdAPI.Commands,
		cmdAddInbounds,
		cmdAddOutbounds,
		cmdRemoveInbounds,
		cmdRemoveOutbounds)
}
