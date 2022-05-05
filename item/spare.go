package item

import "bytes"

// Spare is a spare data item in UAP, this type is interesting for debug.
type Spare struct {
	Base
}

// All function: not used, it's for implement DataItem interface

func (s *Spare) Clone() DataItem {
	return &Spare{
		Base: s.Base,
	}
}

func (s Spare) GetSubItems() []SubItem {
	return nil
}
func (s Spare) String() string {
	return ""
}
func (s Spare) Reader(rb *bytes.Reader) error {
	return nil
}

/*
func (s Spare) Payload() []byte {
	return nil
}*/
