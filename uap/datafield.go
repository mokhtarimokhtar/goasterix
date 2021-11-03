package uap

// DataField describes FRN(Field Reference Number)
type DataField struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        TypeField
	Payload     []byte
}

type TypeField struct {
	Name string
	Size uint8
	Meta MetaField
}
type Bit uint8                  // It is the bit number
type Size uint8                 // It is the number of bytes (size) of the corresponding field
type MetaField map[Bit]Subfield // It is used for compound data type

type Subfield struct {
	Name string
	Size uint8
}

// StandardUAP is User Application Profile
// Cat is ASTERIX Category number (integer)
// Version is ASTERIX version for a category
type StandardUAP struct {
	Name     string
	Category uint8
	Version  float64
	Items    []DataField
}
