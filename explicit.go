package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// Explicit length Data Fields shall start with a one-octet length indicator giving
// the total field length in octets including the length indicator itself.
type Explicit struct {
	MetaItem
	Len  uint8
	Data []byte
}

// Reader extracts a number of bytes define by the first byte.
func (e *Explicit) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	e.MetaItem.NewMetaItem(field)

	err = binary.Read(rb, binary.BigEndian, &e.Len)
	if err != nil {
		return err
	}
	tmp := make([]byte, e.Len-1)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	e.Data = tmp

	return err
}

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

	buf.WriteString(e.MetaItem.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}

func (e Explicit) Frn() uint8 {
	return e.MetaItem.FRN
}
