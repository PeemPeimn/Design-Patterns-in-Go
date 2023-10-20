package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBridge(t *testing.T) {
	tests := map[string]struct {
		device Device
	}{
		"Tv":    {&Tv{}},
		"Radio": {&Radio{}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			remote := &AdvancedRemote{}
			remote.SetDevice(test.device)

			remote.TogglePower()
			assert.Equal(t, true, remote.device.IsEnabled())

			remote.TogglePower()
			assert.Equal(t, false, remote.device.IsEnabled())

			remote.VolumeUp()
			assert.Equal(t, 1, remote.device.GetVolume())

			remote.VolumeDown()
			assert.Equal(t, 0, remote.device.GetVolume())

			remote.ChannelUp()
			assert.Equal(t, 1, remote.device.GetChannel())

			remote.ChannelDown()
			assert.Equal(t, 0, remote.device.GetChannel())

			remote.VolumeUp()
			remote.VolumeUp()
			remote.VolumeUp()
			remote.VolumeUp()
			remote.Mute()
			assert.Equal(t, 0, remote.device.GetVolume())

		})
	}
}
