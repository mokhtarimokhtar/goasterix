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

/*type Item struct {
	Meta       MetaItem
	Fixed      *Fixed
	Extended   *Extended
	Explicit   *Explicit
	Repetitive *Repetitive
	Compound   *Compound
	RFS        *RandomFieldSequencing
	SP         *SpecialPurpose
}*/

/*func NewItem(field uap.DataField) *Item {
	return &Item{
		Meta: MetaItem{
			FRN:         field.FRN,
			DataItem:    field.DataItem,
			Description: field.Description,
			Type:        field.Type,
		},
	}
}*/

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
func (i *Item) String() string {
	var str string
	str = i.Meta.DataItem
	switch i.Meta.Type {
	case uap.Fixed:
		str = str + ": " + i.Fixed.String()
	case uap.Extended:
		str = str + ": " + i.Extended.String()
	case uap.Explicit:
		str = str + ": " + i.Explicit.String()
	case uap.Repetitive:
		str = str + ": " + i.Repetitive.String()
	case uap.Compound:
		str = str + ": " + i.Compound.String()
	}
	return str
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