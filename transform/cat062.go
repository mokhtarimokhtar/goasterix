package transform

import (
	"encoding/hex"
	"github.com/mokhtarimokhtar/goasterix"
	"github.com/mokhtarimokhtar/goasterix/uap"
	"math"
	"strconv"
	"strings"
)

type TrackVelocity struct {
	Vx float32 `json:"vx,omitempty"`
	Vy float32 `json:"vy,omitempty"`
}
type Acceleration struct {
	Ax float32 `json:"ax,omitempty"`
	Ay float32 `json:"ay,omitempty"`
}
type PositionWGS84 struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}
type TrackMode3A struct {
	V      string `json:"v"`
	G      string `json:"g"`
	CH     string `json:"ch"`
	Squawk string `json:"squawk"`
}
type TargetIdent struct {
	Target string `json:"target,omitempty"`
	STI    string `json:"sti,omitempty"`
}
type BarometricAltitude struct {
	QNH      string  `json:"qnh,omitempty"`
	Altitude float64 `json:"altitude,omitempty"`
}

type IAS struct {
	IM string `json:"im"`
	AirSpeed float64 `json:"airSpeed"`
}
type DerivedData struct {
	TargetAddress        string  `json:"targetAddress,omitempty"`
	TargetIdentification string  `json:"targetIdentification,omitempty"`
	MagneticHeading      float64 `json:"magneticHeading,omitempty"`
	IndicatedAirspeed    *IAS    `json:"indicatedAirspeed,omitempty"`
	AirSpeed             uint16  `json:"airSpeed,omitempty"`
}

type Cat062Model struct {
	SacSic                *SourceIdentifier    `json:"sourceIdentifier,omitempty"`
	ServiceIdentification uint8                `json:"serviceIdentification,omitempty"`
	TimeOfDay             float64              `json:"timeOfDay,omitempty"`
	TrackPositionWGS84    *PositionWGS84       `json:"trackPositionWGS84"`
	CartesianXY           *CartesianXYPosition `json:"cartesianXY,omitempty"`
	TrackVelocity         *TrackVelocity       `json:"trackVelocity,omitempty"`
	Acceleration          *Acceleration        `json:"acceleration,omitempty"`
	Mode3ACode            *TrackMode3A         `json:"mode3ACode,omitempty"`
	TargetIdentification  *TargetIdent         `json:"targetIdentification,omitempty"`
	AircraftDerivedData   *DerivedData         `json:"aircraftDerivedData,omitempty"`
	TrackNumber           uint16               `json:"trackNumber,omitempty"`
	FlightLevel           float32              `json:"flightLevel,omitempty"`
	GeometricAltitude     float32              `json:"geometricAltitude,omitempty"`
	BarometricAltitude    *BarometricAltitude  `json:"barometricAltitude,omitempty"`
	RateOfClimbDescent    float32              `json:"rateOfClimbDescent,omitempty"`
}

