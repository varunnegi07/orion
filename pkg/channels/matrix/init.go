package matrix

import (
	"path/filepath"

	"github.com/varunnegi07/orion/pkg/bus"
	"github.com/varunnegi07/orion/pkg/channels"
	"github.com/varunnegi07/orion/pkg/config"
)

func init() {
	channels.RegisterFactory(
		config.ChannelMatrix,
		func(channelName, channelType string, cfg *config.Config, b *bus.MessageBus) (channels.Channel, error) {
			bc := cfg.Channels[channelName]
			decoded, err := bc.GetDecoded()
			if err != nil {
				return nil, err
			}
			c, ok := decoded.(*config.MatrixSettings)
			if !ok {
				return nil, channels.ErrSendFailed
			}
			cryptoDatabasePath := c.CryptoDatabasePath
			if cryptoDatabasePath == "" {
				cryptoDatabasePath = filepath.Join(cfg.WorkspacePath(), "matrix")
			}
			ch, err := NewMatrixChannel(bc, c, b, cryptoDatabasePath)
			if err != nil {
				return nil, err
			}
			if channelName != config.ChannelMatrix {
				ch.SetName(channelName)
			}
			return ch, nil
		},
	)
}
