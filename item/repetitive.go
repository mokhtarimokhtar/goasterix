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
	Data        []byte
	SubItems    []SubItemBits
}

func (r *Repetitive) Clone() DataItem {
	return &Repetitive{
		Base:        r.Base,
		SubItemSize: r.SubItemSize,
		Rep:         r.Rep,
		SubItems:    r.SubItems,
	}
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

	// check if they are defined
	if r.SubItems != nil {
		tmpSubItems := r.SubItems
		r.SubItems = nil
		for i := uint8(0); i < r.Rep; i++ {
			tmp := make([]byte, r.SubItemSize)
			err = binary.Read(rb, binary.BigEndian, &tmp)
			if err != nil {
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
	} else {
		r.Data = make([]byte, r.Rep*r.SubItemSize)
		err = binary.Read(rb, binary.BigEndian, &r.Data)
		if err != nil {
			r.Data = nil
			return err
		}
	}
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
	tmp := []byte{r.Rep}

	buf.Reset()
	buf.WriteString(r.Base.DataItemName)
	buf.WriteByte(':')

	if r.SubItems != nil {
		buf.WriteByte('[')
		buf.WriteString("rep:")
		buf.WriteString(hex.EncodeToString(tmp))
		buf.WriteByte(']')
		for _, subItem := range r.SubItems {
			buf.WriteByte('[')
			buf.WriteString(subItem.String())
			buf.WriteByte(']')
		}
	} else {
		tmp = append(tmp, r.Data...)
		buf.WriteByte('[')
		buf.WriteString(hex.EncodeToString(tmp))
		buf.WriteByte(']')
	}
	return buf.String()
}
