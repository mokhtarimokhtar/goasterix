package goasterix

import (
	"bytes"
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

// Compound Data Fields, being of a variable length, shall comprise a primary subfield, followed by data subfields.
// The primary subfield determines the presence or absence of the subsequent data subfields. It comprises a first part
// of one octet extendable using the Field Extension (FX) mechanism.
// The definition, structure and format of the data subfields are part of the description of the relevant Compound Data
// Item. Data subfields shall be either fixed length, extended length, explicit length or repetitive, but not compound.
type Compound struct {
	Base
	Fields    []uap.DataField
	Primary   []byte
	Secondary []Item
}

func NewCompound(field uap.IDataField) Item {
	f := &Compound{}
	f.Base.NewBase(field)
	f.Fields = field.GetCompound()
	return f
}

func (c *Compound) Reader(rb *bytes.Reader) error {
	var err error

	c.Primary, err = FspecReader(rb)
	if err != nil {
		return err
	}
	frnIndex := FspecIndex(c.Primary)

	for _, frn := range frnIndex {
		uapItem := c.Fields[frn-1]
		var item Item
		switch uapItem.Type {
		case uap.Fixed:
			tmp := NewFixed(uapItem)
			err = tmp.Reader(rb)
			if err != nil {
				return err
			}
			item = tmp

		case uap.Extended:
			tmp := NewExtended(uapItem)
			err = tmp.Reader(rb)
			if err != nil {
				return err
			}
			item = tmp

		case uap.Explicit:
			tmp := NewExplicit(uapItem)
			err = tmp.Reader(rb)
			if err != nil {
				return err
			}
			item = tmp

		case uap.Repetitive:
			tmp := NewRepetitive(uapItem)
			err = tmp.Reader(rb)
			if err != nil {
				return err
			}
			item = tmp
		default:
			err = ErrDataFieldUnknown
			return err
		}
		c.Secondary = append(c.Secondary, item)
	}

	return err
}

// Payload returns this dataField as bytes.
func (c Compound) Payload() []byte {
	var p []byte
	p = append(p, c.Primary...)
	for _, item := range c.Secondary {
		tmp := item.Payload()
		p = append(p, tmp...)
	}
	return p
}

// String implements fmt.Stringer in hexadecimal
func (c Compound) String() string {
	var buf bytes.Buffer
	buf.Reset()
	buf.WriteString(c.Base.DataItem)
	buf.WriteByte(':')
	buf.WriteByte('[')
	buf.WriteString("primary:")
	buf.WriteString(hex.EncodeToString(c.Primary))
	buf.WriteByte(']')

	for _, item := range c.Secondary {
		buf.WriteByte('[')
		buf.WriteString(item.String())
		buf.WriteByte(']')
	}

	return buf.String()
}