// Write writes a single ASTERIX Record to Cat062Model.
// CompoundItems is a slice of CompoundItems DataField.
func (data *Cat062Model) write(items []uap.DataField) {
	for _, item := range items {
		switch item.FRN {
		case 1:
			// decode sac sic
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp, _ := sacSic(payload)
			data.SacSic = &tmp
		// case 2 is spare
		case 3:
			// Service Identification
			data.ServiceIdentification = item.Payload[0]
		case 4:
			// Time Of Track Information
			var payload [3]byte
			copy(payload[:], item.Payload[:])
			data.TimeOfDay, _ = timeOfDay(payload)
		case 5:
			// Calculated Track Position (WGS-84)
			var payload [8]byte
			copy(payload[:], item.Payload[:])
			tmp := calculatedTrackPositionWGS84(payload)
			data.TrackPositionWGS84 = &tmp
		case 6:
			// Calculated Track Position. (Cartesian)
			var payload [6]byte
			copy(payload[:], item.Payload[:])
			tmp := calculatedTrackPositionCartesian(payload)
			data.CartesianXY = &tmp
		case 7:
			// Calculated Track Velocity (Cartesian)
			var payload [4]byte
			copy(payload[:], item.Payload[:])
			tmp := calculatedTrackVelocityCartesian(payload)
			data.TrackVelocity = &tmp
		case 8:
			// Calculated Acceleration (Cartesian)
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := calculatedAccelerationCartesian(payload)
			data.Acceleration = &tmp
		case 9:
			// Track Mode 3/A Code
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := mode3ACode(payload)
			data.Mode3ACode = &tmp
		case 10:
			// Target Identification
			var payload [7]byte
			copy(payload[:], item.Payload[:])
			tmp := targetIdentification(payload)
			data.TargetIdentification = &tmp
		case 11:
			// Aircraft Derived Data
			tmp := extractDerivedData(item.Payload)
			data.AircraftDerivedData = &tmp

		case 12:
			// Track Number
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			data.TrackNumber = trackNumber(payload)
		// todo case 13
		// todo case 14
		// todo case 15
		// todo case 16
		case 17:
			// Measured Flight Level
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			data.FlightLevel = measuredFlightLevel(payload)
		case 18:
			// Calculated Track Geometric Altitude
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			data.GeometricAltitude = trackGeometricAltitude(payload)
		case 19:
			// Calculated Track Barometric Altitude
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			tmp := trackBarometricAltitude(payload)
			data.BarometricAltitude = &tmp
		case 20:
			// Calculated Rate Of Climb/Descent
			var payload [2]byte
			copy(payload[:], item.Payload[:])
			data.RateOfClimbDescent = rateOfClimbDescent(payload)
			// todo case 21
			// todo case 22
			// todo case 23
			// todo case 24
			// todo case 25
			// todo case 26
			// todo case 27
			// todo case 28
			// todo case 34
			// todo case 35
		}
	}
}


// extractDerivedData returns Data derived directly by the aircraft.
func extractDerivedData(data []byte) DerivedData {
	var dd DerivedData
	offset := uint8(1)
	for {
		if data[offset-1]&0x01 == 0 {
			break
		}
		offset++
	}
	if data[0]&0x80 != 0 {
		dd.TargetAddress = strings.ToUpper(hex.EncodeToString(data[offset : offset+3]))
		offset = offset + 3
	}
	if data[0]&0x40 != 0 {
		tmp := [6]byte{data[offset], data[offset+1], data[offset+2], data[offset+3], data[offset+4], data[offset+5]}
		dd.TargetIdentification, _ = modeSIdentification(tmp)
		offset = offset + 6
	}
	if data[0]&0x20 != 0 {
		dd.MagneticHeading = float64(uint16(data[offset])<<8+uint16(data[offset+1])) * 0.0055
		_ = offset + 2
	}
	if data[0]&0x10 != 0 {
		var lsb float64
		if data[offset]&0x80 != 0 {
			dd.IndicatedAirspeed.IM = "mach"
			lsb = 0.001
		} else {
			dd.IndicatedAirspeed.IM = "ias"
			lsb = 0.000061035
		}
		dd.IndicatedAirspeed.AirSpeed = float64(uint16(data[offset]&0x7f)<<8+uint16(data[offset+1])) * lsb
		offset = offset + 2
	}
	if data[0]&0x08 != 0 {
		dd.AirSpeed = uint16(data[offset])<<8+uint16(data[offset+1])
		_ = offset + 2
	}

	return dd
}

// calculatedTrackPositionWGS84 returns Latitude and Longitude.
// Calculated Position in WGS-84 Co-ordinates with a resolution of 180/2^25 degrees
func calculatedTrackPositionWGS84(data [8]byte) PositionWGS84 {
	var pos PositionWGS84
	lsb := 180 / math.Pow(2, 25)
	pos.Latitude = float64(int64(data[0])<<24+int64(data[1])<<16+int64(data[2])<<8+int64(data[3])) * lsb
	pos.Longitude = float64(int64(data[4])<<24+int64(data[5])<<16+int64(data[6])<<8+int64(data[7])) * lsb
	return pos
}

