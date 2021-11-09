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
	Name NameTypeField
	Size uint8
	Meta MetaField
}
type Bit uint8 // It is the bit number

type Size uint8 // It is the number of bytes (size) of the corresponding field

type MetaField map[Bit]Subfield // It is used for compound data type

type Subfield struct {
	Name NameTypeField
	Size uint8
}

type UAP struct {
	Name     string
	Category uint8
	Version  float64
}
type Category uint8
type Version float32

// Profiles contains the defaults User Application Profiles version.
var Profiles = map[uint8]StandardUAP{
	1:   Cat001V12,
	2:   Cat002V10,
	21:  Cat021v10,
	30:  Cat030StrV51,
	32:  Cat032StrV70,
	34:  Cat034V127,
	48:  Cat048V127,
	255: Cat255StrV51,
}
