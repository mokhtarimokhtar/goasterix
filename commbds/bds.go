package commbds

import (
	"encoding/hex"
	. "github.com/mokhtarimokhtar/goasterix/commbds/bdscode"
	"strconv"
	"strings"
)

var (
//ErrCodeUndefined = errors.New("BDS Error: Code CodeNotProcessed")
)

// Bds Comm-B Data Selector.
// ICAO Doc 9871: Technical Provisions for Mode S Services and Extended Squitter
// The 8-bit BDS code determines the transponder register whose contents are to be transferred in
// the MB field of a Comm-B reply.
// It is expressed in two groups of 4 bits each, BDS1 (most significant 4 bits) and BDS2 (least significant 4 bits)
// Ref. Technical Provisions for Mode S Services and Extended Squitter
// MB = Message Comm-B
// TransponderRegisterNumber: BDS1 + BDS2
// Code00: Not valid
// Code40: Selected vertical intention
// Code50: Track and turn report
// Code60: Heading and speed report
// CodeNotProcessed: code bds undefined return hex string value
type Bds struct {
	TransponderRegisterNumber string  `json:"transponderRegisterNumber"`
	Code00                    *string `json:"code00,omitempty"`
	Code40                    *Code40 `json:"code40,omitempty"`
	Code50                    *Code50 `json:"code50,omitempty"`
	Code60                    *Code60 `json:"code60,omitempty"`
	CodeNotProcessed          *string `json:"codeNotProcessed,omitempty"`
}

// Decode reads a BDS fields
// data of 8 bytes = BDS fields 7 bytes + code 1 byte
func (ds *Bds) Decode(data [8]byte) (err error) {
	transponderRegisterNumber := data[7]
	// get type of BDS and convert to hex string
	ds.TransponderRegisterNumber = strconv.FormatUint(uint64(data[7]), 16)

	var tmpMBData [7]byte
	copy(tmpMBData[:], data[:7])

	switch transponderRegisterNumber {
	case 0:
		msg := "Not valid"
		ds.Code00 = &msg
	case 96:
		// 96 = 0x60
		code60 := new(Code60)
		_ = code60.Decode(tmpMBData)
		ds.Code60 = code60
	case 80:
		// 80 = 0x50
		code50 := new(Code50)
		_ = code50.Decode(tmpMBData)
		ds.Code50 = code50
	case 64:
		// 64 = 0x40
		code40 := new(Code40)
		_ = code40.Decode(tmpMBData)
		ds.Code40 = code40
	default:
		//err = ErrCodeUndefined
		valHex := strings.ToUpper(hex.EncodeToString(data[0:7]))
		ds.CodeNotProcessed = &valHex
	}

	return err
}
