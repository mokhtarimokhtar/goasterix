package goasterix

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
	"math/bits"

	"github.com/mokhtarimokhtar/goasterix/uap"
)

var (
	// ErrDataFieldUnknown reports which ErrDatafield Unknown.
	ErrDataFieldUnknown = errors.New("type of datafield not found")
)

type Record struct {
	Cat   uint8
	Fspec []byte
	Items []Item
}

func NewRecord() *Record {
	return &Record{}
}

// Decode extracts a Record of asterix data block (only one record).
// An asterix data block can contain a or more records.
// It returns the number of bytes unread and fills the Record Struct(Fspec, Items array) in byte.
func (rec *Record) Decode(data []byte, stdUAP uap.StandardUAP) (unRead int, err error) {
	rec.Cat = stdUAP.Category

	rb := bytes.NewReader(data)
	rec.Fspec, err = FspecReader(rb)
	unRead = rb.Len()
	if err != nil {
		return unRead, err
	}

	frnIndex := FspecIndex(rec.Fspec)
	offset := uint8(0) // offset shifts the index for a conditional UAP

	for _, frn := range frnIndex {
		uapItem := stdUAP.Items[frn-1-offset] // here the index corresponds to the FRN

		item := NewItem(uapItem)
		switch uapItem.Type {
		case uap.Fixed:
			tmp, err := FixedDataFieldReader(rb, uapItem.Fixed.Size)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.Fixed = &tmp

		case uap.Extended:
			tmp, err := ExtendedDataFieldReader(rb, uapItem.Extended.PrimarySize, uapItem.Extended.SecondarySize)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.Extended = &tmp

		case uap.Explicit:
			tmp, err := ExplicitDataFieldReader(rb)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.Explicit = &tmp

		case uap.Repetitive:
			tmp, err := RepetitiveDataFieldReader(rb, uapItem.Repetitive.SubItemSize)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.Repetitive = &tmp

		case uap.Compound:
			tmp, err := CompoundDataFieldReader(rb, uapItem.Compound)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.Compound = &tmp

		case uap.SP, uap.RE:
			tmp, err := SPAndREDataFieldReader(rb)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.SP = &tmp

		case uap.RFS:
			tmp, err := RFSDataFieldReader(rb, stdUAP.Items)
			if err != nil {
				unRead = rb.Len()
				return unRead, err
			}
			item.RFS = &tmp

		default:
			err = ErrDataFieldUnknown
			return unRead, err
		}
		unRead = rb.Len()
		rec.Items = append(rec.Items, *item)

		if uapItem.Conditional {
			switch item.Meta.Type {
			case uap.Fixed:
				stdUAP.Items = selectUAPConditional(stdUAP.Category, item.Fixed.Data)
			case uap.Extended:
				stdUAP.Items = selectUAPConditional(stdUAP.Category, item.Extended.Primary)
			}
			offset = frn
		}
	}
	return unRead, nil
}

// String returns a string(hex) representation of one asterix record (only existing items).
func (rec *Record) String() []string {
	var items []string
	tmp := "FSPEC: " + hex.EncodeToString(rec.Fspec)
	items = append(items, tmp)

	for _, item := range rec.Items {
		tmp := item.String()
		items = append(items, tmp)
	}
	return items
}

// Payload returns a slice of byte for one asterix record.
func (rec *Record) Payload() (b []byte) {
	b = append(b, rec.Fspec...)
	for _, item := range rec.Items {
		b = append(b, item.Payload()...)
	}
	return b
}

func selectUAPConditional(category uint8, field []byte) []uap.DataField {
	var selectedUAP []uap.DataField
	switch category {
	case 1:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = uap.Cat001TrackV12
		} else {
			selectedUAP = uap.Cat001PlotV12
		}
	case 26:
		tmp := field[0] & 0x80 >> 7
		if tmp == 1 {
			selectedUAP = uap.Cat4TestTrack
		} else {
			selectedUAP = uap.Cat4TestPlot
		}
	}
	return selectedUAP
}

