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
	Compound    []DataField
	RFS         []DataField
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

