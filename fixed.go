package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// Fixed length Data Fields shall comprise a fixed number of octets.
type Fixed struct {
	MetaItem
	Data []byte
}

// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
func (f *Fixed) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	f.MetaItem.NewMetaItem(field)

	size := field.Fixed.Size
	f.Data = make([]byte, size)
	err = binary.Read(rb, binary.BigEndian, &f.Data)
	if err != nil {
		f.Data = nil
		return err
	}
	return err
}

// Payload returns this field as bytes.
func (f Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (f Fixed) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(f.MetaItem.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(f.Data))
	return buf.String()
}

// Frn returns FRN number of field from UAP
func (f Fixed) Frn() uint8 {
	return f.MetaItem.FRN
}
