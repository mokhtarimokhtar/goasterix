package item

import (
	"bytes"
	"encoding/binary"
)

// Fixed length Data Fields shall comprise a fixed number of octets.
type Fixed struct {
	Base
	Size     uint8
	SubItems []SubItem
}

func (f *Fixed) Clone() DataItem {
	return &Fixed{
		Base:     f.Base,
		Size:     f.Size,
		SubItems: f.SubItems,
	}
}

func (f Fixed) GetSubItems() []SubItem {
	return f.SubItems
}

// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
func (f *Fixed) Reader(rb *bytes.Reader) error {
	var err error
	tmp := make([]byte, f.Size)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}

	tmpSubItems := f.SubItems
	f.SubItems = make([]SubItem, 0, len(f.SubItems))
	for _, subItem := range tmpSubItems {
		sub := subItem.Clone()
		err = sub.Reader(tmp)
		if err != nil {
			f.SubItems = nil
			return err
		}
		f.SubItems = append(f.SubItems, *sub)
	}

	return err
}

// String implements fmt.Stringer in hexadecimal
func (f Fixed) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(f.Base.DataItemName)
	buf.WriteByte(':')

	for _, subItem := range f.SubItems {
		buf.WriteByte('[')
		buf.WriteString(subItem.String())
		buf.WriteByte(']')
	}

	return buf.String()
}

/*
// Payload returns this dataField as bytes.
func (f Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}
*/
