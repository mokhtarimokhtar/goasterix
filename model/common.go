package model

import (
	"errors"
	"github.com/mokhtarimokhtar/goasterix/item"
)

var (
	// ErrCharUnknown reports which not found equivalent International Alphabet 5 char.
	ErrCharUnknown = errors.New("char unknown")
)

func GetValueOfStruct(subItem item.SubItem, v ...string) string {
	switch subItem.Data[0] {
	case 0:
		return v[0]
	case 1:
		return v[1]
	case 2:
		return v[2]
	case 3:
		return v[3]
	case 4:
		return v[4]
	case 5:
		return v[5]
	case 6:
		return v[6]
	case 7:
		return v[7]
	}
	return ""
}

type SourceIdentifier struct {
	Sac uint8 `json:"sac" xml:"sac"`
	Sic uint8 `json:"sic" xml:"sic"`
}

// timeOfDay returns a float64 in second (1 bit = 1/128 s)
// Absolute time stamping expressed as Co-ordinated Universal Time (UTC).
// The time information, coded in three octets, shall reflect the exact time of an event,
// expressed as a number of 1/128 s elapsed since last midnight.
// The time of day value is reset to 0 each day at midnight.
func timeOfDay(data []byte) (tod float64, err error) {
	tmp := uint32(data[0])<<16 + uint32(data[1])<<8 + uint32(data[2])
	tod = float64(tmp) / 128
	return tod, nil
}

// GetModeSIdentification returns a string.
// It converts each char into ASCII char.
// Aircraft identification (in 8 characters) obtained from an aircraft equipped with a Mode S transponder.
// Ref: 5.2.24 Records Item I048/240, Aircraft Identification.
// The CALL SIGN shall consist of eight characters, which must contain only decimal digits 0-9, the capital letters A-Z,
// and – as trailing pad characters only – the “space” character.
/*
func GetModeSIdentification(data [6]byte) (string, error) {
	var s string
	var err error

	ch1 := data[0] & 0xFC >> 2
	str1, found1 := TableIA5[ch1]

	ch2 := data[0]&0x03<<4 + data[1]&0xF0>>4
	str2, found2 := TableIA5[ch2]

	ch3 := data[1]&0x0F<<2 + data[2]&0xC0>>6
	str3, found3 := TableIA5[ch3]

	ch4 := data[2] & 0x3F
	str4, found4 := TableIA5[ch4]

	ch5 := data[3] & 0xFC >> 2
	str5, found5 := TableIA5[ch5]

	ch6 := data[3]&0x03<<4 + data[4]&0xF0>>4
	str6, found6 := TableIA5[ch6]

	ch7 := data[4]&0x0F<<2 + data[5]&0xC0>>6
	str7, found7 := TableIA5[ch7]

	ch8 := data[5] & 0x3F
	str8, found8 := TableIA5[ch8]

	if !found1 || !found2 || !found3 || !found4 || !found5 || !found6 || !found7 || !found8 {
		err = ErrCharUnknown
	}

	//s = strings.TrimSpace(str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8)
	s = str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8

	return s, err
}
*/
