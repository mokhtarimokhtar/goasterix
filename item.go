package goasterix

import (
	"bytes"
	"encoding/binary"
	//"github.com/mokhtarimokhtar/goasterix/_uap"
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
	GetCompound() []Item
	Item
}*/

type Item interface {
	IBase
	Payload() []byte
	String() string
	Reader(*bytes.Reader) error
	GetSize() SizeField
	GetCompound() []Item
}

// Readers extracts data from the corresponding Item type.
func Readers(i Item, rb *bytes.Reader) error {
	err := i.Reader(rb)
	return err
}

// GetItem returns the corresponding Item type: Fixed, Extended, etc.
// GetItem is a factory function
//func GetItem(df _uap.IDataField) (Item, error) {
func GetItem(i Item) (Item, error) {
	var err error
	var item Item
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

// GetItemCompound returns the corresponding Item type: Fixed, Extended, etc.
// GetItemCompound is a factory function for compound item type
//func GetItemCompound(df uap.IDataField) (Item, error) {
func GetItemCompound(i Item) (Item, error) {
	var err error
	var item Item
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

type IBase interface {
	GetFrn() uint8
	GetType() TypeField
	GetDataItem() string
	GetDescription() string
}

type Base struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        TypeField
}

func (b *Base) NewBase(field Item) {
	b.FRN = field.GetFrn()
	b.DataItem = field.GetDataItem()
	b.Description = field.GetDescription()
	b.Type = field.GetType()
}

// GetFrn returns FRN number of dataField from UAP
func (b Base) GetFrn() uint8 {
	return b.FRN
}
func (b Base) GetType() TypeField {
	return b.Type
}
func (b Base) GetDataItem() string {
	return b.DataItem
}
func (b Base) GetDescription() string {
	return b.Description
}

func FromToBitReader8(data byte, from uint8, to uint8) byte {
	var v byte
	mask := uint8(0xff)
	diff := from - to
	shift := to - uint8(1)
	keep := mask >> (uint8(8) - uint8(diff+uint8(1)))
	v = data >> shift & keep
	return v
}

//func FromToBitReader16(data uint16, from uint8, to uint8) uint16 {
func FromToBitReader16(data []byte, from uint8, to uint8) []byte {
	tmp := binary.BigEndian.Uint16(data)
	var v uint16
	diff := from - to
	mask := uint16(0xffff)
	shift := to - uint8(1)
	keep := mask >> (uint8(16) - uint8(diff+uint8(1)))
	//v = data >> shift & keep
	v = tmp >> shift & keep
	binary.BigEndian.PutUint16(data, v)
	return data
}

func FromToBitReader32(data uint32, from uint8, to uint8) uint32 {
	var v uint32
	diff := from - to
	mask := uint32(0xffffffff)
	shift := to - uint8(1)
	keep := mask >> (uint8(32) - uint8(diff+uint8(1)))
	v = data >> shift & keep
	return v
}

func Uint64ToByte(val uint64) []byte {
	r := make([]byte, 8)
	for i := uint64(0); i < 8; i++ {
		r[i] = byte((val >> (i * 8)) & 0xff)
	}
	return r
}

func ByteToUint64(val []byte) uint64 {
	r := uint64(0)
	for i := uint64(0); i < 8; i++ {
		r |= uint64(val[i]) << (8 * i)
	}
	return r
}

func Uint32ToByte(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}

func ByteToUint32(val []byte) uint32 {
	r := uint32(0)
	for i := uint32(0); i < 4; i++ {
		r |= uint32(val[i]) << (8 * i)
	}
	return r
}
