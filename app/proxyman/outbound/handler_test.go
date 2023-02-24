package outbound_test

import (
	"context"
	"testing"
	_ "unsafe"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/imannamdari/v2ray-core/v5"
	"github.com/imannamdari/v2ray-core/v5/app/policy"
	. "github.com/imannamdari/v2ray-core/v5/app/proxyman/outbound"
	"github.com/imannamdari/v2ray-core/v5/app/stats"
	"github.com/imannamdari/v2ray-core/v5/common/net"
	"github.com/imannamdari/v2ray-core/v5/common/serial"
	"github.com/imannamdari/v2ray-core/v5/features/outbound"
	"github.com/imannamdari/v2ray-core/v5/proxy/freedom"
	"github.com/imannamdari/v2ray-core/v5/transport/internet"
)

func TestInterfaces(t *testing.T) {
	_ = (outbound.Handler)(new(Handler))
	_ = (outbound.Manager)(new(Manager))
}

//go:linkname toContext github.com/imannamdari/v2ray-core/v5.toContext
func toContext(ctx context.Context, v *core.Instance) context.Context

func TestOutboundWithoutStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						InboundUplink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if ok {
		t.Errorf("Expected conn to not be StatCouterConnection")
	}
}

func TestOutboundWithStatCounter(t *testing.T) {
	config := &core.Config{
		App: []*anypb.Any{
			serial.ToTypedMessage(&stats.Config{}),
			serial.ToTypedMessage(&policy.Config{
				System: &policy.SystemPolicy{
					Stats: &policy.SystemPolicy_Stats{
						OutboundUplink:   true,
						OutboundDownlink: true,
					},
				},
			}),
		},
	}

	v, _ := core.New(config)
	v.AddFeature((outbound.Manager)(new(Manager)))
	ctx := toContext(context.Background(), v)
	h, _ := NewHandler(ctx, &core.OutboundHandlerConfig{
		Tag:           "tag",
		ProxySettings: serial.ToTypedMessage(&freedom.Config{}),
	})
	conn, _ := h.(*Handler).Dial(ctx, net.TCPDestination(net.DomainAddress("localhost"), 13146))
	_, ok := conn.(*internet.StatCouterConnection)
	if !ok {
		t.Errorf("Expected conn to be StatCouterConnection")
	}
}
