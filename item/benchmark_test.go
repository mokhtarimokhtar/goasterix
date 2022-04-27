package item

import (
	"bytes"
	"testing"
)

func BenchmarkFspecReader(b *testing.B) {
	input := []byte{0xFF, 0x01, 0xF2}
	rb := bytes.NewReader(input)

	for n := 0; n < b.N; n++ {
		_, _ = FspecReader(rb)
	}
}

func BenchmarkFspecIndex(b *testing.B) {
	input := []byte{0xef, 0x98}

	for n := 0; n < b.N; n++ {
		_ = FspecIndex(input)
	}
}

func BenchmarkSubItemBitReader(b *testing.B) {
	input := []byte{0x10, 0x00, 0x00, 0x00}
	for n := 0; n < b.N; n++ {
		sub := new(SubItemBit)
		sub.Name = "item1"
		sub.Type = BitField
		sub.Pos = 29
		_ = sub.Reader(input)
	}
}

func BenchmarkSubItemFromToReader(b *testing.B) {
	input := []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	for n := 0; n < b.N; n++ {
		sub := new(SubItemFromTo)
		sub.Name = "item1"
		sub.Type = FromToField
		sub.From = 27
		sub.To = 8
		_ = sub.Reader(input)
	}

}

func BenchmarkOneBitReader(b *testing.B) {
	var input byte = 0xd5
	var pos = uint8(6)

	for n := 0; n < b.N; n++ {
		_ = OneBitReader(input, pos)
	}
}
