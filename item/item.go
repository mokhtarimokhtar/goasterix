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

/*type IBase interface {
	GetFrn() FieldReferenceNumber
	GetType() TypeField
	GetDataItemName() string
	GetDescription() string
}*/
type DataItem interface {
	//IBase
	Clone() DataItem
	GetType() TypeField
	Reader(*bytes.Reader) error
	String() string
	Payload() []byte
}

/*type DataItem interface {
	IBase
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	GetSize() SizeField
	GetCompound() []DataItem
	GetSubItem() []SubItem
}*/
