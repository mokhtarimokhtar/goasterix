package item

import "bytes"

type Spare struct {
	Base
}

func newSpare(field DataItem) DataItem {
	f := &Spare{}
	f.Base.NewBase(field)
	return f
}
func (s Spare) GetSize() SizeField {
	return SizeField{}
}
func (s Spare) Payload() []byte {
	return nil
}
func (s Spare) GetSubItem() []SubItem {
	return nil
}
func (s Spare) String() string {
	return ""
}
func (s Spare) Reader(rb *bytes.Reader) error {
	return nil
}
func (s Spare) GetCompound() []DataItem {
	return nil // not used, it's for implement DataItemName interface
}
