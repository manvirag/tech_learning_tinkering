package main

// PersonalComputer Product to build
type PersonalComputer struct {
	ramCap int
	cpu    string
	os     string
}

// PCBuilder Each builder should implement this interface
type PCBuilder interface {
	SetRAM() PCBuilder
	SetCPU() PCBuilder
	SetOS() PCBuilder
	GetPC() PersonalComputer
}

type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

func (b *HomeEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.ramCap = 4
	return b
}

func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.cpu = "i3"
	return b
}

func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.os = "Windows 7 Home Edition"
	return b
}

func (b *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}
