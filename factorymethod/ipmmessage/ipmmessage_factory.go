package ipmmessage

import (
	"encoding/json"
	"fmt"
)

const (
	TypeFirstPresentmentAR = "AR"
	TypeFirstPresentmentBR = "BR"
	TypeFirstPresentmentMX = "MX"
	TypeHeader             = "HEADER"
)

type IPMMessageFactory struct {
}

/*
	Problema: Todos los tipos de mensajes se deben poder crear a partir del mismo input.
	Es un problema si queremos incluir distintos tipos de IPM Messages (header,trailer, addendum, second presentment, etc) a la solución
*/
func (factory *IPMMessageFactory) NewIPMMessage(clearingDataCountry string) (IipmMessage, error) {
	switch clearingDataCountry {
	case TypeFirstPresentmentAR:
		return NewARFirstPresentment(), nil
	case TypeFirstPresentmentBR:
		return NewBRFirstPresentment(), nil
	case TypeFirstPresentmentMX:
		return NewMXFirstPresentment(), nil
	case TypeHeader:
		return NewHeader(), nil
	default:
		return nil, fmt.Errorf("wrong first presentment country")
	}
}

func PrintIPMMessage(message IipmMessage) {
	marshaledFP, _ := json.Marshal(message)
	fmt.Println(string(marshaledFP))
}
