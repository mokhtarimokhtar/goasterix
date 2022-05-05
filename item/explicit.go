package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// Explicit length Data Fields shall start with a one-octet length indicator giving
// the total dataField length in octets including the length indicator itself.
type Explicit struct {
	Base
	Len      uint8
	SubItems []SubItem
}

func (e *Explicit) Clone() DataItem {
	return &Explicit{
		Base:     e.Base,
		SubItems: e.SubItems,
	}
}

func (e Explicit) GetSubItems() []SubItem {
	return e.SubItems
}

// Reader extracts a number of bytes define by the first byte.
func (e *Explicit) Reader(rb *bytes.Reader) error {
	var err error
	e.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	e.SubItems = make([]SubItem, 0, 1)
	tmp := new(SubItem)
	tmp.Type = FromToField
	tmp.From = (e.Len - 1) * 8
	tmp.To = 1
	tmp.Data = make([]byte, e.Len-1)

	err = binary.Read(rb, binary.BigEndian, &tmp.Data)
	if err != nil {
		e.SubItems = nil
		return err
	}
	e.SubItems = append(e.SubItems, *tmp)
	return err
}

// String implements fmt.Stringer in hexadecimal
func (e Explicit) String() string {
	var buf bytes.Buffer
	tmp := []byte{e.Len}

	buf.Reset()
	buf.WriteString(e.Base.DataItemName)
	buf.WriteByte(':')

	buf.WriteByte('[')
	buf.WriteString("len:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteByte(']')

	if e.SubItems != nil {
		buf.WriteByte('[')
		buf.WriteString(e.SubItems[0].String())
		buf.WriteByte(']')
	}
	return buf.String()
}

/*
// Payload returns this dataField as bytes.
func (e Explicit) Payload() []byte {
	var p []byte
	p = append(p, e.Len)
	p = append(p, e.Data...)
	return p
}
*/
