package main

import (
	"fmt"

	bguru "github.com/dariolpz/ipm-classification-refactor/builder/guru"
	bipm "github.com/dariolpz/ipm-classification-refactor/builder/ipmmessage"
	fmguru "github.com/dariolpz/ipm-classification-refactor/factorymethod/guru"
	fmipm "github.com/dariolpz/ipm-classification-refactor/factorymethod/ipmmessage"
)

func main() {
	// factoryMethodGuru()
	// factoryMethodIPMMessage()
	//builderGuru()
	builderIPM()
}

func factoryMethodGuru() {
	ak47, _ := fmguru.GetGun("ak47")
	musket, _ := fmguru.GetGun("musket")

	fmguru.PrintDetails(ak47)
	fmguru.PrintDetails(musket)
}

func factoryMethodIPMMessage() {
	factory := fmipm.IPMMessageFactory{}
	ipmFPAR, _ := factory.NewIPMMessage(fmipm.TypeFirstPresentmentAR)
	fmipm.PrintIPMMessage(ipmFPAR)
	ipmFPBR, _ := factory.NewIPMMessage(fmipm.TypeFirstPresentmentBR)
	fmipm.PrintIPMMessage(ipmFPBR)
	ipmFPMX, _ := factory.NewIPMMessage(fmipm.TypeFirstPresentmentMX)
	fmipm.PrintIPMMessage(ipmFPMX)
	ipmHeader, _ := factory.NewIPMMessage(fmipm.TypeHeader)
	fmipm.PrintIPMMessage(ipmHeader)
}

func builderGuru() {
	normalBuilder := bguru.GetBuilder("normal")
	iglooBuilder := bguru.GetBuilder("igloo")

	director := bguru.NewDirector(normalBuilder)
	normalHouse := director.BuildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.DoorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.WindowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.Floor)

	director.SetBuilder(iglooBuilder)
	iglooHouse := director.BuildHouse()

	fmt.Printf("Igloo House Door Type: %s\n", iglooHouse.DoorType)
	fmt.Printf("Igloo House Window Type: %s\n", iglooHouse.WindowType)
	fmt.Printf("Igloo House Num Floor: %d\n", iglooHouse.Floor)

}

func builderIPM() {
	BRFirstPresentmentBuilder, _ := bipm.GetBuilder(bipm.FirstPresentmentBRBuilderType, &bipm.FirstPresentmentData{})
	ARFirstPresentmentBuilder, _ := bipm.GetBuilder(bipm.FirstPresentmentARBuilderType, &bipm.FirstPresentmentData{})
	feeCollectionBuilder, _ := bipm.GetBuilder(bipm.FeeCollectionBuilderType, &bipm.FeeCollectionData{})
	headerBuilder, _ := bipm.GetBuilder(bipm.HeaderBuilderType, &bipm.HeaderData{})

	director := bipm.NewDirector(BRFirstPresentmentBuilder)
	brFirstPresentment := director.BuildHouse()
	fmt.Printf("IpmMessage Type: %s\n", brFirstPresentment.GetMessageType())
	brFirstPresentment.PrintIPMMessage()

	director.SetBuilder(ARFirstPresentmentBuilder)
	arFirstPresentment := director.BuildHouse()
	fmt.Printf("IpmMessage Type: %s\n", arFirstPresentment.GetMessageType())
	arFirstPresentment.PrintIPMMessage()

	director.SetBuilder(headerBuilder)
	header := director.BuildHouse()
	fmt.Printf("IpmMessage Type: %s\n", header.GetMessageType())
	header.PrintIPMMessage()

	director.SetBuilder(feeCollectionBuilder)
	feeCollection := director.BuildHouse()
	fmt.Printf("IpmMessage Type: %s\n", feeCollection.GetMessageType())
	feeCollection.PrintIPMMessage()
}
