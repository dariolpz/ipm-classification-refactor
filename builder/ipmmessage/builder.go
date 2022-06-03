package ipmmessage

import "errors"

const (
	FirstPresentmentARBuilderType = "firstpresentmentAR"
	FirstPresentmentBRBuilderType = "firstpresentmentBR"
	FeeCollectionBuilderType      = "feeCollection"
	HeaderBuilderType             = "builder"
)

type iBuilder interface {
	GetIPMMessage(ipmMessageData interface{}) IPMMessage
	SetDataElementss(message *IPMMessage)
	SetPrivateDataSubelements(de48 *DE48)
}

func GetBuilder(builderType string) (iBuilder, error) {
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

func (arFPBuilder *firstPresentmentARBuilder) SetDataElementss(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "001"

}

func (arFPBuilder *firstPresentmentARBuilder) SetPrivateDataSubelements(de48 *DE48) {
	de48.PDS1002 = "01"
}

func (arFPBuilder *firstPresentmentARBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	arFPBuilder.SetDataElementss(&firstPresentment)
	arFPBuilder.SetPrivateDataSubelements(&firstPresentment.DE48)
	// somethingelse()
	return firstPresentment
}

type firstPresentmentBRBuilder struct {
	FirstPresentmentData *FirstPresentmentData
}

func (brFpBuilder *firstPresentmentBRBuilder) SetDataElementss(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "200"
	message.DE04 = "1000"
	message.DE71 = "003"

}

func (brFpBuilder *firstPresentmentBRBuilder) SetPrivateDataSubelements(de48 *DE48) {
	de48.PDS220 = "03"
}

func (brFpBuilder *firstPresentmentBRBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	brFpBuilder.SetDataElementss(&firstPresentment)
	brFpBuilder.SetPrivateDataSubelements(&firstPresentment.DE48)
	return firstPresentment
}

type feeCollectionBuilder struct {
	FeeCollectionData *FeeCollectionData
}

func (feeBuilder *feeCollectionBuilder) SetDataElementss(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "700"
	message.DE04 = "100"
	message.DE71 = "001"

}

func (feeBuilder *feeCollectionBuilder) SetPrivateDataSubelements(de48 *DE48) {
	de48.PDS0137 = "000000001"
}

func (feeBuilder *feeCollectionBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	feeBuilder.SetDataElementss(&firstPresentment)
	feeBuilder.SetPrivateDataSubelements(&firstPresentment.DE48)
	return firstPresentment
}

type headerBuilder struct {
	HeaderData *HeaderData
}

func (headerBuilder *headerBuilder) SetDataElementss(message *IPMMessage) {
	message.DE48 = DE48{}
	message.DE24 = "697"
	message.DE71 = "000"

}

func (headerBuilder *headerBuilder) SetPrivateDataSubelements(de48 *DE48) {
	de48.PDS0105 = "456123132456"
}

func (headerBuilder *headerBuilder) GetIPMMessage() IPMMessage {
	var firstPresentment IPMMessage
	headerBuilder.SetDataElementss(&firstPresentment)
	headerBuilder.SetPrivateDataSubelements(&firstPresentment.DE48)
	return firstPresentment
}

/*
Header
First presentment AR
First presentment AR
...
First presentment AR
Trailer
*/

/*
Header
First presentment BR
First presentment BR
Addendum
...
First presentment BR
Trailer
*/
