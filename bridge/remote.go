package bridge

type Remote struct {
	device Device
}

func (r *Remote) TogglePower() {
	if r.device.IsEnabled() {
		r.device.Disable()
		return
	}

	r.device.Enable()
}

func (r *Remote) VolumeDown() {
	v := r.device.GetVolume()
	r.device.SetVolume(v - 1)
}

func (r *Remote) VolumeUp() {
	v := r.device.GetVolume()
	r.device.SetVolume(v + 1)
}

func (r *Remote) ChannelDown() {
	c := r.device.GetChannel()
	r.device.SetChannel(c - 1)
}

func (r *Remote) ChannelUp() {
	c := r.device.GetChannel()
	r.device.SetChannel(c + 1)
}

func (r *Remote) SetDevice(device Device) {
	r.device = device
}

type AdvancedRemote struct {
	Remote
}

func (r *AdvancedRemote) Mute() {
	r.device.SetVolume(0)
}
