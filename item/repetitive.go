package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

type Repetitive struct {
	Base
	SubItemSize uint8
	Rep         uint8
	SubItems    []SubItem
}

func (r *Repetitive) Clone() DataItem {
	return &Repetitive{
		Base:        r.Base,
		SubItemSize: r.SubItemSize,
		Rep:         r.Rep,
		SubItems:    r.SubItems,
	}
}

func (r Repetitive) GetSubItems() []SubItem {
	return r.SubItems
}

// Reader extracts data item type Repetitive(1+rep*N byte).
// The first byte is REP(factor), nb is the size of bytes to repetition.
// Repetitive Data Fields, being of a variable length, shall comprise a one-octet Field Repetition Indicator (REP)
// signalling the presence of N consecutive sub-fields each of the same pre-determined length.
func (r *Repetitive) Reader(rb *bytes.Reader) error {
	var err error

	r.Rep, err = rb.ReadByte()
	if err != nil {
		return err
	}

	tmpSubItems := r.SubItems
	r.SubItems = make([]SubItem, 0, r.Rep*r.SubItemSize)

	for i := uint8(0); i < r.Rep; i++ {
		tmp := make([]byte, r.SubItemSize)
		err = binary.Read(rb, binary.BigEndian, &tmp)
		if err != nil {
			r.SubItems = nil
			return err
		}
		for _, subItem := range tmpSubItems {
			sub := subItem.Clone()
			err = sub.Reader(tmp)
			if err != nil {
				return err
			}

			r.SubItems = append(r.SubItems, *sub)
		}
	}

	return err
}

// String implements fmt.Stringer in hexadecimal
func (r Repetitive) String() string {
	var buf bytes.Buffer
	tmp := []byte{r.Rep}

	buf.Reset()
	buf.WriteString(r.Base.DataItemName)
	buf.WriteByte(':')

	buf.WriteByte('[')
	buf.WriteString("rep:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteByte(']')
	for _, subItem := range r.SubItems {
		buf.WriteByte('[')
		buf.WriteString(subItem.String())
		buf.WriteByte(']')
	}
	return buf.String()
}

/*
// Payload returns this dataField as bytes.
func (r Repetitive) Payload() []byte {
	var p []byte
	p = append(p, r.Rep)
	p = append(p, r.Data...)
	return p
}
*/
