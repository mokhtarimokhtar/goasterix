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
	SubItems []SubItem
}

func (re *ReservedExpansion) Clone() DataItem {
	return &ReservedExpansion{
		Base:     re.Base,
		Len:      re.Len,
		SubItems: re.SubItems,
	}
}

func (re ReservedExpansion) GetSubItems() []SubItem {
	return re.SubItems
}

func (re *ReservedExpansion) Reader(rb *bytes.Reader) error {
	var err error

	re.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	re.SubItems = make([]SubItem, 0, 1)
	tmp := new(SubItem)
	tmp.Name = "RE"
	tmp.Type = FromToField
	tmp.From = (re.Len - 1) * 8
	tmp.To = 1
	tmp.Data = make([]byte, re.Len-1)

	err = binary.Read(rb, binary.BigEndian, &tmp.Data)
	if err != nil {
		re.SubItems = nil
		return err
	}
	re.SubItems = append(re.SubItems, *tmp)

	return err
}

// String implements fmt.Stringer in hexadecimal
func (re ReservedExpansion) String() string {
	var buf bytes.Buffer
	tmp := []byte{re.Len}

	buf.Reset()
	buf.WriteString(re.Base.DataItemName)
	buf.WriteByte(':')

	buf.WriteByte('[')
	buf.WriteString("len:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteByte(']')

	if re.SubItems != nil {
		buf.WriteByte('[')
		buf.WriteString(re.SubItems[0].String())
		buf.WriteByte(']')
	}

	return buf.String()
}

/*
// Payload returns this dataField as bytes.
func (re ReservedExpansion) Payload() []byte {
	var p []byte
	p = append(p, re.Len)
	p = append(p, re.Data...)
	return p
}
*/
