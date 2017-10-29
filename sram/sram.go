package sram

import (
	"encoding/json"
	"fmt"
)

const (
	Base = 0x1E00
)

type Bottle byte

const (
	BottleNone Bottle = iota
	BottleEmpty
	BottleTwo
	BottleRed
	BottleGreen
	BottleBlue
	BottleSix
	BottleBee
)

var WhatInBottle = map[Bottle]string{
	BottleNone:  "None",
	BottleEmpty: "Empty",
	BottleRed:   "Red",
	BottleGreen: "Green",
	BottleBlue:  "Blue",
	BottleBee:   "Bee",
}

func (b Bottle) MarshalJSON() ([]byte, error) {
	if content, ok := WhatInBottle[b]; ok {
		return json.Marshal(content)
	}

	return json.Marshal(fmt.Sprintf("Unknown0x%02X", b))
}

type Crystal byte

const (
	CrystalHera Crystal = 1 << (iota)
	CrystalDesert
	CrystalEastern
)

func (c Crystal) MarshalJSON() ([]byte, error) {
	m := make(map[string]byte)

	if c&CrystalHera != 0 {
		m["Hera"] = 1
	}

	if c&CrystalDesert != 0 {
		m["Desert"] = 1
	}

	if c&CrystalEastern != 0 {
		m["Eastern"] = 1
	}

	return json.Marshal(m)
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
	m := make(map[string]byte)

	if p&PendantMiseryMire != 0 {
		m["MiseryMire"] = 1
	}

	if p&PendantPalaceOfDarkness != 0 {
		m["PalaceOfDarkness"] = 1
	}

	if p&PendantIcePalace != 0 {
		m["IcePalace"] = 1
	}

	if p&PendantTurtleRock != 0 {
		m["TurtleRock"] = 1
	}

	if p&PendantSwampPalace != 0 {
		m["SwampPalace"] = 1
	}

	if p&PendantThievesTown != 0 {
		m["ThievesTown"] = 1
	}

	if p&PendantSkullWoods != 0 {
		m["SkullWoods"] = 1
	}

	return json.Marshal(m)
}

type Doubles byte

const (
	DoublesFluteWorking Doubles = 1 << (iota)
	DoublesFluteFake
	DoublesShovel
	DoublesReserved
	DoublesPowder
	DoublesMushroom
	DoublesRedBoomerang
	DoublesBlueBoomerang
)

func (d Doubles) MarshalJSON() ([]byte, error) {
	m := make(map[string]byte)

	if d&DoublesFluteWorking != 0 {
		m["FluteWorking"] = 1
	}

	if d&DoublesFluteFake != 0 {
		m["FluteFake"] = 1
	}

	if d&DoublesShovel != 0 {
		m["Shovel"] = 1
	}

	if d&DoublesPowder != 0 {
		m["Powder"] = 1
	}

	if d&DoublesMushroom != 0 {
		m["Mushroom"] = 1
	}

	if d&DoublesRedBoomerang != 0 {
		m["RedBoomerang"] = 1
	}

	if d&DoublesBlueBoomerang != 0 {
		m["BlueBoomerang"] = 1
	}

	return json.Marshal(m)
}

type GameStage byte

const (
	GameStageFindZelda GameStage = iota
	GameStageTheEscape
	GameStageBeAHero
)

var StageToName = map[GameStage]string{
	GameStageFindZelda: "FindZelda",
	GameStageTheEscape: "Escape",
	GameStageBeAHero: "BeAHero",
}

func (gs GameStage) MarshalJSON() ([]byte, error) {
	if name, ok := StageToName[gs]; ok {
		return json.Marshal(name)
	}

	return json.Marshal(fmt.Sprintf("Unknown0x%02X", gs))
}

func init() {
	var m json.Marshaler

	m = Bottle(0)
	m = GameStage(0)
	m = Pendant(0)
	m = Crystal(0)
	m = Doubles(0)

	_ = m
}

type ZeldaData struct {
	Bow       byte `json:",omitempty"` // idx: 0x00
	Boomerang byte `json:",omitempty"` // idx: 0x01; 0: none, 1: Blue, 2: Red
	Hookshot  byte `json:",omitempty"` // idx: 0x02
	Bombs     byte `json:",omitempty"` // idx: 0x03
	Mushroom  byte `json:",omitempty"` // idx: 0x04; 0: no, 1: Mushroom, 2: Powder
	FireRod   byte `json:",omitempty"` // idx: 0x05
	IceRod    byte `json:",omitempty"` // idx: 0x06
	Bombos    byte `json:",omitempty"` // idx: 0x07
	Ether     byte `json:",omitempty"` // idx: 0x08
	Quake     byte `json:",omitempty"` // idx: 0x09
	Lamp      byte `json:",omitempty"` // idx: 0x0A
	Hammer    byte `json:",omitempty"` // idx: 0x0B
	Shovel    byte `json:",omitempty"` // idx: 0x0C; 0: no, 1: Shovel, 2: Flute
	Net       byte `json:",omitempty"` // idx: 0x0D
	Book      byte `json:",omitempty"` // idx: 0x0E

	Reserved0x0f byte `json:"-"` // idx: 0x0F

	CaneOfSomaria byte `json:",omitempty"` // idx: 0x10
	CaneOfByrna   byte `json:",omitempty"` // idx: 0x11
	Cape          byte `json:",omitempty"` // idx: 0x12
	Mirror        byte `json:",omitempty"` // idx: 0x13
	Gloves        byte `json:",omitempty"` // idx: 0x14
	Boots         byte `json:",omitempty"` // idx: 0x15
	Flippers      byte `json:",omitempty"` // idx: 0x16
	MoonPearl     byte `json:",omitempty"` // idx: 0x17

	Reserved0x18 byte `json:""` // idx: 0x18; ???

	Swords  byte      `json:",omitempty"` // idx: 0x19
	Shield  byte      `json:",omitempty"` // idx: 0x1A
	Tunic   byte      // idx: 0x1B
	Bottles [4]Bottle `json:",omitempty"` // idx: 0x1C-0x1F

	Reserved0x20 [0x0F]byte `json:""` // idx: 0x20-0x28
	// 20: low-byte rupees
	// 21: high-byte rupees
	// 22: low-byte ruppees ?
	// 23: high-byte ruppees
	// …

	HeartQuarters byte `json:",omitempty"` // idx:0x29
	ExtraBombs byte `json:",omitempty"` // idx:0x30
	ExtraArrows byte `json:",omitempty"` // idx: 0x31

	Reserved0x32 [0x02]byte `json:""` // idx: 0x32-0x33

	Pendants Pendant // idx: 0x34

	Reserved0x35 [0x05]byte `json:""` // idx: 0x35-0x39
	// 35: just got more bombs
	// 36: just got more arrows

	Crystals Crystal // idx: 0x3A

	Reserved0x3B [0x4A]byte `json:""` // idx: 0x3B-0x84
	// 3B: 1/2 magic?

	GameStage GameStage // idx: 0x85

	Reserved0x86 [0x4C]byte `json:""` // idx: 0x86-0xD1
	// BF: 2→3 at half magic

	Doubles Doubles `json:",omitempty"` // idx: 0xD2
}
