package goasterix

import "math"

/*
 Package util complement implements the functions to calculate the two's complement value for specific bit size.
*/

// TwoComplement16 returns an int16 (signed).
// sizeBits is the number of bit complement.
func TwoComplement16(sizeBits uint8, data uint16) (v int16) {
	n := float64(sizeBits - 1)
	p := math.Pow(2, n) // 2^(N-1)
	mask := uint16(p)

	tmp1 := -int16(data & mask)
	tmp2 := int16(data & ^mask)

	v = tmp1 + tmp2
	return v
}

// TwoComplement32 returns an int32 (signed).
// sizeBits is the number of bit complement.
func TwoComplement32(sizeBits uint8, data uint32) (v int32) {
	n := float64(sizeBits - 1)
	p := math.Pow(2, n) // 2^(N-1)
	mask := uint32(p)

	tmp1 := -int32(data & mask)
	tmp2 := int32(data & ^mask)

	v = tmp1 + tmp2

	// checking example
	// mask := uint32(0x007F)
	/*signed := data[2] & 0x80 >> 7
	if signed == 1 {
		complement := ^tmpLatitude 	// one complement
		v := complement & mask	// apply mask 2^(N-1)
		v = v + 1
		latitude = -float64(v) * 0.000021458
	} else {
		latitude = float64(tmpLatitude) * 0.000021458
	}*/

	return v
}
