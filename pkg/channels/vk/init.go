package vk

import (
	"github.com/varunnegi07/orion/pkg/bus"
	"github.com/varunnegi07/orion/pkg/channels"
	"github.com/varunnegi07/orion/pkg/config"
)

func init() {
	channels.RegisterFactory(
		config.ChannelVK,
		func(channelName, channelType string, cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
			bc := cfg.Channels[channelName]
			if bc == nil {
				return nil, channels.ErrSendFailed
			}
			return NewVKChannel(channelName, bc, b)
		},
	)
}
