package extension

import "github.com/imannamdari/v2ray-core/v5/features"

type SubscriptionManager interface {
	features.Feature
}

func SubscriptionManagerType() interface{} {
	return (*SubscriptionManager)(nil)
}
