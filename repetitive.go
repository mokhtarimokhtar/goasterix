package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Repetitive struct {
	Base
	SubItemSize uint8
	Rep         uint8
	Data        []byte
}

func NewRepetitive(field uap.DataField) Item {
	f := &Repetitive{}
	f.Base.NewBase(field)
	f.SubItemSize = field.Repetitive.SubItemSize
	return f
}

// Reader extracts data item type Repetitive(1+rep*N byte).
// The first byte is REP(factor), nb is the size of bytes to repetition.
// Repetitive Data Fields, being of a variable length, shall comprise a one-octet Field Repetition Indicator (REP)
// signalling the presence of N consecutive sub-fields each of the same pre-determined length.
func (r *Repetitive) Reader(rb *bytes.Reader) error {
	var err error
	err = binary.Read(rb, binary.BigEndian, &r.Rep)
	if err != nil {
		return err
	}
	tmp := make([]byte, r.Rep*r.SubItemSize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	r.Data = tmp

	return err
}

// Payload returns this dataField as bytes.
func (r Repetitive) Payload() []byte {
	var p []byte
	p = append(p, r.Rep)
	p = append(p, r.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (r Repetitive) String() string {
	var buf bytes.Buffer
	buf.Reset()

	tmp := []byte{r.Rep}
	tmp = append(tmp, r.Data...)

	buf.WriteString(r.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}
