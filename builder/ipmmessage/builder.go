package ipmmessage

import "errors"

const (
	FirstPresentmentARBuilderType = "firstpresentmentAR"
	FirstPresentmentBRBuilderType = "firstpresentmentBR"
	FeeCollectionBuilderType      = "feeCollection"
	HeaderBuilderType             = "builder"
)

type iBuilder interface {
	GetIPMMessage() IPMMessage
	SetDEs(message *IPMMessage)
	SetPDSs(de48 *DE48)
}

func GetBuilder(builderType string, ipmMessageData interface{}) (iBuilder, error) {
	switch builderType {
	case FirstPresentmentARBuilderType:
		fpData, ok := ipmMessageData.(*FirstPresentmentData)
		if !ok {
			return nil, errors.New("wrong data type passed")
		}
		return &firstPresentmentARBuilder{FirstPresentmentData: fpData}, nil
	case FirstPresentmentBRBuilderType:
		fpData, ok := ipmMessageData.(*FirstPresentmentData)
		if !ok {
			return nil, errors.New("wrong data type passed")
		}
		return &firstPresentmentBRBuilder{FirstPresentmentData: fpData}, nil
	case FeeCollectionBuilderType:
		feeData, ok := ipmMessageData.(*FeeCollectionData)
		if !ok {
			return nil, errors.New("wrong data type passed")
		}
		return &feeCollectionBuilder{FeeCollectionData: feeData}, nil
	case HeaderBuilderType:
		headerData, ok := ipmMessageData.(*HeaderData)
		if !ok {
			return nil, errors.New("wrong data type passed")
		}
		return &headerBuilder{HeaderData: headerData}, nil
	default:
		return nil, errors.New("invalid builder type")
	}
}

// Usariamos estos datos para agregar los PDS/DEs al mensaje
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

type firstPresentmentARBuilder struct {
	FirstPresentmentData *FirstPresentmentData
}

func (arFPBuilder *firstPresentmentARBuilder) SetDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "001"

}

func (arFPBuilder *firstPresentmentARBuilder) SetPDSs(de48 *DE48) {
	de48.PDS1002 = "01"
}

func (arFPBuilder *firstPresentmentARBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	arFPBuilder.SetDEs(&firstPresentment)
	arFPBuilder.SetPDSs(&firstPresentment.DE48)
	return firstPresentment
}

type firstPresentmentBRBuilder struct {
	FirstPresentmentData *FirstPresentmentData
}

func (brFpBuilder *firstPresentmentBRBuilder) SetDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "003"

}

func (brFpBuilder *firstPresentmentBRBuilder) SetPDSs(de48 *DE48) {
	de48.PDS220 = "03"
}

func (brFpBuilder *firstPresentmentBRBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	brFpBuilder.SetDEs(&firstPresentment)
	brFpBuilder.SetPDSs(&firstPresentment.DE48)
	return firstPresentment
}

type feeCollectionBuilder struct {
	FeeCollectionData *FeeCollectionData
}

func (feeBuilder *feeCollectionBuilder) SetDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "700"
	message.DE04 = "100"
	message.DE71 = "001"

}

func (feeBuilder *feeCollectionBuilder) SetPDSs(de48 *DE48) {
	de48.PDS0137 = "000000001"
}

func (feeBuilder *feeCollectionBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	feeBuilder.SetDEs(&firstPresentment)
	feeBuilder.SetPDSs(&firstPresentment.DE48)
	return firstPresentment
}

type headerBuilder struct {
	HeaderData *HeaderData
}

func (headerBuilder *headerBuilder) SetDEs(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "697"
	message.DE71 = "000"

}

func (headerBuilder *headerBuilder) SetPDSs(de48 *DE48) {
	de48.PDS0105 = "456123132456"
}

func (headerBuilder *headerBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	headerBuilder.SetDEs(&firstPresentment)
	headerBuilder.SetPDSs(&firstPresentment.DE48)
	return firstPresentment
}
