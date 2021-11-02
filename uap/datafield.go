package uap

import (
	"encoding/json"
	"os"
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
	Name      string
	Size      uint8
	Meta      MetaField
}
type Bit uint8                 // It is the bit number
type Size uint8                // It is the number of bytes (size) of the corresponding field
type MetaField map[Bit]Subfield // It is used for compound data type

type Subfield struct {
	Name      string
	Size      uint8
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

func (uap *StandardUAP) ReadUAPJson(fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&uap)
	if err != nil {
		//log.Println("Error json:", err)
		return err
	}
	return nil
}

// RegisterUAP todo: finish build register and config
var RegisterUAP = map[uint8][]DataField{
	uint8(48): Cat048V127.Items,
	uint8(34): Cat034V127.Items,
	uint8(30): Cat030StrV51.Items,
	uint8(32): Cat032StrV70.Items,
}

// ConfigUAP defines the versions used while decoding.
var ConfigUAP = map[uint8]StandardUAP{
	48: {Name: "CAT048", Category: 48, Version: 1.27},
	34: {Name: "CAT034", Category: 34, Version: 1.27},
	30: {Name: "STR", Category: 30, Version: 5.1},
	1:  {Name: "CAT001", Category: 1, Version: 1.2},
	2:  {Name: "CAT002", Category: 2, Version: 1.0},
}
