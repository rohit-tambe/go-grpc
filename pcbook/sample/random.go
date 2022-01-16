package sample

import (
	"math/rand"
	"time"

	"github.com/rohit-tambe/go-grpc/pcbook/pb"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func RandomKeyBoardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTZ
	}
}

func RandomBool() bool {
	return rand.Intn(2) == 1
}

func RandomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}
func RandomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}
func RandomLaptopBrand() string {
	return randomStringFromSet("Apple", "Asus", "Hp", "Dell")
}
func randomStringFromSet(s ...string) string {
	m := len(s)
	if m == 0 {
		return ""
	}
	return s[rand.Intn(m)]
}

func RandomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("IntelCPU-1", "IntelCPU-2", "IntelCPU-3", "IntelCPU-4")
	} else {
		return randomStringFromSet("AMDCPU-1", "AMDCPU-2", "AMDCPU-3", "AMDCPU-4")
	}
}
func RandomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Apple-1", "Apple-2", "Apple-3", "Apple-4")
	case "Hp":
		return randomStringFromSet("HP-1", "HP-2")
	case "Dell":
		return randomStringFromSet("DELL-1", "DELL-2")
	default:
		return randomStringFromSet("ASUS", "ASUS-2")
	}
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}
func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
func RandomScreenPanel() pb.Screen_Panel {
	panelSize := rand.Intn(2)
	if panelSize == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_QLED
}
