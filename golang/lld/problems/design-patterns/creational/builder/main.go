package main

import "fmt"

func main() {
	homePCBuilder := &HomeEditionPCBuilder{}
	homePC := homePCBuilder.SetCPU().SetOS().SetRAM().SetOS().GetPC()

	fmt.Println(homePC.cpu)
	fmt.Println(homePC.os)
	fmt.Println(homePC.ramCap)

}
