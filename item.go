package goasterix

import (
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type MetaItem struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        uap.TypeField
}

type Item struct {
	Meta       MetaItem
	Fixed      *Fixed
	Extended   *Extended
	Explicit   *Explicit
	Repetitive *Repetitive
	Compound   *Compound
	RFS        *RandomFieldSequencing
	SP         *SPandREField
	RE         *SPandREField
}

func NewItem(field uap.DataField) *Item {
	return &Item{
		Meta: MetaItem{
			FRN:         field.FRN,
			DataItem:    field.DataItem,
			Description: field.Description,
			Type:        field.Type,
		},
	}
}

func (i *Item) Payload() []byte {
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

type Fixed struct {
	Data []byte
}

func (f *Fixed) Payload() []byte {
	var p []byte
	p = append(p, f.Data...)
	return p
}

func (f *Fixed) String() string {
	return hex.EncodeToString(f.Data)
}

type Extended struct {
	Primary   []byte
	Secondary []byte
}

func (e *Extended) Payload() []byte {
	var p []byte
	p = append(p, e.Primary...)
	p = append(p, e.Secondary...)
	return p
}

func (e *Extended) String() string {
	return hex.EncodeToString(e.Primary) + hex.EncodeToString(e.Secondary)
}

type Explicit struct {
	Len  uint8
	Data []byte
}

func (e *Explicit) Payload() []byte {
	var p []byte
	p = append(p, e.Len)
	p = append(p, e.Data...)
	return p
}

func (e *Explicit) String() string {
	tmp := []byte{e.Len}
	return hex.EncodeToString(tmp) + hex.EncodeToString(e.Data)
}

type Repetitive struct {
	Rep  uint8
	Data []byte
}

func (r *Repetitive) Payload() []byte {
	var p []byte
	p = append(p, r.Rep)
	p = append(p, r.Data...)
	return p
}

func (r *Repetitive) String() string {
	tmp := []byte{r.Rep}
	return hex.EncodeToString(tmp) + hex.EncodeToString(r.Data)
}

type Compound struct {
	Primary   []byte
	Secondary []Item
}

func (c *Compound) Payload() []byte {
	var p []byte
	p = append(p, c.Primary...)
	for _, item := range c.Secondary {
		tmp := item.Payload()
		p = append(p, tmp...)
	}
	return p
}

func (c Compound) String() string {
	var str string
	str = "[primary: " + hex.EncodeToString(c.Primary) + "]"
	for _, item := range c.Secondary {
		str = str + "[" + item.String() + "]"
	}
	return str
}

type RandomFieldSequencing struct {
	N        uint8
	Sequence []RandomField
}

type RandomField struct {
	FRN   uint8
	Field Item
}

// SPandREField
// The Reserved Expansion Field is used only until and including Edition 2.1 of Part 1 of the ASTERIX documentation.
// The Reserved Expansion Data field is intended to provide a mechanism to introduce  intermediate changes to a given
// category, as explained in 7.2.3. In a way similar to the Special Purpose mechanism,
type SPandREField struct {
	Len  uint8
	Data []byte
}