// FspecReader returns a slice of FSPEC data record asterix.
func FspecReader(reader io.Reader) ([]byte, error) {
	var fspec []byte
	var err error
	for {
		var tmp uint8
		err = binary.Read(reader, binary.BigEndian, &tmp)
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

// FspecIndex returns an array of uint8 corresponding to number FRN(Field Reference Number of Items).
// In other words, it converts a fspec bits to an array FRNs.
// e.g. fspec = 1010 1010 => frnIndex = []uint8{1, 3, 5, 7}
func FspecIndex(fspec []byte) []uint8 {
	var frnIndex []uint8
	for j, val := range fspec {
		for i := 0; i < 7; i++ {
			frn := 7*j + i + 1
			tmp := bits.RotateLeft8(val, i)
			if tmp&0x80 != 0 {
				frnIndex = append(frnIndex, uint8(frn))
			}
		}
	}
	return frnIndex
}

// FixedDataFieldReader extracts a number(nb) of bytes(size) and returns a slice of bytes(data of item).
// Fixed length Data Fields shall comprise a fixed number of octets.
func FixedDataFieldReader(rb *bytes.Reader, size uint8) (Fixed, error) {
	var err error
	item := Fixed{}

	tmp := make([]byte, size)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return item, err
	}
	item.Data = tmp
	return item, err
}

// ExtendedDataFieldReader extracts data item type Extended (FX: last bit = 1).
// primarySize parameter defines the Primary Subitem of extended field.
// secondarySize parameter defines the Secondary Subitem of extended field.
// Extended length Data Fields, being of a variable length, shall contain a primary part of predetermined length,
// immediately followed by a number of secondary parts, each of predetermined length.
// The presence of the next following secondary part shall be indicated by the setting to one of the
// Least Significant Bit (LSB) of the last octet of the preceding part (either the primary part or a secondary part).
// This bit which is reserved for that purpose is called the Field Extension Indicator (FX).
func ExtendedDataFieldReader(rb *bytes.Reader, primarySize uint8, secondarySize uint8) (Extended, error) {
	var err error
	item := Extended{}

	tmp := make([]byte, primarySize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return item, err
	}
	item.Primary = tmp

	if tmp[primarySize-1]&0x01 != 0 {
		for {
			tmp := make([]byte, secondarySize)
			err = binary.Read(rb, binary.BigEndian, &tmp)
			if err != nil {
				return item, err
			}
			item.Secondary = append(item.Secondary, tmp...)
			if tmp[secondarySize-1]&0x01 == 0 {
				break
			}
		}
	}
	return item, err
}

// ExplicitDataFieldReader extracts a number of bytes define by the first byte.
// Explicit length Data Fields shall start with a one-octet length indicator giving
// the total field length in octets including the length indicator itself.
func ExplicitDataFieldReader(rb *bytes.Reader) (Explicit, error) {
	var err error
	item := Explicit{}

	err = binary.Read(rb, binary.BigEndian, &item.Len)
	if err != nil {
		return item, err
	}

	tmp := make([]byte, item.Len-1)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return item, err
	}
	item.Data = tmp
	return item, err
}

// RepetitiveDataFieldReader extracts data item type Repetitive(1+rep*N byte).
// The first byte is REP(factor), nb is the size of bytes to repetition.
// Repetitive Data Fields, being of a variable length, shall comprise a one-octet Field Repetition Indicator (REP)
// signalling the presence of N consecutive sub-fields each of the same pre-determined length.
func RepetitiveDataFieldReader(rb *bytes.Reader, SubItemSize uint8) (Repetitive, error) {
	var err error
	item := Repetitive{}

	err = binary.Read(rb, binary.BigEndian, &item.Rep)
	if err != nil {
		return item, err
	}

	tmp := make([]byte, item.Rep*SubItemSize)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return item, err
	}
	item.Data = tmp
	return item, err
}

