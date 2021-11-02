package model

// TableIA5 International Alphabet 5
// A - Z = 1 - 26
// 0 - 9 = 48 - 57
// space :  32
var TableIA5 = map[uint8]string{
		uint8(1):  "A",
		uint8(2):  "B",
		uint8(3):  "C",
		uint8(4):  "D",
		uint8(5):  "E",
		uint8(6):  "F",
		uint8(7):  "G",
		uint8(8):  "H",
		uint8(9):  "I",
		uint8(10): "J",
		uint8(11): "K",
		uint8(12): "L",
		uint8(13): "M",
		uint8(14): "N",
		uint8(15): "O",
		uint8(16): "P",
		uint8(17): "Q",
		uint8(18): "R",
		uint8(19): "S",
		uint8(20): "T",
		uint8(21): "U",
		uint8(22): "V",
		uint8(23): "W",
		uint8(24): "X",
		uint8(25): "Y",
		uint8(26): "Z",
		uint8(32): " ", // SP
		uint8(48): "0",
		uint8(49): "1",
		uint8(50): "2",
		uint8(51): "3",
		uint8(52): "4",
		uint8(53): "5",
		uint8(54): "6",
		uint8(55): "7",
		uint8(56): "8",
		uint8(57): "9",
	}


// obsolete
// SixBitToASCII converts a char of six bits to ASCII char.
// A - Z :   1 - 26
// 0 - 9 :  48 - 57
// _ :  32
/*func SixBitToASCII(b uint16) (s string, err error) {
	switch {
	case 48 <= b && b <= 57:
		return string(b), nil
	case b == 32:
		return string(b), nil
	//case b < 48:
	case 1 <= b && b <= 26:
		return string(b + 64), nil
	default:
		return "", ErrSixBitCharUnknown
	}
}*/