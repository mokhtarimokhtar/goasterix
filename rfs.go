package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type RandomField struct {
	FRN   uint8
	Field Item
}

func (rf RandomField) Payload() []byte {
	var p []byte
	p = append(p, rf.FRN)
	p = append(p, rf.Field.Payload()...)
	return p
}

func (rf RandomField) String() string {
	tmp := []byte{rf.FRN}
	return hex.EncodeToString(tmp) + hex.EncodeToString(rf.Field.Payload())
}

// RandomFieldSequencing
// The RFS organised field is a collection of Data Fields which in
// contrast to the OFS organisation, can occur in any order.
// The RFS organised field shall be structured as follows:
// - the first octet provides the number, N, of Data Fields following;
// - N fields in any arbitrary order each consisting of a one-octet FRN immediately followed by the contents of the
// Data Item associated with the preceding FRN.
type RandomFieldSequencing struct {
	MetaItem
	N        uint8
	Sequence []RandomField
}

func (rfs *RandomFieldSequencing) Reader(rb *bytes.Reader, field uap.DataField) error {
	var err error
	rfs.MetaItem.NewMetaItem(field)
	// N is the total number of datafields
	err = binary.Read(rb, binary.BigEndian, &rfs.N)
	if err != nil {
		return err
	}

	for i := uint8(0); i < rfs.N; i++ {
		// retrieve random FRN
		var frn uint8
		err := binary.Read(rb, binary.BigEndian, &frn)
		if err != nil {
			return err
		}

		for _, uapItem := range field.RFS {
			if frn == uapItem.FRN {
				rf := new(RandomField)
				rf.FRN = frn
				// todo: add other datafield use case (work just for Fixed)
				switch uapItem.Type {
					case uap.Fixed:
						tmp := new(Fixed)
						err = tmp.Reader(rb, uapItem)
						if err != nil {
							return err
						}
						rf.Field = tmp
						rfs.Sequence = append(rfs.Sequence, *rf)
				}
			}
		}
	}

	return err
}

func (rfs RandomFieldSequencing) Payload() []byte {
	var p []byte
	p = append(p, rfs.N)
	for _, field := range rfs.Sequence {
		p = append(p, field.Payload()...)
	}
	return p
}

func (rfs RandomFieldSequencing) String() string {
	var str string
	tmp := []byte{rfs.N}
	str = hex.EncodeToString(tmp)
	for _, field := range rfs.Sequence {
		str = str + field.String()
	}
	return str
}

func (rfs RandomFieldSequencing) Frn() uint8 {
	return rfs.MetaItem.FRN
}
