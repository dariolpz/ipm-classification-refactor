package ipmmessage

type IipmMessage interface {
	AddDEs()
	ADDPDSs()
}

type ipmMessage struct {
	DE0     string `json:"de0,omitempty"`
	DE1     string `json:"de1,omitempty"`
	DE2     string `json:"de2,omitempty"`
	DE24    string `json:"de24,omitempty"`
	PDS220  string `json:"pds220,omitempty"`
	PDS1016 string `json:"pds1016,omitempty"`
	PDS9999 string `json:"pds9999,omitempty"`
}

type ARFirstPresentment struct {
	ipmMessage
}

func NewARFirstPresentment() IipmMessage {
	fp := &ARFirstPresentment{}
	fp.AddDEs()
	fp.ADDPDSs()
	return fp
}

func (a *ARFirstPresentment) AddDEs() {
	a.DE0 = "DE0"
	a.DE24 = "200"
}

func (a *ARFirstPresentment) ADDPDSs() {
	a.PDS1016 = "PDS1016"
}

type BRFirstPresentment struct {
	ipmMessage
}

func NewBRFirstPresentment() IipmMessage {
	fp := &BRFirstPresentment{}
	fp.AddDEs()
	fp.ADDPDSs()
	return fp
}

func (b *BRFirstPresentment) AddDEs() {
	b.DE0 = "DE0"
	b.DE24 = "200"
}

func (b *BRFirstPresentment) ADDPDSs() {
	b.PDS220 = "PDS220"
}

type MXFirstPresentment struct {
	ipmMessage
}

func NewMXFirstPresentment() IipmMessage {
	fp := &MXFirstPresentment{}
	fp.AddDEs()
	fp.ADDPDSs()
	return fp
}

func (m *MXFirstPresentment) AddDEs() {
	m.DE0 = "DE0"
	m.DE24 = "200"
}

func (m *MXFirstPresentment) ADDPDSs() {
	m.PDS9999 = "PDS9999"
}

type Header struct {
	ipmMessage
}

func NewHeader() IipmMessage {
	fp := &Header{}
	fp.AddDEs()
	fp.ADDPDSs()
	return fp
}

func (h *Header) AddDEs() {
	h.DE0 = "DE0"
	h.DE24 = "697"
}

func (h *Header) ADDPDSs() {
}
