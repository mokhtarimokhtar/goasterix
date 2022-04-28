package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

type SubItem interface {
	Reader(data []byte) error
	GetName() string
	GetType() TypeField
	GetPosition() BitPosition
	String() string
}

// GetSubItem returns the corresponding DataItem type: BitField, FromToField.
// GetSubItem is a factory function
func GetSubItem(s SubItem) (SubItem, error) {
	var err error
	var item SubItem
	switch s.GetType() {
	case BitField:
		item = newSubItemBit(s)
	case FromToField:
		item = newSubItemFromTo(s)
	default:
		err = ErrSubDataFieldUnknown
		return nil, err
	}
	return item, err
}

type BitPosition struct {
	Bit  uint8
	From uint8
	To   uint8
}

type SubItemBit struct {
	Name string
	Type TypeField
	Pos  BitPosition
	Data []byte
}

func newSubItemBit(field SubItem) SubItem {
	f := &SubItemBit{}
	f.Name = field.GetName()
	f.Type = field.GetType()
	f.Pos = field.GetPosition()
	return f
}

func (s *SubItemBit) Reader(data []byte) error {
	var err error
	totalBits := uint8(len(data)) * 8
	indexData := (totalBits - s.Pos.Bit) / 8
	relativePos := s.Pos.Bit - (uint8((s.Pos.Bit-1)/8) * 8)

	s.Data = make([]byte, 1)
	s.Data[0] = OneBitReader(data[indexData], relativePos)

	return err
}
func (s SubItemBit) GetName() string {
	return s.Name
}
func (s SubItemBit) GetType() TypeField {
	return s.Type
}
func (s SubItemBit) GetPosition() BitPosition {
	return s.Pos
}

// String implements fmt.Stringer in hexadecimal
func (s SubItemBit) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(s.Name)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(s.Data))
	return buf.String()
}

type SubItemFromTo struct {
	Name string
	Type TypeField
	Pos  BitPosition
	Data []byte
}

func newSubItemFromTo(field SubItem) SubItem {
	f := &SubItemFromTo{}
	f.Name = field.GetName()
	f.Type = field.GetType()
	f.Pos = field.GetPosition()
	return f
}

func (s *SubItemFromTo) Reader(data []byte) error {
	var err error
	s.Data, err = FromToBitReader(data, s.Pos.From, s.Pos.To)
	return err
}
func (s SubItemFromTo) GetName() string {
	return s.Name
}
func (s SubItemFromTo) GetType() TypeField {
	return s.Type
}
func (s SubItemFromTo) GetPosition() BitPosition {
	return s.Pos
}

// String implements fmt.Stringer in hexadecimal
func (s SubItemFromTo) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(s.Name)
	buf.WriteByte(':')
	buf.WriteString(hex.EncodeToString(s.Data))
	return buf.String()
}

// OneBitReader returns a byte equal to the value of bit position
func OneBitReader(data byte, b uint8) byte {
	pos := b - 1
	return data >> pos & 0x01
}

// FromToBitReader returns a slice of byte corresponding to range the From and To.
func FromToBitReader(data []byte, from uint8, to uint8) ([]byte, error) {
	var err error

	if from <= to {
		return nil, ErrSubDataFieldFormat
	}

	fromToRangeBits := from - to

	dataLen := uint8(len(data))
	tmp := make([]byte, 8-dataLen, 8)
	tmp = append(tmp, data...)

	value := binary.BigEndian.Uint64(tmp)

	maskFrom := uint64(1<<uint8(from)) - 1
	maskTo := uint64(1<<uint8(to-1)) - 1
	mask := maskFrom & ^maskTo

	res := value & mask >> (to - 1)

	nbBytes := fromToRangeBits/8 + 1
	data = Uint64ToByteLess(res, nbBytes)

	return data, err
}

// Uint64ToByteLess converts uint64 to slice byte in Big Endian,
// it fills the slice juste enough.
// size in byte
func Uint64ToByteLess(val uint64, size uint8) []byte {
	r := make([]byte, size)
	for i := size; i > 0; i-- {
		r[size-i] = byte((val >> ((i - 1) * 8)) & 0xff)
	}
	return r
}
