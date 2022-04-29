package item

import "bytes"

// FspecReader returns a slice of FSPEC data record asterix.
func FspecReader(rb *bytes.Reader) ([]byte, error) {
	var fspec []byte
	var err error
	var tmp byte
	for {
		tmp, err = rb.ReadByte()
		if err != nil {
			return nil, err
		}
		fspec = append(fspec, tmp)
		if tmp&0x01 == 0 {
			break
		}
	}
	return fspec, err
}

// FspecIndex returns an array of uint8 corresponding to number FieldReferenceNumber(Field Reference Number of DataItems).
// In other words, it transposes a fspec pos to an array FRNs.
// e.g. fspec = 1010 1010 => frnIndex = []uint8{1, 3, 5, 7}
func FspecIndex(fspec []byte) []uint8 {
	l := uint8(len(fspec))
	var frnIndex = make([]uint8, 0, l*7)
	var tmp byte
	for j := uint8(0); j < l; j++ {
		for i := uint8(0); i < 7; i++ {
			tmp = fspec[j] << i
			if tmp&0x80 != 0 {
				frnIndex = append(frnIndex, 7*j+i+1)
			}
		}
	}
	return frnIndex
}
