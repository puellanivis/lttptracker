package sram

import (
	"encoding/json"
)

const (
	Base = 0x1E00
)

type Crystal byte
const (
	CrystalHera Crystal = 1 << (iota)
	CrystalDesert
	CrystalEastern
)

func sgn(b byte) byte {
	if b != 0 {
		return 1
	}

	return 0
}

func (c Crystal) MarshalJSON() ([]byte, error) {
	return json.Marshal([3]byte{
		sgn(byte(c & CrystalHera)),
		sgn(byte(c & CrystalDesert)),
		sgn(byte(c & CrystalEastern)),
	})
}

type Pendant byte
const (
	PendantMiseryMire Pendant = 1 << (iota)
	PendantPalaceOfDarkness
	PendantIcePalace
	PendantTurtleRock
	PendantSwampPalace
	PendantThievesTown
	PendantSkullWoods
)

func (p Pendant) MarshalJSON() ([]byte, error) {
	return json.Marshal([7]byte{
		sgn(byte(p & PendantMiseryMire)),
		sgn(byte(p & PendantPalaceOfDarkness)),
		sgn(byte(p & PendantIcePalace)),
		sgn(byte(p & PendantTurtleRock)),
		sgn(byte(p & PendantSwampPalace)),
		sgn(byte(p & PendantThievesTown)),
		sgn(byte(p & PendantSkullWoods)),
	})
}

type ZeldaData struct {
	Bow byte		// idx: 0x00
	Boomerang byte		// idx: 0x01; 0: none, 1: Blue, 2: Red
	Hookshot byte		// idx: 0x02
	Bombs byte		// idx: 0x03
	Mushroom byte		// idx: 0x04; 0: no, 1: Mushroom, 2: Powder
	FireRod byte		// idx: 0x05
	IceRod byte		// idx: 0x06
	Bombos byte		// idx: 0x07
	Ether byte		// idx: 0x08
	Quake byte		// idx: 0x09
	Lamp byte		// idx: 0x0A
	Hammer byte		// idx: 0x0B
	Shovel byte		// idx: 0x0C; 0: no, 1: Shovel, 2: Flute
	Net byte		// idx: 0x0D
	Book byte		// idx: 0x0E

	Reserved0x0f byte	`json:"-"` // idx: 0x0F

	CaneOfSomaria byte	// idx: 0x10
	CaneOfByrna byte	// idx: 0x11
	Cape byte		// idx: 0x12
	Mirror byte		// idx: 0x13
	Gloves byte		// idx: 0x14
	Boots byte		// idx: 0x15
	Flippers byte		// idx: 0x16
	MoonPearl byte		// idx: 0x17

	Reserved0x18 byte	`json:"-"` // idx: 0x18; ???

	Swords byte		// idx: 0x19
	Shield byte		// idx: 0x1A
	Tunic byte		// idx: 0x1B
	Bottles [4]byte		// idx: 0x1C-0x1F

	Reserved0x20 [0x14]byte	`json:"-"` // idx: 0x20-0x33

	Pendants Pendant	// idx: 0x34

	Reserved0x35 [0x05]byte	`json:"-"` // idx: 0x35-0x39

	Crystals Crystal	// idx: 0x3A

	Reserved0x3B [0x4A]byte	`json:"-"` // idx: 0x3B-0x84

	Agahnim byte		// idx: 0x85
}

