package bdscode

import (
	"github.com/mokhtarimokhtar/goasterix"
	"math"
)

// Code60 Heading and speed report
// MagneticHeading Range = [–180, +180] degrees
// IndicatedAirspeed Range = [0, 1023] knots
// Mach Range = [0, 4.092] MACH
// BarometricAltitudeRate Range = [–16 384, +16 352] feet/minute
// InertialVerticalVelocity
type Code60 struct {
	MagneticHeading                int16   `json:"magneticHeading,omitempty"`
	MagneticHeadingStatus          bool    `json:"-"`
	IndicatedAirspeed              uint16  `json:"indicatedAirspeed,omitempty"`
	IndicatedAirspeedStatus        bool    `json:"-"`
	Mach                           float64 `json:"mach,omitempty"`
	MachStatus                     bool    `json:"-"`
	BarometricAltitudeRate         int16   `json:"barometricAltitudeRate,omitempty"`
	BarometricAltitudeRateStatus   bool    `json:"-"`
	InertialVerticalVelocity       int16   `json:"inertialVerticalVelocity,omitempty"`
	InertialVerticalVelocityStatus bool    `json:"-"`
}

func (c *Code60) Decode(data [7]byte) (err error) {
	// Extract MagneticHeading
	if (data[0] & 0x80 >> 7) == 1 {
		c.MagneticHeadingStatus = true
		mh := uint16(data[0]&0x7F)<<4 + uint16(data[1]&0xF0)>>4
		tmpMh := goasterix.TwoComplement16(11, mh)
		c.MagneticHeading = int16(float64(tmpMh) * 90 / 512)
	}

	// Extract IndicatedAirspeed
	if (data[1] & 0x08 >> 3) == 1 {
		c.IndicatedAirspeedStatus = true
		ias := uint16(data[1]&0x07)<<7 + uint16(data[2]&0xFE)>>1
		c.IndicatedAirspeed = ias
	}

	// Extract Mach
	if (data[2] & 0x01) == 1 {
		c.MachStatus = true
		mach := uint16(data[3]&0xFF)<<2 + uint16(data[4]&0xC0)>>6
		c.Mach = math.Round(float64(mach)*2.048/512*1000) / 1000
	}

	// Extract BarometricAltitudeRate
	if (data[4] & 0x20 >> 5) == 1 {
		c.BarometricAltitudeRateStatus = true
		bar := uint16(data[4]&0x1F)<<5 + uint16(data[5]&0xF8)>>3
		tmpBar := goasterix.TwoComplement16(10, bar)
		c.BarometricAltitudeRate = tmpBar * 32
	}

	// Extract InertialVerticalVelocity
	if (data[5] & 0x04 >> 2) == 1 {
		c.InertialVerticalVelocityStatus = true
		ivv := uint16(data[5]&0x03)<<8 + uint16(data[6]&0xFF)
		tmpIvv := goasterix.TwoComplement16(10, ivv)
		c.InertialVerticalVelocity = tmpIvv * 32
	}

	return err
}
