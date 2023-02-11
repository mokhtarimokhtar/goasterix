package transform

import (
	"errors"
	"math"
)

var (
	// ErrCharUnknown reports which not found equivalent International Alphabet 5 char.
	ErrCharUnknown = errors.New("[ASTERIX Error] char unknown")
)

type CartesianXYPosition struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type SourceIdentifier struct {
	Sac uint8 `json:"sac" xml:"sac"`
	Sic uint8 `json:"sic" xml:"sic"`
}

type TimeOfDayHighPrecision struct {
	FSI             string  `json:"FSI"`
	TimeOfReception float64 `json:"TimeOfReception"`
}

// sacSic returns a SourceIdentifier with:
// Sac: an integer of System Area TransponderRegisterNumber.
// Sic: an integer of System Identification TransponderRegisterNumber.
func sacSic(data [2]byte) (src SourceIdentifier, err error) {
	src.Sac = data[0]
	src.Sic = data[1]
	return src, nil
}

// TODO: Check if it applies for similar functions. If so, refactor name
// timeOfDay returns a float64 in second (1 bit = 1/128 s)
// Absolute time stamping expressed as Coordinated Universal Time (UTC).
// The time information, coded in three octets, shall reflect the exact time of an event,
// expressed as a number of 1/128 s elapsed since last midnight.
// The time of day value is reset to 0 each day at midnight.
func timeOfDay(data [3]byte) (tod float64, err error) {
	tmp := uint32(data[0])<<16 + uint32(data[1])<<8 + uint32(data[2])
	tod = float64(tmp) / 128
	return tod, nil
}

// timeOfDay returns a float64 in second (1 bit = 1/2^30 s)
// Absolute time stamping expressed as Co-ordinated Universal Time (UTC).
// The time information, coded in three octets, shall reflect the exact time of an event,
// expressed as a number of 1/2^30 s elapsed since last midnight.
// The time of day value is reset to 0 each day at midnight.
func timeOfDayHighPrecision(data [4]byte) (TimeOfDayHighPrecision, error) {
	var tmp TimeOfDayHighPrecision

	fsiVal := data[0] & 0xc0 >> 6
	switch fsiVal {
	case 3:
		tmp.FSI = "Reserved"
	case 2:
		tmp.FSI = "-1"
	case 1:
		tmp.FSI = "+1"
	case 0:
		tmp.FSI = "+0"
	}

	tod := uint32(data[0]&0x3F)<<24 + uint32(data[1])<<16 + uint32(data[2])<<8 + uint32(data[3])
	tmp.TimeOfReception = float64(tod) * math.Pow(2, -30)
	return tmp, nil
}

// trackNumber returns an integer (Identification of a track).
// An integer value representing a unique reference to a track record within a particular track file.
func trackNumber(data [2]byte) uint16 {
	tn := uint16(data[0])<<8 + uint16(data[1])
	return tn
}

// modeSIdentification returns a string.
// It converts each char into ASCII char.
// Aircraft identification (in 8 characters) obtained from an aircraft equipped with a Mode S transponder.
// Ref: 5.2.24 Records Item I048/240, Aircraft Identification.
// The CALL SIGN shall consist of eight characters, which must contain only decimal digits 0-9, the capital letters A-Z,
// and – as trailing pad characters only – the “space” character.
func modeSIdentification(data [6]byte) (string, error) {
	var s string
	var err error

	ch1 := data[0] & 0xFC >> 2
	str1, found1 := TableIA5[ch1]

	ch2 := data[0]&0x03<<4 + data[1]&0xF0>>4
	str2, found2 := TableIA5[ch2]

	ch3 := data[1]&0x0F<<2 + data[2]&0xC0>>6
	str3, found3 := TableIA5[ch3]

	ch4 := data[2] & 0x3F
	str4, found4 := TableIA5[ch4]

	ch5 := data[3] & 0xFC >> 2
	str5, found5 := TableIA5[ch5]

	ch6 := data[3]&0x03<<4 + data[4]&0xF0>>4
	str6, found6 := TableIA5[ch6]

	ch7 := data[4]&0x0F<<2 + data[5]&0xC0>>6
	str7, found7 := TableIA5[ch7]

	ch8 := data[5] & 0x3F
	str8, found8 := TableIA5[ch8]

	if !found1 || !found2 || !found3 || !found4 || !found5 || !found6 || !found7 || !found8 {
		err = ErrCharUnknown
	}

	//s = strings.TrimSpace(str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8)
	s = str1 + str2 + str3 + str4 + str5 + str6 + str7 + str8

	return s, err
}

// checkEqualLatLong returns a boolean (whether two WGS84 supplied coordinates are equal +- epsilon).
func checkEqualLatLong(resultCoordinates WGS84Coordinates, actualCoordinates WGS84Coordinates, epsilon float64) bool {
	compareLatitudes := equalWithinErrorBounds(float64(resultCoordinates.Latitude), float64(actualCoordinates.Latitude), epsilon)
	compareLongitudes := equalWithinErrorBounds(float64(resultCoordinates.Longitude), float64(actualCoordinates.Longitude), epsilon)
	return compareLatitudes && compareLongitudes
}

// equalWithinErrorBounds returns a boolean (whether two float64 values are equal +- epsilon/error).
func equalWithinErrorBounds(actualValue float64, targetValue float64, epsilon float64) bool {
	return math.Abs(targetValue-actualValue) < epsilon
}
