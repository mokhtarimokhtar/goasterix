package uap

type NameTypeField int

const (
	Fixed NameTypeField = iota
	Extended
	Compound
	Repetitive
	Explicit
	SP
	RE
	RFS
	Spare
)

// DataField describes FRN(Field Reference Number)
type DataField struct {
	FRN         uint8
	DataItem    string
	Description string
	Type        TypeField
	Payload     []byte
}

type TypeField struct {
	//Name string
	Name NameTypeField
	Size uint8
	Meta MetaField
}
type Bit uint8                  // It is the bit number

type Size uint8                 // It is the number of bytes (size) of the corresponding field

type MetaField map[Bit]Subfield // It is used for compound data type

type Subfield struct {
	//Name string
	Name NameTypeField
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

type UAP struct {
	Name string
	Category uint8
	Version float64
}
type Category uint8
type Version float32
/*
// Profiles contains the defaults User Application Profiles version.
var Profiles = map[uint8]StandardUAP{
	2:Cat002V10,
	21:Cat021v10,
	48:Cat048V127,
	34:Cat034V127,
	30:Cat032StrV70,
	32:Cat032StrV70,
	255:Cat255StrV51,
}
//func NewUAProfiles(confUAP map[uint8]string) (map[uint8]StandardUAP) {
//
//	return nil
//}
var Registry = []StandardUAP{
	Cat048V127,
	Cat034V127,
	Cat030ArtasV70,
	Cat001TrackV12,
	Cat001PlotV12,
	Cat002V10,
	Cat030StrV51,
	Cat032StrV70,
	Cat255StrV51,
	Cat021v10,
}


var Prof = []UAP{
	{Name: "CAT048", Category: 48, Version: 1.27},
	{Name: "CAT034", Category: 34, Version: 1.27},
	{Name: "ARTAS", Category: 30, Version: 7.0},
}

func Register(profiles []UAP) (r map[uint8]StandardUAP) {
	r = make(map[uint8]StandardUAP)
	for _, uap := range profiles {
		for _, standardUAP := range Registry {
			if uap.Name == standardUAP.Name && uap.Category == standardUAP.Category && uap.Version == standardUAP.Version {
				r[uap.Category] = standardUAP
			}
		}
	}

	return r
}
*/