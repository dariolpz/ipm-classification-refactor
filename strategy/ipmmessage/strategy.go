package ipmmessage

type ipmMssageStrategy interface {
	setDEs(message *IPMMessage)
	setPDSs(de48 *DE48)
}

type FirstPresentmentData struct {
	Foo string
	Bar string
}

type FeeCollectionData struct {
	Foo string
	Bar string
}

type HeaderData struct {
	Foo string
	Bar string
}

type FirstPresentmentARStrategy struct {
}

func (arFP FirstPresentmentARStrategy) setDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "001"
}

func (arFP FirstPresentmentARStrategy) setPDSs(de48 *DE48) {
	de48.PDS1002 = "01"
}

type FirstPresentmentBRStrategy struct {
	FirstPresentmentData *FirstPresentmentData
}

func (brFpBuilder FirstPresentmentBRStrategy) setDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "003"

}

func (brFpBuilder FirstPresentmentBRStrategy) setPDSs(de48 *DE48) {
	de48.PDS220 = "03"
}

type FeeCollectionStrategy struct {
	FeeCollectionData *FeeCollectionData
}

func (feeStrategy FeeCollectionStrategy) setDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "700"
	message.DE04 = "100"
	message.DE71 = "001"

}

func (feeStrategy FeeCollectionStrategy) setPDSs(de48 *DE48) {
	de48.PDS0137 = "000000001"
}

type HeaderStrategy struct {
	HeaderData *HeaderData
}

func (headerStrategy HeaderStrategy) setDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "697"
	message.DE71 = "000"

}

func (headerStrategy HeaderStrategy) setPDSs(de48 *DE48) {
	de48.PDS0105 = "456123132456"
}
