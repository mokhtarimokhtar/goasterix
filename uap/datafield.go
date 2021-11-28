package uap

type TypeField uint8

const (
	Fixed TypeField = iota + 1
	Extended
	Compound
	Repetitive
	Explicit
	SP
	RE
	RFS
	Spare
)

// StandardUAP is User Application Profile
// Cat is ASTERIX Category number (integer)
// Version is ASTERIX version for a category
type StandardUAP struct {
	Name     string
	Category uint8
	Version  float64
	Items    []DataField
}

// DataField describes FRN(Field Reference Number)
type DataField struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        TypeField
	Fixed       FixedField
	Extended    ExtendedField
	Repetitive  RepetitiveField
	Explicit    ExplicitField
	Compound    []DataField
	Conditional bool
}
type FixedField struct {
	Size uint8
}
type ExtendedField struct {
	PrimarySize   uint8
	SecondarySize uint8
}
type RepetitiveField struct {
	SubItemSize uint8
}
type ExplicitField struct {
}

/*
type TypeField struct {
	NameType TypeField
	Size     uint8
	PrimarySize   uint8    // used for extended
	SecondarySize uint8    // used for extended
	Primary       *Primary // used for compound
}

// Bit is a bit number position
type Bit uint8

// Size is the number of bytes (size) of the corresponding field
type Size uint8

// Primary subitem, followed by data subitems.
// The Compound Data Item primary subitem determines the presence or absence of the subsequent
// data subitems and is made of a first part of one octet extendible using the Field Extension (FX) mechanism
type Primary []MetaField

// MetaField contains a hashmap [number of bit key] <=> [type of field and size]
type MetaField map[Bit]Subfield

// Subfield describes the type of field and the size if necessary
type Subfield struct {
	NameType      TypeField
	Size          uint8
	PrimarySize   uint8
	SecondarySize uint8
	Item          string
	Description   string
}
*/