// calculatedTrackPositionCartesian returns X and Y float64 in m
// Calculated position in Cartesian co-ordinates with a resolution of 0.5m
// LSB = 0.5
func calculatedTrackPositionCartesian(data [6]byte) CartesianXYPosition {
	var pos CartesianXYPosition

	tmpX := uint32(data[0])<<16 + uint32(data[1])<<8 + uint32(data[2])
	pos.X = float64(goasterix.TwoComplement32(24, tmpX)) * 0.5

	tmpY := uint32(data[3])<<16 + uint32(data[4])<<8 + uint32(data[5])
	pos.Y = float64(goasterix.TwoComplement32(24, tmpY)) * 0.5

	return pos
}

// calculatedTrackVelocityCartesian returns Vx and Vy float32 in m/s
// Calculated track velocity expressed in Cartesian co-ordinates
func calculatedTrackVelocityCartesian(data [4]byte) TrackVelocity {
	var vel TrackVelocity
	vel.Vx = float32(int16(data[0])<<8+int16(data[1])) / 4
	vel.Vy = float32(int16(data[2])<<8+int16(data[3])) / 4
	return vel
}

// calculatedAccelerationCartesian returns Ax and Ay float32 in m/s^2.
// Calculated Acceleration of the target expressed in Cartesian co-ordinates
// LSB = 0.25 m/s^2
func calculatedAccelerationCartesian(data [2]byte) Acceleration {
	var acc Acceleration
	acc.Ax = float32(int8(data[0])) / 4
	acc.Ay = float32(int8(data[1])) / 4
	return acc
}

// mode3ACode returns the squawk.
// Mode-3/A code converted into octal representation
func mode3ACode(data [2]byte) TrackMode3A {
	var mode3A TrackMode3A
	if data[0]&0x80 != 0 {
		mode3A.V = "code_not_validated"
	} else {
		mode3A.V = "code_validated"
	}

	if data[0]&0x40 != 0 {
		mode3A.G = "garbled_code"
	} else {
		mode3A.G = "default"
	}

	if data[0]&0x20 != 0 {
		mode3A.CH = "changed"
	} else {
		mode3A.CH = "no_change"
	}

	tmp := uint16(data[0])&0x000F<<8 + uint16(data[1])&0x00FF
	mode3A.Squawk = strconv.FormatUint(uint64(tmp), 8)

	return mode3A
}

// Target (aircraft or vehicle) identification in 8 characters
func targetIdentification(data [7]byte) TargetIdent {
	var target TargetIdent
	tmp := data[0] & 0xc0 >> 6

	switch tmp {
	case 0:
		target.STI = "downlinked_target"
	case 1:
		target.STI = "callsign_not_downlinked_target"
	case 2:
		target.STI = "registration_not_downlinked_target"
	case 3:
		target.STI = "invalid"
	}
	tmpData := [6]byte{data[1], data[2], data[3], data[4], data[5], data[6]}
	target.Target, _ = modeSIdentification(tmpData)

	return target
}

// measuredFlightLevel returns level in 100's ft
func measuredFlightLevel(data [2]byte) float32 {
	fl := float32(int16(data[0])<<8+int16(data[1])) / 4
	return fl
}

// trackGeometricAltitude returns a float32 in ft
// Vertical distance between the target and the projection of its position
// on the earth’s ellipsoid, as defined by WGS84
func trackGeometricAltitude(data [2]byte) float32 {
	fl := float32(int16(data[0])<<8+int16(data[1])) * 6.25
	return fl
}

// trackBarometricAltitude returns Altitude in FL
// Calculated Barometric Altitude of the track
func trackBarometricAltitude(data [2]byte) BarometricAltitude {
	var ba BarometricAltitude
	if data[0]&0x80 != 0 {
		ba.QNH = "qnh_correction_applied"
	} else {
		ba.QNH = "no_qnh_correction_applied"
	}
	ba.Altitude = float64((uint16(data[0])<<8+uint16(data[1]))&0x7FFF) / 4
	return ba
}

// rateOfClimbDescent returns a float32 in feet/minute
// Calculated rate of Climb/Descent of an aircraft in feet/minute
// A positive value indicates a climb, whereas a negative value indicates a descent.
func rateOfClimbDescent(data [2]byte) float32 {
	rate := float32(int16(data[0])<<8+int16(data[1])) * 6.25
	return rate
}
