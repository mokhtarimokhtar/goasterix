package uap

type NameTypeField uint8

const (
	Fixed NameTypeField = iota + 1
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
	Conditional bool
	Payload     []byte
}

type TypeField struct {
	NameType      NameTypeField
	Size          uint8
	PrimarySize   uint8
	SecondarySize uint8
	Primary       *Primary
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
	NameType      NameTypeField
	Size          uint8
	PrimarySize   uint8
	SecondarySize uint8
	Item          string
	Description   string
}

type UAP struct {
	Name     string
	Category uint8
	Version  float64
}
type Category uint8
type Version float32

// DefaultProfiles contains the defaults User Application Profiles version.
var DefaultProfiles = map[uint8]StandardUAP{
	1:   Cat001V12,
	2:   Cat002V10,
	21:  Cat021v10,
	30:  Cat030StrV51,
	32:  Cat032StrV70,
	34:  Cat034V127,
	48:  Cat048V127,
	255: Cat255StrV51,
	62:  Cat062V119,
	// Category for testing not exist
	26: Cat4Test,
}
