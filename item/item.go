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

/*type ICompound interface {
	GetCompound() []DataItemName
	DataItemName
}*/

type DataItem interface {
	IBase
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	GetSize() SizeField
	GetCompound() []DataItem
	GetSubItem() []SubItem
}

type IBase interface {
	GetFrn() uint8
	GetType() TypeField
	GetDataItemName() string
	GetDescription() string
}

// Readers extracts data from the corresponding DataItem type.
func Readers(i DataItem, rb *bytes.Reader) error {
	err := i.Reader(rb)
	return err
}

// GetItem returns the corresponding DataItem type: Fixed, Extended, etc.
// GetItem is a factory function
//func GetItem(df _uap.IDataField) (DataItem, error) {
func GetItem(i DataItem) (DataItem, error) {
	var err error
	var item DataItem
	switch i.GetType() {
	case FixedField:
		item = newFixed(i)
	case ExtendedField:
		item = NewExtended(i)
	case ExplicitField:
		item = NewExplicit(i)
	case RepetitiveField:
		item = NewRepetitive(i)
	case SPField:
		item = NewSpecialPurpose(i)
	case REField:
		item = NewReservedExpansion(i)
	case CompoundField:
		//var c ICompound
		item = NewCompound(i)
		/*
			case uap.RFS:
				item = NewRandomFieldSequencing(i)*/
	default:
		err = ErrDataFieldUnknown
		return nil, err
	}
	return item, err
}

// GetItemCompound returns the corresponding DataItem type: Fixed, Extended, etc.
// GetItemCompound is a factory function for compound item type
//func GetItemCompound(df uap.IDataField) (DataItem, error) {
func GetItemCompound(i DataItem) (DataItem, error) {
	var err error
	var item DataItem
	switch i.GetType() {
	case FixedField:
		item = newFixed(i)
	case ExtendedField:
		item = NewExtended(i)
	case RepetitiveField:
		item = NewRepetitive(i)
	case ExplicitField:
		item = NewExplicit(i)
	default:
		err = ErrDataFieldUnknown
		return nil, err
	}
	return item, err
}

type SizeField struct {
	ForFixed             uint8
	ForExtendedPrimary   uint8
	ForExtendedSecondary uint8
	ForRepetitive        uint8
}
