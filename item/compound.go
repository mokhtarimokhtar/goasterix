package item

import (
	"bytes"
	"encoding/hex"
)

// Compound Data Fields, being of a variable length, shall comprise a primary subfield, followed by data subfields.
// The primary subfield determines the presence or absence of the subsequent data subfields. It comprises a first part
// of one octet extendable using the Field Extension (FX) mechanism.
// The definition, structure and format of the data subfields are part of the description of the relevant Compound Data
// DataItem. Data subfields shall be either fixed length, extended length, explicit length or repetitive, but not compound.
type Compound struct {
	Base
	Primary   []byte
	Secondary []DataItem
}

func (c *Compound) Clone() DataItem {
	return &Compound{
		Base:      c.Base,
		Primary:   c.Primary,
		Secondary: c.Secondary,
	}
}

func (c *Compound) Reader(rb *bytes.Reader) error {
	var err error

	c.Primary, err = FspecReader(rb)
	if err != nil {
		c.Secondary = nil
		return err
	}
	frnIndex := FspecIndex(c.Primary)
	tmp := c.Secondary // save temporary meta data DataItemName
	c.Secondary = make([]DataItem, 0, len(frnIndex))

	for _, frn := range frnIndex {
		uapItem := tmp[frn-1]
		dataItem := uapItem.Clone()

		typeItem := dataItem.GetType()
		if (typeItem != FixedField) && (typeItem != ExtendedField) && (typeItem != RepetitiveField) && (typeItem != ExplicitField) {
			return ErrDataFieldUnknown
		}

		err = dataItem.Reader(rb)
		if err != nil {
			return err
		}

		c.Secondary = append(c.Secondary, dataItem)
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
	buf.WriteString(c.Base.DataItemName)
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
