package main

import (
	"fmt"

	bguru "github.com/dariolpz/ipm-classification-refactor/builder/guru"
	bipm "github.com/dariolpz/ipm-classification-refactor/builder/ipmmessage"
	fmguru "github.com/dariolpz/ipm-classification-refactor/factorymethod/guru"
	fmipm "github.com/dariolpz/ipm-classification-refactor/factorymethod/ipmmessage"
	sguru "github.com/dariolpz/ipm-classification-refactor/strategy/guru"
	sipm "github.com/dariolpz/ipm-classification-refactor/strategy/ipmmessage"
)

func main() {
	// factoryMethodGuru()
	// factoryMethodIPMMessage()
	// builderGuru()
	builderIPM()
	// strategyGuru()
	// strategyIPMMessage()
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
	BRFirstPresentmentBuilder, _ := bipm.GetBuilder(bipm.FirstPresentmentBRBuilderType)
	ARFirstPresentmentBuilder, _ := bipm.GetBuilder(bipm.FirstPresentmentARBuilderType)
	feeCollectionBuilder, _ := bipm.GetBuilder(bipm.FeeCollectionBuilderType)
	headerBuilder, _ := bipm.GetBuilder(bipm.HeaderBuilderType)

	director := bipm.NewDirector(BRFirstPresentmentBuilder)
	brFirstPresentment := director.Build(bipm.FirstPresentmentData{})
	fmt.Printf("IpmMessage Type: %s\n", brFirstPresentment.GetMessageType())
	brFirstPresentment.PrintIPMMessage()

	director.SetBuilder(ARFirstPresentmentBuilder)
	arFirstPresentment := director.Build(bipm.FirstPresentmentData{})
	fmt.Printf("IpmMessage Type: %s\n", arFirstPresentment.GetMessageType())
	arFirstPresentment.PrintIPMMessage()

	director.SetBuilder(headerBuilder)
	header := director.Build(bipm.HeaderData{})
	fmt.Printf("IpmMessage Type: %s\n", header.GetMessageType())
	header.PrintIPMMessage()

	director.SetBuilder(feeCollectionBuilder)
	feeCollection := director.Build(bipm.FeeCollectionData{})
	fmt.Printf("IpmMessage Type: %s\n", feeCollection.GetMessageType())
	feeCollection.PrintIPMMessage()
}

func strategyGuru() {
	lfu := &sguru.LFU{}
	cache := sguru.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &sguru.LRU{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &sguru.FIFO{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}

func strategyIPMMessage() {
	BRFPStrategy := sipm.FirstPresentmentBRStrategy{}
	ARFPStrategy := sipm.FirstPresentmentARStrategy{}
	feeCollectionStrategy := sipm.FeeCollectionStrategy{}
	headerStrategy := sipm.HeaderStrategy{}

	brFirstPresentment := sipm.NewIPMMessage(BRFPStrategy)
	fmt.Printf("IpmMessage Type: %s\n", brFirstPresentment.GetMessageType())
	brFirstPresentment.PrintIPMMessage()

	arFirstPresentment := sipm.NewIPMMessage(ARFPStrategy)
	fmt.Printf("IpmMessage Type: %s\n", arFirstPresentment.GetMessageType())
	arFirstPresentment.PrintIPMMessage()

	header := sipm.NewIPMMessage(headerStrategy)
	fmt.Printf("IpmMessage Type: %s\n", header.GetMessageType())
	header.PrintIPMMessage()

	feeCollection := sipm.NewIPMMessage(feeCollectionStrategy)
	fmt.Printf("IpmMessage Type: %s\n", feeCollection.GetMessageType())
	feeCollection.PrintIPMMessage()
}
