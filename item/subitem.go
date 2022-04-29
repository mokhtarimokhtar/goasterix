package item

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
)

// SubItemBits has two types: One bit or From To bits
// Bit field is for one bit position
// From and To is for the range of bits position
type SubItemBits struct {
	Name string
	Type TypeField
	Bit  uint8
	From uint8
	To   uint8
	Data []byte
}

func (s *SubItemBits) Clone() *SubItemBits {
	return &SubItemBits{
		Name: s.Name,
		Type: s.Type,
		Bit:  s.Bit,
		From: s.From,
		To:   s.To,
	}
}

func (s *SubItemBits) Reader(data []byte) error {
	var err error
	switch s.Type {
	case BitField:
		totalBits := uint8(len(data)) * 8
		indexData := (totalBits - s.Bit) / 8
		relativePos := s.Bit - (uint8((s.Bit-1)/8) * 8)

		s.Data = make([]byte, 1)
		s.Data[0] = OneBitReader(data[indexData], relativePos)

	case FromToField:
		s.Data, err = FromToBitReader(data, s.From, s.To)

	default:
		err = ErrSubDataFieldUnknown
		return err
	}
	return err
}

// String implements fmt.Stringer in hexadecimal
func (s SubItemBits) String() string {
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
