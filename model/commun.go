package model

import (
	"encoding/json"
	"encoding/xml"
	"github.com/mokhtarimokhtar/goasterix/uap"
)

type Writer interface {
	write([]uap.DataField)
}

func WriteModel(w Writer, items []uap.DataField) {
	w.write(items)
}
func WriteModelJSON(w Writer, items []uap.DataField) (j []byte, err error) {
	w.write(items)
	j, err = json.Marshal(w)
	return j, err
}
func WriteModelXML(w Writer, items []uap.DataField) (x []byte, err error) {
	w.write(items)
	x, err = xml.Marshal(w)
	return x, err
}

type CartesianXYPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type SourceIdentifier struct {
	Sac uint8 `json:"sac" xml:"sac"`
	Sic uint8 `json:"sic" xml:"sic"`
}

// sacSic returns a SourceIdentifier with:
// Sac: an integer of System Area TransponderRegisterNumber.
// Sic: an integer of System Identification TransponderRegisterNumber.
func sacSic(data [2]byte) (src SourceIdentifier, err error) {
	src.Sac = data[0]
	src.Sic = data[1]
	return src, nil
}

// timeOfDay returns a float64 in second (1 bit = 1/128 s)
// Absolute time stamping expressed as Co-ordinated Universal Time (UTC).
// The time information, coded in three octets, shall reflect the exact time of an event,
// expressed as a number of 1/128 s elapsed since last midnight.
// The time of day value is reset to 0 each day at midnight.
func timeOfDay(data [3]byte) (tod float64, err error) {
	var tmp uint32
	tmp = uint32(data[0])<<16 + uint32(data[1])<<8 + uint32(data[2])
	tod = float64(tmp) / 128
	return tod, nil
}