// CompoundDataFieldReader
// Compound Data Fields, being of a variable length, shall comprise a primary subfield, followed by data subfields.
// The primary subfield determines the presence or absence of the subsequent data subfields. It comprises a first part
// of one octet extendable using the Field Extension (FX) mechanism.
// The definition, structure and format of the data subfields are part of the description of the relevant Compound Data
// Item. Data subfields shall be either fixed length, extended length, explicit length or repetitive, but not compound.
func CompoundDataFieldReader(rb *bytes.Reader, cp []uap.DataField) (Compound, error) {
	var err error
	items := Compound{}

	items.Primary, err = FspecReader(rb)
	if err != nil {
		return items, err
	}
	frnIndex := FspecIndex(items.Primary)

	for _, frn := range frnIndex {
		uapItem := cp[frn-1]
		item := NewItem(uapItem)
		switch uapItem.Type {
		case uap.Fixed:
			tmp, err := FixedDataFieldReader(rb, uapItem.Fixed.Size)
			if err != nil {
				return items, err
			}
			item.Fixed = &tmp
			items.Secondary = append(items.Secondary, *item)

		case uap.Extended:
			tmp, err := ExtendedDataFieldReader(rb, uapItem.Extended.PrimarySize, uapItem.Extended.SecondarySize)
			if err != nil {
				return items, err
			}
			item.Extended = &tmp
			items.Secondary = append(items.Secondary, *item)

		case uap.Explicit:
			tmp, err := ExplicitDataFieldReader(rb)
			if err != nil {
				return items, err
			}
			item.Explicit = &tmp
			items.Secondary = append(items.Secondary, *item)

		case uap.Repetitive:
			tmp, err := RepetitiveDataFieldReader(rb, uapItem.Repetitive.SubItemSize)
			if err != nil {
				return items, err
			}
			item.Repetitive = &tmp
			items.Secondary = append(items.Secondary, *item)

		default:
			err = ErrDataFieldUnknown
			return items, err
		}
	}
	return items, err
}

// RFSDataFieldReader
// The RFS organised field is a collection of Data Fields which in
// contrast to the OFS organisation, can occur in any order.
// The RFS organised field shall be structured as follows:
// - the first octet provides the number, N, of Data Fields following;
// - N fields in any arbitrary order each consisting of a one-octet FRN immediately followed by the contents of the
// Data Item associated with the preceding FRN.
func RFSDataFieldReader(rb *bytes.Reader, items []uap.DataField) (RandomFieldSequencing, error) {
	var err error
	rfs := RandomFieldSequencing{}
	// N is the total number of datafields
	err = binary.Read(rb, binary.BigEndian, &rfs.N)
	if err != nil {
		return rfs, err
	}
	for i := uint8(0); i < rfs.N; i++ {
		// retrieve random FRN
		var frn uint8
		err := binary.Read(rb, binary.BigEndian, &frn)
		if err != nil {
			return rfs, err
		}

		for _, field := range items {
			if frn == field.FRN {
				rf := new(RandomField)
				rf.FRN = frn

				// todo: work just for Fixed datafield use case
				tmp := make([]byte, field.Fixed.Size)

				err := binary.Read(rb, binary.BigEndian, &tmp)
				if err != nil {
					return rfs, err
				}
				item := NewItem(field)
				fixed := Fixed{}
				fixed.Data = tmp
				item.Fixed = &fixed
				rf.Field = *item
				rfs.Sequence = append(rfs.Sequence, *rf)
			}
		}
	}

	return rfs, err
}

// SPAndREDataFieldReader extracts returns a slice
// ref. EUROCONTROL-SPEC-0149 2.4
// 4.3.5 Non-Standard Data Fields:
// Reserved Expansion Data
// Field Special Purpose field
func SPAndREDataFieldReader(rb *bytes.Reader) (SpecialPurpose, error) {
	var err error
	sp := SpecialPurpose{}

	err = binary.Read(rb, binary.BigEndian, &sp.Len)
	if err != nil {
		return sp, err
	}

	tmp := make([]byte, sp.Len-1)
	err = binary.Read(rb, binary.BigEndian, &tmp)
	if err != nil {
		return sp, err
	}
	sp.Data = tmp

	return sp, err
}
