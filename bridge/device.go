package bridge

type Device interface {
	IsEnabled() bool
	Enable()
	Disable()
	GetVolume() int
	SetVolume(percent int)
	GetChannel() int
	SetChannel(channel int)
}

type Tv struct {
	enable  bool
	volume  int
	channel int
}

func (tv *Tv) IsEnabled() bool {
	return tv.enable
}

func (tv *Tv) Enable() {
	tv.enable = true
}

func (tv *Tv) Disable() {
	tv.enable = false
}

func (tv *Tv) GetVolume() int {
	return tv.volume
}

func (tv *Tv) SetVolume(volume int) {
	tv.volume = volume
}

func (tv *Tv) GetChannel() int {
	return tv.channel
}

func (tv *Tv) SetChannel(channel int) {
	tv.channel = channel
}

type Radio struct {
	enable  bool
	volume  int
	channel int
}

func (r *Radio) IsEnabled() bool {
	return r.enable
}

func (r *Radio) Enable() {
	r.enable = true
}

func (r *Radio) Disable() {
	r.enable = false
}

func (r *Radio) GetVolume() int {
	return r.volume
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
}

func (r *Radio) GetChannel() int {
	return r.channel
}

func (r *Radio) SetChannel(channel int) {
	r.channel = channel
}
