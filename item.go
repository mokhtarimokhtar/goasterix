package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Item interface {
	Payload() []byte
	String() string
	Reader(*bytes.Reader, uap.DataField) error
	Frn() uint8
}

/*func (i *Item) Payload() []byte {
	var p []byte
	switch i.Meta.Type {
	case uap.Fixed:
		p = i.Fixed.Payload()
	case uap.Extended:
		p = i.Extended.Payload()
	case uap.Explicit:
		p = i.Explicit.Payload()
	case uap.Repetitive:
		p = i.Repetitive.Payload()
	case uap.Compound:
		p = i.Compound.Payload()
	}
	return p
}

*/

type MetaItem struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        uap.TypeField
}

func (m *MetaItem) NewMetaItem(field uap.DataField) {
	m.FRN = field.FRN
	m.DataItem = field.DataItem
	m.Description = field.Description
	m.Type = field.Type
}
