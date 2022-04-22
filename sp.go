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
// Field Special Purpose dataField
type SpecialPurpose struct {
	Base
	Len  uint8
	Data []byte
}

func NewSpecialPurpose(field uap.IDataField) Item {
	f := &SpecialPurpose{}
	f.Base.NewBase(field)
	return f
}

func (sp *SpecialPurpose) Reader(rb *bytes.Reader) error {
	var err error

	sp.Len, err = rb.ReadByte()
	if err != nil {
		return err
	}

	sp.Data = make([]byte, sp.Len-1)
	err = binary.Read(rb, binary.BigEndian, &sp.Data)
	if err != nil {
		sp.Data = nil
		return err
	}

	return err
}

// Payload returns this dataField as bytes.
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

	buf.WriteString(sp.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(tmp))
	return buf.String()
}

