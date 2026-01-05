package model

type ArgonConfig struct {
	HashRow    []byte
	Salt       []byte
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLen     uint32
}
