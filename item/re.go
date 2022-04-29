package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// ReservedExpansion extracts returns a slice
// ref. EUROCONTROL-SPEC-0149 2.4
// 4.3.5 Non-Standard Data Fields:
// Reserved Expansion Data
// Field Special Purpose dataField
type ReservedExpansion struct {
	Base
	Len      uint8
	Data     []byte
	SubItems []SubItemBits
}

func (re *ReservedExpansion) Clone() DataItem {
	return &ReservedExpansion{
		Base:     re.Base,
		Len:      re.Len,
		SubItems: re.SubItems,
	}
}

func (re *ReservedExpansion) Reader(rb *bytes.Reader) error {
	var err error

	re.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	re.Data = make([]byte, re.Len-1)
	err = binary.Read(rb, binary.BigEndian, &re.Data)
	if err != nil {
		re.Data = nil
		return err
	}

	return err
}

// Payload returns this dataField as bytes.
func (re ReservedExpansion) Payload() []byte {
	var p []byte
	p = append(p, re.Len)
	p = append(p, re.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (re ReservedExpansion) String() string {
	var buf bytes.Buffer
	buf.Reset()

	tmp := []byte{re.Len}
	tmp = append(tmp, re.Data...)

	buf.WriteString(re.Base.DataItemName)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}
