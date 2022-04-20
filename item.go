package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

/*type ContextType struct {
	Item Item
}
func (c *ContextType) setReader(tf dataField.TypeField) error {
	var err error
	switch tf {
	case dataField.Fixed:
		c.Item = new(Fixed)
	case dataField.Extended:
		c.Item = new(Extended)
	case dataField.Explicit:
		c.Item = new(Explicit)
	case dataField.Repetitive:
		c.Item = new(Repetitive)
	case dataField.Compound:
		c.Item = new(Compound)
	case dataField.SP, dataField.RE:
		c.Item = new(SpecialPurpose)
	case dataField.RFS:
		c.Item = new(RandomFieldSequencing)
	default:
		err = ErrDataFieldUnknown
		return err
	}
	return err
}

func (c *ContextType) Reader(rb *bytes.Reader, df dataField.dataField) error {
	err := c.Item.Reader(rb, df)
	return err
}*/

type Item interface {
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	Frn() uint8
}

func Readers(i Item, rb *bytes.Reader) error {
	err := i.Reader(rb)
	return err
}

// GetItem returns the corresponding Item type: Fixed, Extended, etc.
// GetItem is a factory function
func GetItem(df uap.DataField) (Item, error) {
	var err error
	var item Item
	switch df.Type {
	case uap.Fixed:
		item = NewFixed(df)
	case uap.Extended:
		item = NewExtended(df)
	case uap.Repetitive:
		item = NewRepetitive(df)
	case uap.Explicit:
		item = NewExplicit(df)
	case uap.Compound:
		item = NewCompound(df)
	case uap.SP, uap.RE:
		item = NewSpecialPurpose(df)
	case uap.RFS:
		item = NewRandomFieldSequencing(df)
	default:
		err = ErrDataFieldUnknown
		return nil, err
	}
	return item, err
}


type Base struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        uap.TypeField
}

func (b *Base) NewBase(field uap.DataField) {
	b.FRN = field.FRN
	b.DataItem = field.DataItem
	b.Description = field.Description
	b.Type = field.Type
}

// Frn returns FRN number of dataField from UAP
func (b Base) Frn() uint8 {
	return b.FRN
}
