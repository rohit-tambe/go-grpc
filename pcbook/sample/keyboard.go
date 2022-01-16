package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/rohit-tambe/go-grpc/pcbook/pb"
	uuid "github.com/satori/go.uuid"
)

func NewKeyBoard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  RandomKeyBoardLayout(),
		Backlit: RandomBool(),
	}
	return keyboard
}

func NewCpu() *pb.CPU {
	brand := RandomCPUBrand()
	name := RandomCPUName(brand)
	numberCores := RandomInt(2, 8)
	numberOfThread := RandomInt(numberCores, 12)
	minGhz := RandomFloat64(2.0, 3.05)
	maxGhz := RandomFloat64(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberOfThread),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func NewGPU() *pb.GPU {
	brand := RandomGPUBrand()
	name := RandomCPUName(brand)
	minGhz := RandomFloat64(1.0, 2.05)
	maxGhz := RandomFloat64(minGhz, 9.0)
	memory := &pb.Memory{Value: uint64(RandomInt(2, 6)),
		Unit: pb.Memory_GEGABYTE,
	}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(RandomInt(2, 64)),
		Unit:  pb.Memory_GEGABYTE,
	}
	return ram
}
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SDD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(128, 1024)),
			Unit:  pb.Memory_GEGABYTE,
		},
	}
	return ssd
}
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(RandomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return hdd
}
func NewScreen() *pb.Screen {
	height := RandomInt(1080, 4320)
	width := height * 16 / 9
	screen := &pb.Screen{
		SizeInch: RandomFloat32(13, 72),
		Resolution: &pb.Screen_Resolution{
			Height: uint32(height),
			Width:  uint32(width),
		},
		Panel:      RandomScreenPanel(),
		Multitouch: RandomBool(),
	}
	return screen
}

func NewLapTop() *pb.Laptop {
	brand := RandomLaptopBrand()
	name := RandomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:           uuid.NewV1().String(),
		Brand:        brand,
		Name:         name,
		Cpu:          NewCpu(),
		Ram:          NewRAM(),
		Gpus:         []*pb.GPU{NewGPU()},
		Storeges:     []*pb.Storage{NewHDD(), NewSSD()},
		Screen:       NewScreen(),
		Keyboard:     NewKeyBoard(),
		Weight:       &pb.Laptop_WeightKg{WeightKg: RandomFloat64(1.0, 3.0)},
		PriceUsd:     RandomFloat64(1500, 3000),
		RealeaseYear: uint32(RandomInt(2015, 2019)),
		UpdatedAt:    ptypes.TimestampNow(),
	}
	return laptop
}
