package onebot

import (
	"github.com/varunnegi07/orion/pkg/bus"
	"github.com/varunnegi07/orion/pkg/channels"
	"github.com/varunnegi07/orion/pkg/config"
)

func init() {
	channels.RegisterFactory(
		config.ChannelOneBot,
		func(channelName, channelType string, cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
			bc := cfg.Channels[channelName]
			decoded, err := bc.GetDecoded()
			if err != nil {
				return nil, err
			}
			c, ok := decoded.(*config.OneBotSettings)
			if !ok {
				return nil, channels.ErrSendFailed
			}
			return NewOneBotChannel(bc, c, b)
		},
	)
}
