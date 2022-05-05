package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// SpecialPurpose or Reserved extracts returns a slice
// ref. EUROCONTROL-SPEC-0149 2.4
// 4.3.5 Non-Standard Data Fields:
// Reserved Expansion Data
// Field Special Purpose dataField
type SpecialPurpose struct {
	Base
	Len      uint8
	SubItems []SubItem
}

func (sp *SpecialPurpose) Clone() DataItem {
	return &SpecialPurpose{
		Base:     sp.Base,
		Len:      sp.Len,
		SubItems: sp.SubItems,
	}
}

func (sp SpecialPurpose) GetSubItems() []SubItem {
	return sp.SubItems
}

func (sp *SpecialPurpose) Reader(rb *bytes.Reader) error {
	var err error

	sp.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	sp.SubItems = make([]SubItem, 0, 1)
	tmp := new(SubItem)
	tmp.Name = "SP"
	tmp.Type = FromToField
	tmp.From = (sp.Len - 1) * 8
	tmp.To = 1
	tmp.Data = make([]byte, sp.Len-1)

	err = binary.Read(rb, binary.BigEndian, &tmp.Data)
	if err != nil {
		sp.SubItems = nil
		return err
	}
	sp.SubItems = append(sp.SubItems, *tmp)

	return err
}

// String implements fmt.Stringer in hexadecimal
func (sp SpecialPurpose) String() string {
	var buf bytes.Buffer
	buf.Reset()
	tmp := []byte{sp.Len}

	buf.Reset()
	buf.WriteString(sp.Base.DataItemName)
	buf.WriteByte(':')

	buf.WriteByte('[')
	buf.WriteString("len:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteByte(']')

	if sp.SubItems != nil {
		buf.WriteByte('[')
		buf.WriteString(sp.SubItems[0].String())
		buf.WriteByte(']')
	}
	return buf.String()
}

/*
// Payload returns this dataField as bytes.
func (sp SpecialPurpose) Payload() []byte {
	var p []byte
	p = append(p, sp.Len)
	p = append(p, sp.Data...)
	return p
}
*/
