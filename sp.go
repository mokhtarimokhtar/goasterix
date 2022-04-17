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

func (sp SpecialPurpose) Payload() []byte {
	var p []byte
	p = append(p, sp.Len)
	p = append(p, sp.Data...)
	return p
}

func (sp SpecialPurpose) String() string {
	tmp := []byte{sp.Len}
	return hex.EncodeToString(tmp) + hex.EncodeToString(sp.Data)
}

func (sp SpecialPurpose) Frn() uint8 {
	return sp.MetaItem.FRN
}