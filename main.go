package main

import (
	"github.com/dariolpz/ipm-classification-refactor/factorymethod/guru"
	"github.com/dariolpz/ipm-classification-refactor/factorymethod/ipmmessage"
)

func main() {
	// factoryMethodGuru()
	factoryMethodIPMMessage()
}

func factoryMethodGuru() {
	ak47, _ := guru.GetGun("ak47")
	musket, _ := guru.GetGun("musket")

	guru.PrintDetails(ak47)
	guru.PrintDetails(musket)
}

func factoryMethodIPMMessage() {
	factory := ipmmessage.IPMMessageFactory{}
	ipmFPAR, _ := factory.NewIPMMessage(ipmmessage.TypeFirstPresentmentAR)
	ipmmessage.PrintIPMMessage(ipmFPAR)
	ipmFPBR, _ := factory.NewIPMMessage(ipmmessage.TypeFirstPresentmentBR)
	ipmmessage.PrintIPMMessage(ipmFPBR)
	ipmFPMX, _ := factory.NewIPMMessage(ipmmessage.TypeFirstPresentmentMX)
	ipmmessage.PrintIPMMessage(ipmFPMX)
	ipmHeader, _ := factory.NewIPMMessage(ipmmessage.TypeHeader)
	ipmmessage.PrintIPMMessage(ipmHeader)
}
