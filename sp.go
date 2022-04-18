package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// SpecialPurpose or Reserved extracts returns a slice
// ref. EUROCONTROL-SPEC-0149 2.4
// 4.3.5 Non-Standard Data Fields:
// Reserved Expansion Data
// Field Special Purpose field
type SpecialPurpose struct {
	MetaItem
	Len  uint8
	Data []byte
}

func (sp *SpecialPurpose) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	sp.MetaItem.NewMetaItem(field)

	err = binary.Read(rb, binary.BigEndian, &sp.Len)
	if err != nil {
		return err
	}

	tmp := make([]byte, sp.Len-1)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	sp.Data = tmp

	return err
}

// Payload returns this field as bytes.
func (sp SpecialPurpose) Payload() []byte {
	var p []byte
	p = append(p, sp.Len)
	p = append(p, sp.Data...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (sp SpecialPurpose) String() string {
	var buf bytes.Buffer
	buf.Reset()

	tmp := []byte{sp.Len}
	tmp = append(tmp, sp.Data...)

	buf.WriteString(sp.MetaItem.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}

// Frn returns FRN number of field from UAP
func (sp SpecialPurpose) Frn() uint8 {
	return sp.MetaItem.FRN
}
