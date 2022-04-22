package goasterix

import (
	"bytes"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Item interface {
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	Frn() uint8
}

// Readers extracts data from the corresponding Item type.
func Readers(i Item, rb *bytes.Reader) error {
	err := i.Reader(rb)
	return err
}

// GetItem returns the corresponding Item type: Fixed, Extended, etc.
// GetItem is a factory function
func GetItem(df uap.IDataField) (Item, error) {
	var err error
	var item Item
	switch df.GetType() {
	case uap.Fixed:
		item = newFixed(df)
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

// GetItemCompound returns the corresponding Item type: Fixed, Extended, etc.
// GetItemCompound is a factory function for compound item type
func GetItemCompound(df uap.IDataField) (Item, error) {
	var err error
	var item Item
	switch df.GetType() {
	case uap.Fixed:
		item = newFixed(df)
	case uap.Extended:
		item = NewExtended(df)
	case uap.Repetitive:
		item = NewRepetitive(df)
	case uap.Explicit:
		item = NewExplicit(df)
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

func (b *Base) NewBase(field uap.IDataField) {
	b.FRN = field.GetFrn()
	b.DataItem = field.GetDataItem()
	b.Description = field.GetDescription()
	b.Type = field.GetType()
}

// Frn returns FRN number of dataField from UAP
func (b Base) Frn() uint8 {
	return b.FRN
}
