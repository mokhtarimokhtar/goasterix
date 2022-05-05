package item

import (
	"bytes"
	"errors"
	//"github.com/mokhtarimokhtar/goasterix/_uap"
)

var (
	// ErrDataFieldUnknown reports which ErrDatafield Unknown.
	ErrDataFieldUnknown = errors.New("type of datafield not found")
	// ErrSubDataFieldUnknown reports which ErrDatafield Unknown.
	ErrSubDataFieldUnknown = errors.New("type of sub datafield not found")
	// ErrSubDataFieldFormat reports which ErrDatafield Format.
	ErrSubDataFieldFormat = errors.New("sub datafield incorrect")
	// ErrSubItemNotFound reports which sub-item not found in list of SubItems.
	ErrSubItemNotFound = errors.New("sub-item not found")
)

type TypeField uint8

const (
	FixedField TypeField = iota + 1
	ExtendedField
	CompoundField
	RepetitiveField
	ExplicitField
	SPField
	REField
	RFSField
	SpareField
	BitField
	FromToField
)

// String displays the name of iota value in string
func (t TypeField) String() string {
	return [...]string{"", "Fixed", "Extended", "Compound", "Repetitive", "Explicit", "SP", "RE", "RFS", "Spare",
		"Bit", "FromToBit"}[t]
}

type IBase interface {
	GetFrn() FieldReferenceNumber
	GetType() TypeField
	GetDataItemName() string
	GetDescription() string
}
type DataItem interface {
	IBase
	Clone() DataItem
	GetType() TypeField
	Reader(*bytes.Reader) error
	GetSubItems() []SubItem
	String() string
	//Payload() []byte
}

/*
type PayloadDataItem struct {
	Data []byte
}

func (p *PayloadDataItem) Payload(dataItem DataItem) {
	switch dataItem.GetType() {
	case FixedField:
		d := dataItem.(*Fixed)
		p.Data = make([]byte, 0, d.Size)
		p.Data = d.Data
	case ExtendedField:
		d := dataItem.(*Extended)
		p.Data = make([]byte, 0, len(d.Primary)+len(d.Secondary))
		p.Data = append(p.Data, d.Primary...)
		p.Data = append(p.Data, d.Secondary...)
	case RepetitiveField:
		d := dataItem.(*Repetitive)
		p.Data = make([]byte, 0, d.Rep*d.SubItemSize)
		p.Data = append(p.Data, d.Rep)
		p.Data = append(p.Data, d.Data...)
	}
}
*/

/*type DataItem interface {
	IBase
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	GetSize() SizeField
	GetCompound() []DataItem
	GetSubItem() []SubItem
}*/
