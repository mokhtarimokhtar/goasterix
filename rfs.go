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

// Payload returns this dataField as bytes.
func (rf RandomField) Payload() []byte {
	var p []byte
	p = append(p, rf.FRN)
	p = append(p, rf.Field.Payload()...)
	return p
}

// String implements fmt.Stringer in hexadecimal
func (rf RandomField) String() string {
	var buf bytes.Buffer
	tmp := []byte{rf.FRN}
	buf.Reset()
	buf.WriteString("FRN:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteString(" ")
	buf.WriteString(rf.Field.String())
	return buf.String()
}

// RandomFieldSequencing
// The RFS organised dataField is a collection of Data Fields which in
// contrast to the OFS organisation, can occur in any order.
// The RFS organised dataField shall be structured as follows:
// - the first octet provides the number, N, of Data Fields following;
// - N fields in any arbitrary order each consisting of a one-octet FRN immediately followed by the contents of the
// Data Item associated with the preceding FRN.
type RandomFieldSequencing struct {
	Base
	Fields   []uap.DataField
	N        uint8
	Sequence []RandomField
}

func NewRandomFieldSequencing(field uap.IDataField) Item {
	f := &RandomFieldSequencing{}
	f.Base.NewBase(field)
	//f.Fields = field.RFS
	f.Fields = field.GetRFS()
	return f
}

func (rfs *RandomFieldSequencing) Reader(rb *bytes.Reader) error {
	var err error
	// N is the total number of datafields
	err = binary.Read(rb, binary.BigEndian, &rfs.N)
	if err != nil {
		return err
	}

	for i := uint8(0); i < rfs.N; i++ {
		// retrieve random FRN
		var frn uint8
		err = binary.Read(rb, binary.BigEndian, &frn)
		if err != nil {
			return err
		}

		for _, uapItem := range rfs.Fields {
			if frn == uapItem.FRN {
				rf := new(RandomField)
				rf.FRN = frn
				// todo: add other datafield use case (work just for Fixed)
				switch uapItem.Type {
				case uap.Fixed:
					tmp := new(Fixed)
					err = tmp.Reader(rb)
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

// Payload returns this dataField as bytes.
func (rfs RandomFieldSequencing) Payload() []byte {
	var p []byte
	p = append(p, rfs.N)
	for _, field := range rfs.Sequence {
		p = append(p, field.Payload()...)
	}
	return p
}

// String implements fmt.Stringer in hexadecimal
func (rfs RandomFieldSequencing) String() string {
	var buf bytes.Buffer
	buf.Reset()
	tmp := []byte{rfs.N}
	buf.WriteString(rfs.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteByte('[')
	buf.WriteString("N:")
	buf.WriteString(hex.EncodeToString(tmp))
	buf.WriteByte(']')

	for _, item := range rfs.Sequence {
		buf.WriteByte('[')
		buf.WriteString(item.String())
		buf.WriteByte(']')
	}

	return buf.String()
}
