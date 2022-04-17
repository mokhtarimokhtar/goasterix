package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)


type Fixed struct {
	MetaItem
	Data []byte
}
// Reader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
// Fixed length Data Fields shall comprise a fixed number of octets.
func (f *Fixed) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	f.MetaItem.NewMetaItem(field)

	size := field.Fixed.Size
	f.Data = make([]byte, size)
	err = binary.Read(rb, binary.BigEndian, &f.Data)
	if err != nil {
		return err
	}

	return err
}

func (f Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}

func (f Fixed) String() string {
	return hex.EncodeToString(f.Data)
}

func (f Fixed) Frn() uint8 {
	return f.MetaItem.FRN
}

