package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// Explicit length Data Fields shall start with a one-octet length indicator giving
// the total dataField length in octets including the length indicator itself.
type Explicit struct {
	Base
	Len  uint8
	Data []byte
}

func NewExplicit(field Item) Item {
	f := &Explicit{}
	f.Base.NewBase(field)
	return f
}

func (e Explicit) GetSize() SizeField {
	return SizeField{} // not used, it's for implement Item interface
}
func (e Explicit) GetCompound() []Item {
	return nil // not used, it's for implement Item interface
}

// Reader extracts a number of bytes define by the first byte.
func (e *Explicit) Reader(rb *bytes.Reader) error {
	var err error
	e.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	e.Data = make([]byte, e.Len-1) // tmp is for if err case then e.Data = nil
	err = binary.Read(rb, binary.BigEndian, &e.Data)
	if err != nil {
		e.Data = nil
		return err
	}

	return err
}

// Payload returns this dataField as bytes.
func (e Explicit) Payload() []byte {
	var p []byte
	p = append(p, e.Len)
	p = append(p, e.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (e Explicit) String() string {
	var buf bytes.Buffer
	buf.Reset()

	tmp := []byte{e.Len}
	tmp = append(tmp, e.Data...)

	buf.WriteString(e.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}
