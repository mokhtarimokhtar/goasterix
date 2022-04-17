package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Repetitive struct {
	MetaItem
	Rep  uint8
	Data []byte
}
func (r *Repetitive) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	r.MetaItem.NewMetaItem(field)

	subItemSize := field.Repetitive.SubItemSize
	err = binary.Read(rb, binary.BigEndian, &r.Rep)
	if err != nil {
		return err
	}
	tmp := make([]byte, r.Rep*subItemSize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	r.Data = tmp

	return err
}

func (r Repetitive) Payload() []byte {
	var p []byte
	p = append(p, r.Rep)
	p = append(p, r.Data...)
	return p
}

func (r Repetitive) String() string {
	tmp := []byte{r.Rep}
	return hex.EncodeToString(tmp) + hex.EncodeToString(r.Data)
}

func (r Repetitive) Frn() uint8 {
	return r.MetaItem.FRN
}