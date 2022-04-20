package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// Fixed length Data Fields shall comprise a fixed number of octets.
type Fixed struct {
	Base
	Data []byte
	Size uint8
}

func NewFixed(field uap.DataField) Item {
	f := &Fixed{}
	f.Base.NewBase(field)
	f.Size = field.Fixed.Size
	return f
}

// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
//func (f *Fixed) Reader(rb *bytes.Reader, dataField uap.DataField) error {
func (f *Fixed) Reader(rb *bytes.Reader) error {
	var err error
	//f.Base.NewBase(dataField)
	//size := dataField.Fixed.Size
	f.Data = make([]byte, f.Size)
	err = binary.Read(rb, binary.BigEndian, &f.Data)
	if err != nil {
		f.Data = nil
		return err
	}
	return err
}

// Payload returns this dataField as bytes.
func (f Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (f Fixed) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(f.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(f.Data))
	return buf.String()
}

