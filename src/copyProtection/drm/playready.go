package drm

type PlayReady struct {
	wrmHeader []byte
	pro       []byte
}

func NewPlayReady() *PlayReady {
	return &PlayReady{}
}

func (p *PlayReady) GenerateContentKey() {
}

func (p *PlayReady) GenerateWRMHeader(kid string, contentID string) {
	var b []byte
	p.wrmHeader = b
}

func (p *PlayReady) ComputeChecksum() {
}

func (p *PlayReady) GeneratePlayReadyObject() {
	var b []byte
	b = p.wrmHeader
	p.pro = b
}
